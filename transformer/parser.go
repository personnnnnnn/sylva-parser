package transformer

import (
	"fmt"
	"slices"
)

type Parser struct {
	Tokens     []*Token
	TokenIndex int
	Errors     []*SylvaError
}

func (p *Parser) Step() {
	if p.TokenIndex >= len(p.Tokens)-1 {
		return
	}
	p.TokenIndex++
}

func (p *Parser) Token() *Token {
	return p.Tokens[p.TokenIndex]
}

func (p *Parser) GetToken() *Token {
	token := p.Token()
	p.Step()
	return token
}

func MakeParser(tokens []*Token) *Parser {
	return &Parser{
		Tokens:     tokens,
		TokenIndex: 0,
		Errors:     []*SylvaError{},
	}
}

func (p *Parser) Parse() Node {
	root := p.Statement()
	p.Errors = root.GetErrors()

	endToken := p.GetToken()
	if endToken.Type != TT_EOF {
		start := endToken.Bounds.Start
		end := p.Tokens[len(p.Tokens)-1].Bounds.End

		p.Errors = append(p.Errors, &SylvaError{
			Type:    "parse error",
			Bounds:  MakeBounds(start, end),
			Message: "expected EOF",
		})
	}

	return root
}

func (p *Parser) Int() Node {
	token := p.GetToken()
	if token.Type != TT_Int {
		return ErrNode(token.Bounds, "parse error", "expected int")
	}
	return &IntNode{
		NodeTraits: NodeTraits{
			Bounds: token.Bounds,
			Errors: nil,
		},
		Number: token.Value.(int64),
	}
}

func (p *Parser) Float() Node {
	token := p.GetToken()
	if token.Type != TT_Float {
		return ErrNode(token.Bounds, "parse error", "expected float")
	}
	return &FloatNode{
		NodeTraits: NodeTraits{
			Bounds: token.Bounds,
			Errors: []*SylvaError{},
		},
		Number: token.Value.(float64),
	}
}

func (p *Parser) Parens() Node {
	opening := p.GetToken()
	if opening.Type != TT_LParen {
		return ErrNode(
			opening.Bounds,
			"parse error",
			"expected '('",
		)
	}
	expr := p.Expr()
	endToken := p.Token()

	if endToken.Type != TT_RParen {
		expr.AppendError(&SylvaError{
			Bounds:  endToken.Bounds,
			Type:    "parse error",
			Message: "expected ')'",
		})
	} else {
		p.Step()
	}

	return expr
}

func (p *Parser) Symbol() Node {
	token := p.GetToken()
	if token.Type != TT_Symbol {
		return ErrNode(
			token.Bounds,
			"parse error",
			"expected symbol",
		)
	}
	return &SymbolNode{
		NodeTraits: NodeTraits{
			Bounds: token.Bounds,
			Errors: nil,
		},
		Symbol: token.Value.(string),
	}
}

func (p *Parser) Literal() Node {
	node := p.Atom()

	for {
		currToken := p.Token()
		done := false
		switch currToken.Type {
		default:
			done = true
		case TT_Dot:
			p.Step()
			oldNode := node
			symbolToken := p.Token()

			var symbol *SymbolNode
			if symbolToken.Type != TT_Symbol {

				symbol = &SymbolNode{
					NodeTraits: NodeTraits{
						Bounds: symbolToken.Bounds,
						Errors: []*SylvaError{
							{
								Bounds:  symbolToken.Bounds,
								Type:    "parse error",
								Message: "expected symbol",
							},
						},
					},
					Symbol: "<NONE>",
				}
			} else {
				symbol = p.Symbol().(*SymbolNode)
			}

			node = &AttributeAccessNode{
				NodeTraits: NodeTraits{
					Bounds: MakeBounds(oldNode.GetBounds().Start, symbol.GetBounds().End),
					Errors: nil,
				},
				Value:     oldNode,
				Attribute: symbol,
			}

			node.AppendErrors(oldNode)
			node.AppendErrors(symbol)
		case TT_LBrack:
			p.Step()
			expr := p.Expr()
			endToken := p.Token()

			oldNode := node
			node = &IndexAccessNode{
				NodeTraits: NodeTraits{
					Bounds: MakeBounds(oldNode.GetBounds().Start, endToken.Bounds.End),
					Errors: nil,
				},
				Value: oldNode,
				Index: expr,
			}

			node.AppendErrors(oldNode)
			node.AppendErrors(expr)

			if endToken.Type == TT_RBrack {
				p.Step()
			} else {
				node.AppendError(&SylvaError{
					Bounds:  endToken.Bounds,
					Type:    "parse error",
					Message: "expected ']'",
				})
			}
		}

		if done {
			break
		}
	}

	return node
}

