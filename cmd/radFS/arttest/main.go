package main

import (
	"github.com/acmpesuecc/radFS/internal/art"
)

func main() {
	var t art.Tree

	t.Insert([]byte("cart"), "v1")
	t.Insert([]byte("car"), "v2")
	t.Insert([]byte("cab"), "v3")
	v, ok := t.Search([]byte("cat"))
	println("cat:", v, ok)

	v, ok = t.Search([]byte("cab"))
	println("cab:", v, ok)

	v, ok = t.Search([]byte("cart"))
	println("cart:", v, ok)

	art.PrintTree(t.Root(), 0)
}
