package transformer

import (
	"fmt"
)

type ErrorNode struct {
	NodeTraits
}

func ErrNode(bounds Bounds, type_ string, msg string) *ErrorNode {
	return &ErrorNode{
		NodeTraits: NodeTraits{
			Errors: []*SylvaError{{
				Type:    type_,
				Bounds:  bounds,
				Message: msg,
			}},
			Bounds: bounds,
		},
	}
}

func (*ErrorNode) KindName() string {
	return "error"
}

func (e *ErrorNode) Marshal() {
	SetKind(e)
}

type IntNode struct {
	NodeTraits
	Number int64 `json:"number"`
}

func (i *IntNode) Marshal() {
	SetKind(i)
}

func (*IntNode) KindName() string {
	return "int"
}

func (i *IntNode) String() string {
	return fmt.Sprintf("(int %v)", i.Number)
}

type FloatNode struct {
	NodeTraits
	Number float64 `json:"number"`
}

func (f *FloatNode) Marshal() {
	SetKind(f)
}

func (*FloatNode) KindName() string {
	return "float"
}

func (i *FloatNode) String() string {
	return fmt.Sprintf("(float %v)", i.Number)
}

type SymbolNode struct {
	NodeTraits
	Symbol string `json:"symbol"`
}

func (s *SymbolNode) Marshal() {
	SetKind(s)
}

func (*SymbolNode) KindName() string {
	return "symbol"
}

func (s *SymbolNode) String() string {
	return fmt.Sprintf("(symbol: %v)", s.Symbol)
}

type ArgListNode struct {
	NodeTraits
	Arguments []Node `json:"arguments"`
}

func (a *ArgListNode) Marshal() {
	SetKind(a)
	for _, arg := range a.Arguments {
		arg.Marshal()
	}
}

func (*ArgListNode) KindName() string {
	return "argList"
}

func (a *ArgListNode) String() string {
	return fmt.Sprintf("(arg-list: %v)", a.Arguments)
}

type CallNode struct {
	NodeTraits
	Function  Node         `json:"function"`
	Arguments *ArgListNode `json:"arguments"`
}

func (c *CallNode) Marshal() {
	SetKind(c)
	c.Function.Marshal()
	c.Arguments.Marshal()
}

func (*CallNode) KindName() string {
	return "call"
}

func (c *CallNode) String() string {
	return fmt.Sprintf("(call kind: %v, fn: %v, args: %v)", c.Kind, c.Function, c.Arguments)
}

type BinOpExpr struct {
	NodeTraits
	Op    BinaryOperator `json:"op"`
	Left  Node           `json:"left"`
	Right Node           `json:"right"`
}

func (b *BinOpExpr) Marshal() {
	SetKind(b)
	b.Left.Marshal()
	b.Right.Marshal()
}

func (*BinOpExpr) KindName() string {
	return "binop"
}

func (b *BinOpExpr) String() string {
	return fmt.Sprintf("(kind: %v, op: %v, left: %v, right: %v)", b.Kind, b.Op, b.Left, b.Right)
}

type BoolNode struct {
	NodeTraits
	Bool bool `json:"bool"`
}

func (*BoolNode) KindName() string {
	return "bool"
}

func (b *BoolNode) Marshal() {
	SetKind(b)
}

func (b *BoolNode) String() string {
	return fmt.Sprintf("(bool: %v)", b.Bool)
}

type StringNode struct {
	NodeTraits
	Text string `json:"text"`
}

func (*StringNode) KindName() string {
	return "string"
}

func (s *StringNode) Marshal() {
	SetKind(s)
}

func (s *StringNode) String() string {
	return fmt.Sprintf("(string: %v)", s.Text)
}

type UnaryOpNode struct {
	NodeTraits
	Op    UnaryOperator `json:"op"`
	Value Node          `json:"value"`
}

func (*UnaryOpNode) KindName() string {
	return "unaryop"
}

func (u *UnaryOpNode) Marshal() {
	SetKind(u)
	u.Value.Marshal()
}

func (u *UnaryOpNode) String() string {
	return fmt.Sprintf("(unaryop: %v, value: %v)", u.Op, u.Value)
}

