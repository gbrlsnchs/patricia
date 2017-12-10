package patricia

// Tree is a PATRICIA Tree.
type Tree struct {
	// name is the tree's name.
	name string
	// root is the tree's root.
	root *Node
	// size is the total number of nodes
	// without the tree's root.
	size uint
}

// New creates a named PATRICIA tree with a single node (its root).
func New(name string) *Tree {
	return &Tree{name: name, root: &Node{}}
}

// Add adds a new node to the tree.
func (t *Tree) Add(s string, v interface{}) {
	tnode := t.root

walk:
	for i := 0; i < len(s); i++ {
		for j := uint8(8); j > 0; j-- {
			exp := byte(1 << (j - 1))
			mask := s[i] & exp
			bit := uint8(0)

			if mask > 0 {
				bit = 1
			}

			if tnode.edges[bit] == nil {
				if i == len(s)-1 && j-1 == 0 {
					tnode.edges[bit] = newEdge(tnode.child(v))
					t.size++

					break walk
				}

				tnode.edges[bit] = newEdge(tnode.child(nil))
				tnode = tnode.edges[bit].node
				t.size++

				continue
			}

			if i == len(s)-1 && j == 0 {
				tnode.edges[bit].node.Value = v

				break walk
			}

			tnode = tnode.edges[bit].node
		}
	}
}

// Del deletes a node.
//
// If a parent node that holds no value ends up holding only one edge
// after a deletion of one of its edges, it gets merged with the remaining edge.
func (t *Tree) Del(s string) {
	tnode := t.root

	for i := 0; i < len(s); i++ {
		for j := uint8(8); j > 0; j-- {
			exp := byte(1 << (j - 1))
			mask := s[i] & exp
			bit := uint8(0)

			if mask > 0 {
				bit = 1
			}

			if tnode.edges[bit] == nil {
				return
			}

			if i == len(s)-1 && j-1 == 0 {
				// TODO: father's grandchildren become its children

				return
			}

			tnode = tnode.edges[bit].node
		}
	}
}

// Get retrieves a node.
func (t *Tree) Get(s string) *Node {
	tnode := t.root

	for i := 0; i < len(s); i++ {
		for j := uint8(8); j > 0; j-- {
			exp := byte(1 << (j - 1))
			mask := s[i] & exp
			bit := uint8(0)

			if mask > 0 {
				bit = 1
			}

			if tnode.edges[bit] == nil {
				return nil
			}

			if i == len(s)-1 && j-1 == 0 {
				return tnode.edges[bit].node
			}

			tnode = tnode.edges[bit].node
		}
	}

	return nil
}

// Size returns the total numbers of nodes the tree has,
// including the root.
func (t *Tree) Size() uint {
	return t.size + 1
}
