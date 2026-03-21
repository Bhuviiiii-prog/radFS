package art

// TODO: Helper functions (e.g., prefix matching)
func addchild(n *Node, k byte, child *Node) {
	in := n.innerNode
	pos := 0

	child1, pos1 := findchild(k, n)
	if child1 != nil {
		in.children[pos1] = child
		return
	}

	for pos < len(in.keys) && in.children[pos] != nil {
		pos++

	}

	var i int
	for i = pos - 1; i >= 0 && in.keys[i] > k; i-- {
		in.keys[i+1] = in.keys[i]
		in.children[i+1] = in.children[i]

	}

	in.keys[i+1] = k
	in.children[i+1] = child

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
	for i := 0; i < len(in.keys); i++ {
		if in.keys[i] == k {
			return in.children[i], i //finds the node and the position
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
