package art

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
	k := keycheck(key, depth)

	// 3. Find the child corresponding to the byte at the current depth
	next, _ := findchild(k, n)
	if next != nil {
		return search(next, key, depth+1)
	}

	return nil
}
