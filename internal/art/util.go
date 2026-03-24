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
	for i = 0; i < in.meta.prefixlen && in.meta.prefix[i] == keycheck(key, depth+i); i++ { //checks prefix until mismatch

	}
	return i

}
func findchild(k byte, n *Node) (*Node, int) {
	in := n.innerNode
	switch in.nodeType {
	case Node4, Node16:
		for i := 0; i < len(in.keys); i++ {
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

func keycheck(key []byte, depth int) byte {
	if depth >= len(key) {
		return 1

	} else {
		return key[depth]
	}
}

func grow(n *Node) *Node {
	switch n.innerNode.nodeType {
	case Node4:
		n16 := newNode16()
		copymeta(n, n16)
		for i := 0; i < 4; i++ {
			n16.innerNode.keys[i] = n.innerNode.keys[i]
			n16.innerNode.children[i] = n.innerNode.children[i]
		}
		return n16
	case Node16:
		n48 := newNode48()
		copymeta(n, n48)
		index := 0
		for i := 0; i < 16; i++ {
			idx := n.innerNode.keys[i]
			child := n.innerNode.children[i]

			if child != nil {
				n48.innerNode.keys[idx] = byte(index + 1) // the reason its index+1 is because we are making 0 a kind of "no children" case

				n48.innerNode.children[index] = child
				index++

			}

		}
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
	new_node.innerNode.meta.prefix = n.innerNode.meta.prefix
	new_node.innerNode.meta.prefixlen = n.innerNode.meta.prefixlen

}
