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

func insert(n *Node, value string, key []byte, depth int) *Node {

	if n == nil {
		return newleaf(value, key)
	}
	if isleaf(n) {
		new_node := newNode4()
		oldkey := n.leaf.key
		i := depth
		for i < len(oldkey) && i < len(key) && oldkey[i] == key[i] {
			new_node.innerNode.meta.prefix[i-depth] = key[i]
			i++
		}

		new_node.innerNode.meta.prefixlen = i - depth
		depth = i

		addchild(new_node, key[depth], newleaf(value, key))
		addchild(new_node, oldkey[depth], n)
		return new_node

	}
	p := checkprefix(n, key, depth)
	if p != n.innerNode.meta.prefixlen {
		new_node := newNode4()
		addchild(new_node, key[depth+p], newleaf(value, key))
		addchild(new_node, n.innerNode.meta.prefix[p], n)
		new_node.innerNode.meta.prefixlen = p
		copy(new_node.innerNode.meta.prefix, n.innerNode.meta.prefix[:p])

		oldprefixlen := n.innerNode.meta.prefixlen
		n.innerNode.meta.prefixlen = n.innerNode.meta.prefixlen - (p + 1)
		copy(n.innerNode.meta.prefix, n.innerNode.meta.prefix[p+1:oldprefixlen])
		return new_node
	}

	depth += n.innerNode.meta.prefixlen
	next, pos := findchild(key[depth], n)
	if next != nil {
		n.innerNode.children[pos] = insert(next, value, key, depth+1)
		return n

	} else {
		addchild(n, key[depth], newleaf(value, key))
		return n

	}

}

func (t *Tree) Search(key []byte) (string, bool) {
	leaf := search(t.root, key, 0) // start from root and depth 0
	if leaf != nil && isleaf(leaf) {
		return leaf.leaf.values, true //Node->innerleaf->values
	}
	return "", false
}

func search(n *Node, key []byte, depth int) *Node {
	if n == nil {
		return nil
	}

	if isleaf(n) {
		// Verify if the leaf's key actually matches our search key
		if string(n.leaf.key) == string(key) {
			return n
		}
		return nil
	}

	// 1. Check if the node's prefix matches the current part of the key
	if n.innerNode.meta.prefixlen > 0 {
		p := checkprefix(n, key, depth)
		if p != n.innerNode.meta.prefixlen {
			return nil
		}
		depth += n.innerNode.meta.prefixlen
	}

	// 2. Bound check: if we've consumed the prefix but the key is finished, and we aren't at a leaf, the key doesn't exist.
	if depth >= len(key) {
		return nil
	}

	// 3. Find the child corresponding to the byte at the current depth
	next, _ := findchild(key[depth], n)
	if next != nil {
		return search(next, key, depth+1)
	}

	return nil
}
