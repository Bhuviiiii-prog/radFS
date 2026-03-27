package art

import "fmt"

func insert(n *Node, value string, key []byte, depth int) *Node {

	if n == nil {
		return newleaf(value, key)
	}
	if isleaf(n) {
		new_node := newNode4()
		oldkey := n.leaf.key
		i := depth
		fmt.Println("depth:", depth)
		fmt.Println("oldkey[depth:]:", string(oldkey[depth:]))
		fmt.Println("key[depth:]:", string(key[depth:]))

		for i < len(oldkey) && i < len(key) && oldkey[i] == key[i] {
			prefix_index := i - depth
			if prefix_index < maxprefixlen { //index goes till 7 so prefix index<8 and not ==8
				new_node.innerNode.meta.prefix[prefix_index] = key[i] // stores only the till max prefix

			}

			i++
		}

		new_node.innerNode.meta.prefixlen = i - depth // stores full prefix len even after maxprefixlen
		depth = i

		new_node = addchild(new_node, keycheck(key, depth), newleaf(value, key))

		new_node = addchild(new_node, keycheck(oldkey, depth), n)
		fmt.Println("depth:", depth)
		fmt.Println("oldkey[depth:]:", string(oldkey[depth:]))
		fmt.Println("key[depth:]:", string(key[depth:]))
		fmt.Println("                             ")
		fmt.Println("prefix", string(new_node.innerNode.meta.prefix))
		return new_node

	}
	p := checkprefix(n, key, depth)
	if p != n.innerNode.meta.prefixlen {
		new_node := newNode4()
		new_node = addchild(new_node, keycheck(key, depth+p), newleaf(value, key))
		if p < maxprefixlen {
			new_node = addchild(new_node, n.innerNode.meta.prefix[p], n)

		} else {
			leaf := fetchleaf(n)
			new_node = addchild(new_node, keycheck(leaf.leaf.key, depth+p), n) // the logic is teh leaf will contain the full key with the same prefix
		}

		new_node.innerNode.meta.prefixlen = p
		if p < maxprefixlen {
			copy(new_node.innerNode.meta.prefix, n.innerNode.meta.prefix[:p])

		} else {
			copy(new_node.innerNode.meta.prefix, n.innerNode.meta.prefix[:maxprefixlen])
		}

		oldprefixlen := n.innerNode.meta.prefixlen
		n.innerNode.meta.prefixlen = n.innerNode.meta.prefixlen - (p + 1)
		if len(n.innerNode.meta.prefix[p+1:oldprefixlen]) < maxprefixlen {
			copy(n.innerNode.meta.prefix, n.innerNode.meta.prefix[p+1:oldprefixlen])

		} else {
			copy(n.innerNode.meta.prefix, n.innerNode.meta.prefix[p+1:maxprefixlen])
		}

		return new_node
	}

	depth += n.innerNode.meta.prefixlen
	next, pos := findchild(keycheck(key, depth), n)
	if next != nil {
		n.innerNode.children[pos] = insert(next, value, key, depth+1)
		return n

	} else {
		n = addchild(n, keycheck(key, depth), newleaf(value, key))
		return n

	}

}
