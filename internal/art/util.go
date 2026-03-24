package art

// TODO: Helper functions (e.g., prefix matching)
func addchild(n *Node, k byte, child *Node) {
	in := n.innerNode
	pos := 0
	for pos < len(in.keys) && in.children[pos] != nil {
		pos++

	}
	var i int
	for i = pos - 1; i > 0 && in.keys[i] > k; i-- {
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
		return 0

	} else {
		return key[depth]
	}
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