func (p *Parser) Atom() Node {
	token := p.Token()
	switch token.Type {
	case TT_Int:
		return p.Int()
	case TT_Float:
		return p.Float()
	case TT_True, TT_False:
		return p.Bool()
	case TT_String:
		return p.StringNode()
	case TT_Symbol:
		return p.Symbol()
	case TT_LParen:
		return p.Parens()
	default:
		p.Step()
		return ErrNode(token.Bounds, "parse error", "expected literal")
	}
}

func (p *Parser) Bool() Node {
	token := p.GetToken()
	if !(token.Type == TT_True || token.Type == TT_False) {
		return ErrNode(token.Bounds, "parse error", "extected true or false")
	}
	return &BoolNode{
		NodeTraits: NodeTraits{
			Bounds: token.Bounds,
			Errors: nil,
		},
		Bool: token.Type == TT_True,
	}
}

func (p *Parser) Expr() Node {
	return p.Logic()
}

func (p *Parser) Logic() Node {
	return p.OpRule(p.Comparison, TT_And, TT_Or)
}

func (p *Parser) Comparison() Node {
	return p.OpRule(p.Concat, TT_Lt, TT_Gt, TT_Lte, TT_Gte, TT_Eq, TT_Neq)
}

func (p *Parser) Concat() Node {
	return p.OpRule(p.AddOrSub, TT_DDot)
}

func (p *Parser) AddOrSub() Node {
	return p.OpRule(p.MulOrDiv, TT_Plus, TT_Minus)
}

func (p *Parser) MulOrDiv() Node {
	return p.OpRule(p.Value, TT_Mul, TT_Div, TT_Mod)
}

func (p *Parser) OpRule(opRule func() Node, tokens ...TokenType) Node {
	left := opRule()
	for {
		opToken := p.Token()
		if !slices.Contains(tokens, opToken.Type) {
			break
		}

		p.Step()
		right := opRule()
		oldLeft := left

		operator, err := TTtoBinaryOperator((opToken.Type))
		if err != nil {
			fmt.Println("error:", err)
		}

		left = &BinOpExpr{
			NodeTraits: NodeTraits{
				Bounds: MakeBounds(left.GetBounds().Start, right.GetBounds().End),
				Errors: nil,
			},
			Left:  left,
			Right: right,
			Op:    operator,
		}

		left.AppendErrors(oldLeft)
		left.AppendErrors(right)
	}

	return left
}

func (p *Parser) Item() Node {
	token := p.Token()
	if token.Type == TT_TDot {
		p.Step()
		value := p.Expr()
		node := &SpreadNode{
			NodeTraits: NodeTraits{
				Bounds: MakeBounds(token.Bounds.Start, value.GetBounds().End),
				Errors: nil,
			},
			Value: value,
		}
		node.AppendErrors(value)
		return node
	}
	return p.Expr()
}

func (p *Parser) ArgList() *ArgListNode {
	startToken := p.GetToken()
	if startToken.Type != TT_LParen {
		return &ArgListNode{
			NodeTraits: NodeTraits{
				Bounds: startToken.Bounds,
				Errors: []*SylvaError{
					{
						Bounds:  startToken.Bounds,
						Type:    "parse error",
						Message: "expected '('",
					},
				},
			},
			Arguments: []Node{},
		}
	}

	args := []Node{}
	node := &ArgListNode{}

	if p.Token().Type == TT_RParen {
		endToken := p.GetToken()
		node.Bounds.Start = startToken.Bounds.Start
		node.Bounds.End = endToken.Bounds.End
		node.Arguments = args
		return node
	}

	firstExpr := p.Expr()
	node.AppendErrors(firstExpr)
	args = append(args, firstExpr)

	for p.Token().Type == TT_Comma {
		p.Step()

		nextToken := p.Token()
		if nextToken.Type == TT_RParen {
			break
		}

		expr := p.Item()
		node.AppendErrors(expr)
		args = append(args, expr)
	}

	if p.Token().Type != TT_RParen {
		toRemove := []int{}
		bounds := p.Token().Bounds
		for i, err := range node.GetErrors() {
			if err.Bounds == bounds && err.Message == "expected literal" {
				toRemove = append(toRemove, i)
			}
		}
		for _, i := range toRemove {
			node.SetErrors(Remove(node.GetErrors(), i))
		}
		if p.Tokens[p.TokenIndex-1].Type == TT_Comma {
			node.AppendError(&SylvaError{
				Bounds:  bounds,
				Type:    "parse error",
				Message: "expected expression or ')'",
			})
		} else {
			node.AppendError(&SylvaError{
				Bounds:  bounds,
				Type:    "parse error",
				Message: "expected ',' or ')'",
			})
		}
	}

	for p.Token().Type != TT_RParen && p.Token().Type != TT_EOF {
		p.Step()
	}

	p.Step()

	node.Arguments = args

	return node
}

