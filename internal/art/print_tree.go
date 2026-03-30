package art

import "fmt"

func PrintTree(n *Node, level int, depth int) {
	if n == nil {
		return
	}

	indent := ""
	for i := 0; i < level; i++ {
		indent += "  "
	}

	if isleaf(n) {
		fmt.Println(indent + "Leaf: " + string(n.leaf.key))
		return
	}

	in := n.innerNode

	prefixlen := in.meta.prefixlen
	prefix := ""
	if prefixlen <= maxprefixlen {
		prefix = string(in.meta.prefix[:prefixlen])

	} else {
		leaf := fetchleaf(n)
		prefix = string(leaf.leaf.key[depth : depth+prefixlen])
	}

	fmt.Println(indent+"Node(prefix=\""+prefix+"\", prefixLen=", prefixlen, ")")

	if in.leaf != nil {
		fmt.Printf("%s  [Internal Leaf]: %s\n", indent, string(in.leaf.leaf.key))
	}
	// Print children

	newDepth := depth + prefixlen
	for i := 0; i < len(in.keys); i++ {
		if in.children[i] != nil {
			fmt.Printf("%s Edge('%c' | %d):\t", indent, in.keys[i], in.keys[i])

			PrintTree(in.children[i], level+1, newDepth+1)
		}
	}
}
