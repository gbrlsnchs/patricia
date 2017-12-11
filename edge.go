package patricia

// edge is a PATRICIA tree edge.
type edge struct {
	node *Node
}

// newEdge creates a new edge.
func newEdge(node *Node) *edge {
	return &edge{
		node: node,
	}
}
