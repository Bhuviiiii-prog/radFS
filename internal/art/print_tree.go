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

	prefixLen := in.meta.prefixlen
	if prefixLen < 0 || prefixLen > len(in.meta.prefix) {
		prefixLen = 0
	}

	prefix := string(in.meta.prefix[:prefixLen])

	fmt.Println(indent+"Node(prefix=\""+prefix+"\", prefixLen=", prefixLen, ")")

	// Print children
	for i := 0; i < len(in.keys); i++ {
		if in.children[i] != nil {
			fmt.Println(indent+" Edge('", string(in.keys[i]), "'):")
			PrintTree(in.children[i], level+1)
		}
	}
}
