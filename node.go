package patricia

// Node is a node of a PATRICIA tree.
type Node struct {
	// Value is a value of any type held by a node.
	Value interface{}
	edges [uint8(2)]*edge
	depth int
	st    SortingTechnique
}

// Depth returns the node's depth.
func (n *Node) Depth() int {
	return n.depth
}

// IsLeaf returns whether the node is a leaf.
func (n *Node) IsLeaf() bool {
	return len(n.edges) == 0
}

// child returns a child of the node.
func (n *Node) child(v interface{}) *Node {
	c := &Node{
		Value: v,
		depth: n.depth + 1,
	}

	return c
}
