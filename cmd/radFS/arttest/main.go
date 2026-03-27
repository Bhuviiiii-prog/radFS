package main

import (
	"github.com/acmpesuecc/radFS/internal/art"
)

func main() {
	tree := &art.Tree{}

	tree.Insert([]byte("abbbbbbbc"), "first")
	tree.Insert([]byte("abbbbbbbcb"), "second")
	tree.Insert([]byte("abbbbbbbcbc"), "second")

	art.PrintTree(tree.Root(), 0)
}
