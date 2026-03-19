package art

// TODO: Public API (Tree struct, Insert, Search, Delete)

type Tree struct {
	root *Node
}

func (t *Tree) Insert(key []byte, value string) {
	t.root = insert(t.root, value, key, 0)
}

func (t *Tree) Root() *Node {
	return t.root
}
