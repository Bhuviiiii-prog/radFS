package main

import (
	"github.com/acmpesuecc/radFS/internal/art"
)

func main() {
	tree := &art.Tree{}

	tree.Insert([]byte("cab"), "first")

	tree.Insert([]byte("can"), "second")
	tree.Insert([]byte("car"), "first")
	tree.Insert([]byte("cat"), "first")

	tree.Insert([]byte("caz"), "second")

	art.PrintTree(tree.Root(), 0, 0)
}
