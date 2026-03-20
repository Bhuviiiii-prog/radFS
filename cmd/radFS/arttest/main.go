package main

import (
	"github.com/acmpesuecc/radFS/internal/art"
)

func main() {
	var t art.Tree

	t.Insert([]byte("cats"), "v1")

	t.Insert([]byte("cat"), "v3")
	t.Insert([]byte("carpet"), "v3")

	art.PrintTree(t.Root(), 0)
}
