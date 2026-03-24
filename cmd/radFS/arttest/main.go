package main

import (
	"fmt"

	"github.com/acmpesuecc/radFS/internal/art"
)

func main() {
	tree := &art.Tree{}

	tree.Insert([]byte{0x0A}, "first")
	tree.Insert([]byte{0x0A, 0x01}, "second")

	val, found := tree.Search([]byte{0x0A})
	if !found || val != "first" {

		fmt.Printf("Expected 'first', got %s\n", val)
	}

	val2, found2 := tree.Search([]byte{0x0A, 0x01})
	if !found2 || val2 != "second" {
		fmt.Printf("Expected 'second', got %s\n", val2)
	}

	//art.PrintTree(tree.Root(), 0)
}
