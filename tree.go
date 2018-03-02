package patricia

import "sync"

// Tree is a PATRICIA Tree.
type Tree struct {
	// Safe tells whether the tree's operations
	// should be thread-safe. By default, the tree's
	// not thread-safe.
	Safe bool
	// name is the tree's name.
	name string
	// root is the tree's root.
	root *Node
	// size is the total number of nodes
	// without the tree's root.
	size uint
	// mtx controls the operations' safety.
	mtx *sync.RWMutex
}

// New creates a named PATRICIA tree with a single node (its root).
func New(name string) *Tree {
	return &Tree{name: name, root: &Node{}, mtx: &sync.RWMutex{}}
}

// Add adds a new node to the tree.
func (t *Tree) Add(s string, v interface{}) {
	if t.Safe {
		defer t.mtx.Unlock()
		t.mtx.Lock()
	}

	if v == nil {
		return
	}

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
func (t *Tree) Del(s string) {
	if t.Safe {
		defer t.mtx.Unlock()
		t.mtx.Lock()
	}

	tnode := t.root

	if tnode.IsLeaf() {
		return
	}

	leaf := tnode
	count := uint(0)

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

			tnode = tnode.edges[bit].node
			count++

			if i == len(s)-1 && j-1 == 0 && tnode.IsLeaf() {
				leaf.edges = [uint8(2)]*edge{}
				t.size -= count

				return
			}

			if tnode != nil && tnode.Value != nil {
				leaf = tnode
				count = 0
			}
		}
	}
}

// Get retrieves a node.
func (t *Tree) Get(s string) *Node {
	if t.Safe {
		defer t.mtx.RUnlock()
		t.mtx.RLock()
	}

	tnode := t.root

	if tnode.IsLeaf() {
		return nil
	}

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
	if t.Safe {
		defer t.mtx.RUnlock()
		t.mtx.RLock()
	}

	return t.size + 1
}
