package art

// TODO: Public API (Tree struct, Insert, Search, Delete)

type Tree struct {
	root *Node
}

func (t *Tree) Root() *Node {
	return t.root
}

func (t *Tree) Insert(key []byte, value string) {
	t.root = insert(t.root, value, key, 0)
}

func (t *Tree) Search(key []byte) (string, bool) {
	leaf := search(t.root, key, 0) // start from root and depth 0
	if leaf != nil && isleaf(leaf) {
		return leaf.leaf.values, true //Node->innerleaf->values
	}
	return "", false
}
func (t *Tree) Delete(key []byte) bool {
	if t.root == nil {
		return false
	}

	newRoot, deleted := deletekey(t.root, key, 0)

	if deleted {
		t.root = newRoot
		return true
	}

	return false
}
