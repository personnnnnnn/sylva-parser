package transformer

type NodeTraits struct {
	Kind   string        `json:"kind"`
	Bounds Bounds        `json:"bounds"`
	Errors []*SylvaError `json:"errors"`
}

func (n *NodeTraits) SetKind(newKind string) {
	n.Kind = newKind
}

func (NodeTraits) KindName() string {
	return "UNKNOWN"
}

func (n *NodeTraits) Marshal() {
	SetKind(n)
}

type Node interface {
	GetBounds() Bounds
	SetBounds(bounds Bounds)
	GetErrors() []*SylvaError
	AppendErrors(other Node)
	AppendError(err *SylvaError)
	SetErrors(errors []*SylvaError)
	KindName() string
	SetKind(newKind string)
	Marshal()
}

func SetKind(node Node) {
	node.SetKind(node.KindName())
	if node.GetErrors() == nil {
		node.SetErrors([]*SylvaError{})
	}
}

func (n *NodeTraits) SetBounds(bounds Bounds) {
	n.Bounds = bounds
}

func (n *NodeTraits) SetErrors(errors []*SylvaError) {
	n.Errors = errors
}

func (n *NodeTraits) AppendErrors(other Node) {
	n.Errors = append(n.Errors, other.GetErrors()...)
}

func (n *NodeTraits) AppendError(err *SylvaError) {
	n.Errors = append(n.Errors, err)
}

func (n *NodeTraits) GetBounds() Bounds {
	return n.Bounds
}

func (n *NodeTraits) GetErrors() []*SylvaError {
	return n.Errors
}