func (p *Parser) Call() Node {
	left := p.Literal()
	for {
		opToken := p.Token()
		if opToken.Type != TT_LParen {
			break
		}

		right := p.ArgList()
		oldLeft := left

		left = &CallNode{
			NodeTraits: NodeTraits{
				Bounds: MakeBounds(left.GetBounds().Start, right.GetBounds().End),
				Errors: nil,
			},
			Function:  left,
			Arguments: right,
		}

		left.AppendErrors(oldLeft)
		left.AppendErrors(right)
	}

	return left
}

func (p *Parser) StringNode() Node {
	token := p.GetToken()
	if token.Type != TT_String {
		return ErrNode(token.Bounds, "parse error", "expected string")
	}
	return &StringNode{
		NodeTraits: NodeTraits{
			Bounds: token.Bounds,
			Errors: nil,
		},
		Text: token.Value.(string),
	}
}

func (p *Parser) Value() Node {
	unaryOp, err := TTtoUnaryOperator(p.Token().Type)
	if err != nil {
		node := p.Call()
		return node
	}
	start := p.Token().Bounds.Start
	p.Step()
	value := p.Value()

	node := &UnaryOpNode{
		NodeTraits: NodeTraits{
			Bounds: MakeBounds(start, value.GetBounds().End),
		},
		Op:    unaryOp,
		Value: value,
	}

	node.AppendErrors(value)

	return node
}

func (p *Parser) Variable() Node {
	oldIndex := p.TokenIndex

	variable := p.Literal()
	switch variable.(type) {
	case *SymbolNode, *IndexAccessNode, *AttributeAccessNode:
		return variable
	}

	p.TokenIndex = oldIndex
	token := p.GetToken()
	return ErrNode(
		token.Bounds,
		"parse error",
		"expected variable",
	)
}

func (p *Parser) VariableAssignment() Node {
	node := &VariableAssignmentNode{}
	variable := p.Variable()
	node.Variable = variable
	node.AppendErrors(node)

	eqToken := p.GetToken()
	if eqToken.Type != TT_EqSign {
		return ErrNode(
			eqToken.Bounds,
			"parse error",
			"expected '='",
		)
	}

	expr := p.Expr()
	node.Value = expr
	node.AppendErrors(expr)

	node.Bounds.Start = variable.GetBounds().Start
	node.Bounds.End = expr.GetBounds().End

	return node
}

func (p *Parser) VariableListAssignment() Node {
	node := &VariableListAssignmentNode{}
	variables := []Node{}
	variable := p.Variable()
	variables = append(variables, variable)
	node.AppendErrors(variable)

	for {
		comma := p.Token()
		if comma.Type == TT_EqSign {
			break
		}
		if comma.Type != TT_Comma {
			return ErrNode(
				comma.Bounds,
				"parse error",
				"expected ','",
			)
		}
		p.Step()
		variable := p.Variable()
		variables = append(variables, variable)
		node.AppendErrors(variable)
	}

	p.Step()

	expr := p.Expr()
	node.Value = expr
	node.Variables = variables

	node.AppendErrors(expr)

	node.Bounds.Start = variable.GetBounds().Start
	node.Bounds.End = expr.GetBounds().End

	return node
}

