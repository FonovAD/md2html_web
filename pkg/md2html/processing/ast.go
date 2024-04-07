package processing

type Node struct {
	operator Token
	operand  []*Node
}

type StatmentsNode struct {
	Node
	CodeString []Node
}

func (SN *StatmentsNode) AddNode(node Node) {
	SN.CodeString = append(SN.CodeString, node)
}
