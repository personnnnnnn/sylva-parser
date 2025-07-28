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
	return "arg-list"
}

func (a *ArgListNode) String() string {
	return fmt.Sprintf("(arglist: %v)", a.Arguments)
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
	return "attr-access"
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
	return "attr-access"
}

func (i *IndexAccessNode) Marshal() {
	SetKind(i)
	i.Value.Marshal()
	i.Index.Marshal()
}

func (i *IndexAccessNode) String() string {
	return fmt.Sprintf("(index-access value: %v, index: %v)", i.Value, i.Index)
}
