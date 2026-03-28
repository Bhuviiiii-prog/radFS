package main

import (
	"github.com/acmpesuecc/radFS/internal/art"
)

func main() {
	t := &art.Tree{}

	t.Insert([]byte("cart"), "v1")
	t.Insert([]byte("car"), "v2")
	t.Insert([]byte("cab"), "v3")
	println("===Tree before deletion===")
	art.PrintTree(t.Root(), 0)
	t.Delete([]byte("car"))
	println("===Tree after deletion of 'car'===")
	art.PrintTree(t.Root(), 0)
	v, ok := t.Search([]byte("cat"))
	println("cat:", v, ok)

	v, ok = t.Search([]byte("car"))
	println("car:", v, ok)

}
