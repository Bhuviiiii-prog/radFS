package main

import (
	"fmt"

	"github.com/acmpesuecc/radFS/internal/art"
)

func main() {
	tree := &art.Tree{}

	for i := 0; i < 50; i++ {
		key := []byte("ca" + string(byte(i+97)))
		tree.Insert(key, fmt.Sprintf("val%d", i))
	}
	art.PrintTree(tree.Root(), 0, 0)
}
