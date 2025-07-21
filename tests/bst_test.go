package tests

import (
	"testing"

	"github.com/michaelorina/go-dsa/latest/treesearch"
)

func TestBSTInsertAndSearch(t *testing.T) {
	bst := &treesearch.BST{}
	values := []int{8, 3, 10, 1, 6, 14}

	for _, v := range values {
		bst.Insert(v)
	}

	for _, v := range values {
		if !bst.Search(v) {
			t.Errorf("Expected %d to be found in the BST", v)
		}
	}

	if bst.Search(99) {
		t.Errorf("Expected 99 not to be found in the BST")
	}
}

func TestBSTInOrderTraversal(t *testing.T) {
	bst := &treesearch.BST{}
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, v := range values {
		bst.Insert(v)
	}

	expected := []int{2, 3, 4, 5, 6, 7, 8}
	got := bst.InOrderTraversal()

	if len(got) != len(expected) {
		t.Fatalf("Expected %v, got %v", expected, got)
	}

	for i := range expected {
		if got[i] != expected[i] {
			t.Errorf("Traversal[%d] = %d, want %d", i, got[i], expected[i])
		}
	}
}
