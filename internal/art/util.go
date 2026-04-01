package art

// TODO: Helper functions (e.g., prefix matching)
func addchild(n *Node, k byte, child *Node) *Node {
	in := n.innerNode
	child1, pos1 := findchild(k, n)
	if child1 != nil {
		in.children[pos1] = child
		return n
	}

	if n.innerNode.num_children == len(in.keys) {
		n = grow(n)
		in = n.innerNode

	}
	var i int
	for i = in.num_children - 1; i >= 0 && in.keys[i] > k; i-- {
		in.keys[i+1] = in.keys[i]
		in.children[i+1] = in.children[i]

	}

	in.keys[i+1] = k
	in.children[i+1] = child
	in.num_children += 1

	return n

}
func checkprefix(n *Node, key []byte, depth int) int {
	in := n.innerNode
	var i int
	maxcmp := min(maxprefixlen, in.meta.prefixlen)

	for i = 0; i < maxcmp && depth+i < len(key); i++ { //checks prefix until mismatch
		if in.meta.prefix[i] != key[depth+i] {
			return i // case when you find mismatch and the mismatch is less than maxprefixlen

		}

	}
	if in.meta.prefixlen > maxprefixlen {
		leaf := fetchleaf(n)
		leafkey := leaf.leaf.key
		for ; i < in.meta.prefixlen && depth+i < len(leafkey) && depth+i < len(key); i++ {
			if key[depth+i] != leaf.leaf.key[depth+i] {
				return i // case when you find mismatch and the mismatch is more than maxprefixlen

			}

		}

	}

	return i // case when you find mismatch and the mismatch is equal maxprefixlen

}

func findchild(k byte, n *Node) (*Node, int) {
	in := n.innerNode
	switch in.nodeType {
	case Node4, Node16:
		for i := 0; i < in.num_children; i++ {
			if in.keys[i] == k {
				return in.children[i], i //finds the node and the position
			}

		}
	case Node48:
		idx := in.keys[k]
		if idx > 0 {
			realindex := int(idx - 1)
			return in.children[realindex], realindex

		}
	case Node256:
		if in.children[k] != nil {
			return in.children[k], int(k)
		}

	}
	return nil, -1

}

// removechild removes the child with key k from node n by shifting
// all subsequent keys and children left to fill the gap.
func removechild(n *Node, k byte) {
	_, pos := findchild(k, n)
	if pos == -1 {
		return // key not found, nothing to remove
	}

	in := n.innerNode
	last := len(in.keys) - 1

	// shift everything after pos one step to the left
	for i := pos; i < last; i++ {
		in.keys[i] = in.keys[i+1]
		in.children[i] = in.children[i+1]
	}

	// clear the now-duplicate last slot to avoid stale pointers
	in.keys[last] = 0
	in.children[last] = nil
}

func grow(n *Node) *Node {
	switch n.innerNode.nodeType {
	case Node4:
		n16 := newNode16()
		copymeta(n, n16)
		index := 0
		for i := 0; i < 4; i++ {
			if n.innerNode.children[i] != nil {
				n16.innerNode.keys[index] = n.innerNode.keys[i]
				n16.innerNode.children[index] = n.innerNode.children[i]
				index++

			}

		}
		n16.innerNode.num_children = index
		return n16
	case Node16:
		n48 := newNode48()
		copymeta(n, n48)
		index := 0
		for i := 0; i < 16; i++ {
			idx := n.innerNode.keys[i]
			child := n.innerNode.children[i]

			if child != nil {
				n48.innerNode.keys[idx] = byte(index + 1)
				n48.innerNode.children[index] = child
				index++

			}

		}
		n48.innerNode.num_children = index
		return n48

	case Node48:
		n256 := newNode256()
		copymeta(n, n256)
		for i := 0; i < 256; i++ {
			idx := n.innerNode.keys[i]

			if n.innerNode.keys[i] != 0 {
				child := n.innerNode.children[int(idx-1)]
				n256.innerNode.children[i] = child
			}

		}

		return n256

	}
	return nil

}
func copymeta(n *Node, new_node *Node) {

	new_node.innerNode.meta.prefixlen = n.innerNode.meta.prefixlen

	limit := n.innerNode.meta.prefixlen
	if limit > maxprefixlen {
		limit = maxprefixlen
	}

	copy(new_node.innerNode.meta.prefix, n.innerNode.meta.prefix[:limit])
}

func fetchleaf(n *Node) *Node {
	if isleaf(n) {
		return n
	}
	if n.innerNode.leaf != nil {
		return n.innerNode.leaf
	}

	for i := 0; i < len(n.innerNode.keys); i++ {
		if n.innerNode.children[i] != nil {
			return fetchleaf(n.innerNode.children[i])

		}

	}
	return nil

}
