package art

// deletekey recursively walks the tree to find and remove the given key.
// Returns the (possibly modified) node and whether the key was deleted.
func deletekey(n *Node, key []byte, depth int) (*Node, bool) {
	if n == nil {
		return nil, false
	}

	if isleaf(n) {
		if string(n.leaf.key) == string(key) {
			return nil, true
		}
		return n, false
	}

	p := checkprefix(n, key, depth)
	if p != n.innerNode.meta.prefixlen {
		return n, false
	}
	depth += n.innerNode.meta.prefixlen

	k := keycheck(key, depth)
	child, pos := findchild(k, n)
	if child == nil {
		return n, false
	}

	newChild, deleted := deletekey(child, key, depth+1)
	if !deleted {
		return n, false
	}

	if newChild == nil {
		removechild(n, k)
	} else {
		n.innerNode.children[pos] = newChild
	}

	return n, true
}
