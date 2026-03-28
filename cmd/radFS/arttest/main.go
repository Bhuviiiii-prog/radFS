package main

import (
	"github.com/acmpesuecc/radFS/internal/art"
)

func main() {
	tree := &art.Tree{}

	tree.Insert([]byte("int"), "first")

	tree.Insert([]byte("intern"), "second")
	tree.Insert([]byte("internet"), "first")
	tree.Insert([]byte("i"), "first")

	art.PrintTree(tree.Root(), 0, 0)
}