type AttributeAccessNode struct {
	NodeTraits
	Value     Node        `json:"value"`
	Attribute *SymbolNode `json:"attribute"`
}

func (*AttributeAccessNode) KindName() string {
	return "attrAccess"
}

func (a *AttributeAccessNode) Marshal() {
	SetKind(a)
	a.Value.Marshal()
	a.Attribute.Marshal()
}

func (a *AttributeAccessNode) String() string {
	return fmt.Sprintf("(attr-access value: %v, attr: %v)", a.Value, a.Attribute.Symbol)
}

type IndexAccessNode struct {
	NodeTraits
	Value Node `json:"value"`
	Index Node `json:"index"`
}

func (*IndexAccessNode) KindName() string {
	return "attrAccess"
}

func (i *IndexAccessNode) Marshal() {
	SetKind(i)
	i.Value.Marshal()
	i.Index.Marshal()
}

func (i *IndexAccessNode) String() string {
	return fmt.Sprintf("(index-access value: %v, index: %v)", i.Value, i.Index)
}

type SpreadNode struct {
	NodeTraits
	Value Node `json:"value"`
}

func (*SpreadNode) KindName() string {
	return "spread"
}

func (s *SpreadNode) Marshal() {
	SetKind(s)
	s.Value.Marshal()
}

func (s *SpreadNode) String() string {
	return fmt.Sprintf("(spread %v)", s.Value)
}

type VariableAssignmentNode struct {
	NodeTraits
	Variable Node `json:"variable"` // SymbolNode, AttributeAccessNode or IndexAccessNode
	Value    Node `json:"value"`
}

func (*VariableAssignmentNode) KindName() string {
	return "variableAssignment"
}

func (v *VariableAssignmentNode) Marshal() {
	SetKind(v)
	v.Variable.Marshal()
	v.Value.Marshal()
}

func (v *VariableAssignmentNode) String() string {
	return fmt.Sprintf("(var-assignment var: %v, value: %v)", v.Variable, v.Value)
}

type VariableListAssignmentNode struct {
	NodeTraits
	Variables []Node `json:"variables"` // SymbolNode, AttributeAccessNode or IndexAccessNode
	Value     Node   `json:"value"`
}

func (*VariableListAssignmentNode) KindName() string {
	return "variableListAssignment"
}

func (v *VariableListAssignmentNode) Marshal() {
	SetKind(v)
	for _, variable := range v.Variables {
		variable.Marshal()
	}
	v.Value.Marshal()
}

func (v *VariableListAssignmentNode) String() string {
	return fmt.Sprintf("(var-list-assignment var: %v, value: %v)", v.Variables, v.Value)
}

type VariableDefinitionNode struct {
	NodeTraits
	Variable *SymbolNode `json:"variable"`
	Value    Node        `json:"value"`
}

func (*VariableDefinitionNode) KindName() string {
	return "variableDefinition"
}

func (v *VariableDefinitionNode) Marshal() {
	SetKind(v)
	v.Variable.Marshal()
	v.Value.Marshal()
}

func (v *VariableDefinitionNode) String() string {
	return fmt.Sprintf("(var-definition var: %v, value: %v)", v.Variable, v.Value)
}

type VariableListDefinitionNode struct {
	NodeTraits
	Variables []*SymbolNode `json:"variables"`
	Value     Node          `json:"value"`
}

func (*VariableListDefinitionNode) KindName() string {
	return "variableListDefinition"
}

func (v *VariableListDefinitionNode) Marshal() {
	SetKind(v)
	for _, variable := range v.Variables {
		variable.Marshal()
	}
	v.Value.Marshal()
}

func (v *VariableListDefinitionNode) String() string {
	return fmt.Sprintf("(var-list-definition vars: %v, value: %v)", v.Variables, v.Value)
}

type VariableDeclarationNode struct {
	NodeTraits
	Variables []*SymbolNode `json:"variables"`
}

func (*VariableDeclarationNode) KindName() string {
	return "variableDeclaration"
}

func (v *VariableDeclarationNode) Marshal() {
	SetKind(v)
	for _, variable := range v.Variables {
		variable.Marshal()
	}
}

func (v *VariableDeclarationNode) String() string {
	return fmt.Sprintf("(var-declaration vars: %v)", v.Variables)
}
