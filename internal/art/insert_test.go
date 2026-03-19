package art

import "testing"

func TestInsertStructure(t *testing.T) {
	var tree Tree

	tree.Insert([]byte("cat"), "v1")
	tree.Insert([]byte("car"), "v2")
	tree.Insert([]byte("cap"), "v3")

	root := tree.root

	if root == nil {
		t.Fatal("root nil")
	}

	if isleaf(root) {
		t.Fatal("root should not be leaf")
	}

	in := root.innerNode

	// check prefix
	prefix := string(in.meta.prefix[:in.meta.prefixlen])
	if prefix != "ca" {
		t.Fatalf("expected prefix 'ca', got '%s'", prefix)
	}

	// check children count
	count := 0
	for _, c := range in.children {
		if c != nil {
			count++
		}
	}

	if count != 3 {
		t.Fatalf("expected 3 children, got %d", count)
	}
}
