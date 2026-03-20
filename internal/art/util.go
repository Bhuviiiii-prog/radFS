package art

// TODO: Helper functions (e.g., prefix matching)
func addchild(n *Node, k byte, child *Node) {
	for i := 0; i < len(n.innerNode.keys); i++ {
		if n.innerNode.children[i] == nil {
			n.innerNode.children[i] = child
			n.innerNode.keys[i] = k
			return

		}

	}

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
		return 0

	} else {
		return key[depth]
	}
}
