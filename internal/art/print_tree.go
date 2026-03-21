package art

import "fmt"

func PrintTree(n *Node, level int) {
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
	if prefixlen < 0 || prefixlen > len(in.meta.prefix) {
		prefixlen = 0
	}

	prefix := string(in.meta.prefix[:prefixlen])

	fmt.Println(indent+"Node(prefix=\""+prefix+"\", prefixLen=", prefixlen, ")")

	// Print children
	for i := 0; i < len(in.keys); i++ {
		if in.children[i] != nil {
			fmt.Printf("%s Edge('%c' | %d):\t", indent, in.keys[i], in.keys[i])

			PrintTree(in.children[i], level+1)
		}
	}
}
