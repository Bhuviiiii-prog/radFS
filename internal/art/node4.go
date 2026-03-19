package art

func newNode4() *Node {
	in := &innerNode{
		nodeType: Node4,
		keys:     make([]byte, Node4max),
		children: make([]*Node, Node4max),
		meta: meta{
			prefix: make([]byte, maxprefixlen),
		},
	}
	return &Node{innerNode: in}

}
