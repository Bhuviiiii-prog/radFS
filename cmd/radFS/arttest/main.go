package main

import (
	"github.com/acmpesuecc/radFS/internal/art"
)

func main() {
	var t art.Tree

	t.Insert([]byte("cat"), "v1")
	t.Insert([]byte("car"), "v2")
	t.Insert([]byte("cap"), "v3")

	art.PrintTree(t.Root(), 0)
}
