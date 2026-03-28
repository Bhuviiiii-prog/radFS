package art

func insert(n *Node, value string, key []byte, depth int) *Node {

	if n == nil {
		return newleaf(value, key)
	}
	if isleaf(n) {
		new_node := newNode4()
		oldkey := n.leaf.key
		i := depth

		for i < len(oldkey) && i < len(key) && oldkey[i] == key[i] {
			prefix_index := i - depth
			if prefix_index < maxprefixlen { //index goes till 7 so prefix index<8 and not ==8
				new_node.innerNode.meta.prefix[prefix_index] = key[i] // stores only the till max prefix

			}

			i++
		}

		new_node.innerNode.meta.prefixlen = i - depth // stores full prefix len even after maxprefixlen
		depth = i
		if depth == len(key) {
			new_node.innerNode.leaf = newleaf(value, key)

		} else {
			new_node = addchild(new_node, key[depth], newleaf(value, key))

		}
		if depth == len(oldkey) {
			new_node.innerNode.leaf = n

		} else {
			new_node = addchild(new_node, oldkey[depth], n)

		}

		return new_node

	}
	p := checkprefix(n, key, depth)

	if p != n.innerNode.meta.prefixlen {

		new_node := newNode4()
		if p+depth == len(key) {
			new_node.innerNode.leaf = newleaf(value, key)

		} else {
			new_node = addchild(new_node, key[depth+p], newleaf(value, key))

		}
		leaf := fetchleaf(n) // either its an actual leaf or innernode leaf
		oldkey := leaf.leaf.key

		var oldkeybyte byte
		if p < maxprefixlen {
			oldkeybyte = n.innerNode.meta.prefix[p]
		} else {
			oldkeybyte = oldkey[depth+p]
		}

		if p+depth == len(oldkey) { //when the split is exactly the prefix ends eg intern was already there and you add internship
			new_node.innerNode.leaf = leaf // promote newnode leaf to previous leaf

		} else {
			new_node = addchild(new_node, oldkeybyte, n)
		}

		new_node.innerNode.meta.prefixlen = p
		if p < maxprefixlen {
			copy(new_node.innerNode.meta.prefix, n.innerNode.meta.prefix[:p])

		} else {
			copy(new_node.innerNode.meta.prefix, n.innerNode.meta.prefix[:maxprefixlen])
		}

		oldprefixlen := n.innerNode.meta.prefixlen
		n.innerNode.meta.prefixlen = oldprefixlen - (p + 1)
		if len(n.innerNode.meta.prefix[p+1:oldprefixlen]) < maxprefixlen {
			copy(n.innerNode.meta.prefix, n.innerNode.meta.prefix[p+1:oldprefixlen])

		} else {
			leaf := fetchleaf(n)
			copy(n.innerNode.meta.prefix, leaf.leaf.key[depth+p+1:depth+p+1+maxprefixlen])
		}

		return new_node
	}

	depth += n.innerNode.meta.prefixlen
	next, pos := findchild(key[depth], n)
	if next != nil {
		n.innerNode.children[pos] = insert(next, value, key, depth+1)
		return n

	} else {
		n = addchild(n, key[depth], newleaf(value, key))
		return n

	}

}
