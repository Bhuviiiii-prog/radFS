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

		addchild(new_node, keycheck(key, depth), newleaf(value, key))
		addchild(new_node, keycheck(oldkey, depth), n)
		return new_node

	}
	p := checkprefix(n, key, depth)
	if p != n.innerNode.meta.prefixlen {
		new_node := newNode4()
		addchild(new_node, keycheck(key, depth+p), newleaf(value, key))
		addchild(new_node, n.innerNode.meta.prefix[p], n)
		new_node.innerNode.meta.prefixlen = p
		copy(new_node.innerNode.meta.prefix, n.innerNode.meta.prefix[:p])

		oldprefixlen := n.innerNode.meta.prefixlen
		n.innerNode.meta.prefixlen = n.innerNode.meta.prefixlen - (p + 1)
		copy(n.innerNode.meta.prefix, n.innerNode.meta.prefix[p+1:oldprefixlen])
		return new_node
	}

	depth += n.innerNode.meta.prefixlen
	next, pos := findchild(keycheck(key, depth), n)
	if next != nil {
		n.innerNode.children[pos] = insert(next, value, key, depth+1)
		return n

	} else {
		addchild(n, keycheck(key, depth), newleaf(value, key))
		return n

	}

}
