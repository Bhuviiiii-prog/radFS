package main

import (
	"fmt"

	"github.com/acmpesuecc/radFS/internal/art"
)

func main() {
	t := &art.Tree{}

<<<<<<< HEAD
	for i := 0; i < 50; i++ {
		key := []byte("ca" + string(byte(i+97)))
		tree.Insert(key, fmt.Sprintf("val%d", i))
	}
	art.PrintTree(tree.Root(), 0, 0)
=======

	t.Insert([]byte("cart"), "v1")
	t.Insert([]byte("car"), "v2")
	t.Insert([]byte("cab"), "v3")
	println("===Tree before deletion===")
	art.PrintTree(t.Root(),0 ,0)
	t.Delete([]byte("car"))
	println("===Tree after deletion of 'car'===")
	art.PrintTree(t.Root(), 0, 0)
	v, ok := t.Search([]byte("cat"))
	println("cat:", v, ok)

	v, ok = t.Search([]byte("car"))
	println("car:", v, ok)


>>>>>>> ffc8e5cd38b2330448f25c992ff8ea98c1c1c6f5
}
