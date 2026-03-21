package art

func search(n *Node, key []byte, depth int) *Node {
	// Base case: nil node means we've reached a dead end.
	// Key does not exist in this path of the tree.
	if n == nil {
		return nil
	}

	// Reached a leaf node, do a full key comparison.
	// Necessary because path compression may have skipped bytes.
	if isleaf(n) {
		if string(n.leaf.key) == string(key) {
			return n
		}
		return nil
	}

	// Check if the compressed prefix at this node matches the search key.
	// If any byte mismatches, the entire subtree is irrelevant.
	if n.innerNode.meta.prefixlen > 0 {
		p := checkprefix(n, key, depth)
		if p != n.innerNode.meta.prefixlen {
			return nil
		}
		depth += n.innerNode.meta.prefixlen
	}

	// Get the next byte to branch on at current depth.
	// Returns 0 (terminator) if key is exhausted.
	k := keycheck(key, depth)

	// Find the child corresponding to byte k and recurse deeper.
	// Return nil if no child exists for this byte.
	next, _ := findchild(k, n)
	if next != nil {
		return search(next, key, depth+1)
	}

	return nil
}
