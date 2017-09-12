package document

// NodeList is a list of parser nodes that implement Node.
type NodeList []Node

func (l *NodeList) Append(n ...Node) {
	for _, node := range n {
		*l = append(*l, node)
	}
}

// last returns the last item added to the slice
func (l *NodeList) LastNode(n ...Node) Node { return (*l)[len(*l)-1] }