func (p *Parser) VariableDefinition() Node {
	letToken := p.Token()
	if letToken.Type != TT_Let {
		return ErrNode(letToken.Bounds, "parse error", "expected 'let'")
	}
	p.Step()

	node := &VariableDefinitionNode{}
	symbolToken := p.Token()

	var variable *SymbolNode
	if symbolToken.Type != TT_Symbol {
		variable = &SymbolNode{
			NodeTraits: NodeTraits{
				Bounds: symbolToken.Bounds,
				Errors: []*SylvaError{
					{
						Bounds:  symbolToken.Bounds,
						Type:    "parse error",
						Message: "expected symbol",
					},
				},
			},
			Symbol: "<NONE>",
		}
	} else {
		variable = p.Symbol().(*SymbolNode)
	}

	node.Variable = variable
	node.AppendErrors(node)

	eqToken := p.GetToken()
	if eqToken.Type != TT_EqSign {
		return ErrNode(
			eqToken.Bounds,
			"parse error",
			"expected '='",
		)
	}

	expr := p.Expr()
	node.Value = expr
	node.AppendErrors(expr)

	node.Bounds.Start = letToken.Bounds.Start
	node.Bounds.End = expr.GetBounds().End

	return node
}

func (p *Parser) VariableListDefinition() Node {
	letToken := p.Token()
	if letToken.Type != TT_Let {
		return ErrNode(letToken.Bounds, "parse error", "expected 'let'")
	}
	p.Step()

	variables := []*SymbolNode{}

	symbolToken := p.Token()
	var variable *SymbolNode
	if symbolToken.Type != TT_Symbol {
		variable = &SymbolNode{
			NodeTraits: NodeTraits{
				Bounds: symbolToken.Bounds,
				Errors: []*SylvaError{
					{
						Bounds:  symbolToken.Bounds,
						Type:    "parse error",
						Message: "expected symbol",
					},
				},
			},
			Symbol: "<NONE>",
		}
	} else {
		variable = p.Symbol().(*SymbolNode)
	}

	variables = append(variables, variable)

	expectValue := false
	for {
		comma := p.Token()
		if comma.Type == TT_EqSign {
			expectValue = true
			break
		}
		if comma.Type != TT_Comma {
			break
		}
		p.Step()

		symbolToken := p.Token()
		var variable *SymbolNode
		if symbolToken.Type != TT_Symbol {
			variable = &SymbolNode{
				NodeTraits: NodeTraits{
					Bounds: symbolToken.Bounds,
					Errors: []*SylvaError{
						{
							Bounds:  symbolToken.Bounds,
							Type:    "parse error",
							Message: "expected symbol",
						},
					},
				},
				Symbol: "<NONE>",
			}
		} else {
			variable = p.Symbol().(*SymbolNode)
		}

		variables = append(variables, variable)
	}

	if expectValue {
		p.Step()

		node := &VariableListDefinitionNode{}

		expr := p.Expr()
		node.Value = expr
		node.Variables = variables

		for _, variable := range node.Variables {
			node.AppendErrors(variable)
		}

		node.AppendErrors(expr)

		node.Bounds.Start = letToken.Bounds.Start
		node.Bounds.End = expr.GetBounds().End
		if len(variables) == 1 {
			return ErrNode(node.Bounds, "parse error", "expected 2 or more variables for a list definition")
		}

		return node
	}

	node := &VariableDeclarationNode{}

	node.Variables = variables
	for _, variable := range node.Variables {
		node.AppendErrors(variable)
	}

	node.Bounds.Start = letToken.Bounds.Start
	node.Bounds.End = variables[len(variables)-1].Bounds.End

	return node
}

func (p *Parser) Statement() Node {
	oldIndex := p.TokenIndex

	token := p.Token()
	if token.Type == TT_Let {
		varListDef := p.VariableListDefinition()
		if _, ok := varListDef.(*ErrorNode); !ok {
			return varListDef
		}
		p.TokenIndex = oldIndex
		return p.VariableDefinition()
	}

	// a simple symbol or a function call (like 'print("Hi :)")')
	// might get confused as a VariableAssignment ('print = 0')
	// but, in that case (where there is no "=" after the symbol)
	// it will return an ErrorNode
	varAssignment := p.VariableAssignment()
	if _, ok := varAssignment.(*ErrorNode); !ok {
		return varAssignment
	} else {
		p.TokenIndex = oldIndex
		varListAssignment := p.VariableListAssignment()
		if _, ok := varListAssignment.(*ErrorNode); !ok {
			return varListAssignment
		}
	}

	p.TokenIndex = oldIndex
	return p.Expr()
}
