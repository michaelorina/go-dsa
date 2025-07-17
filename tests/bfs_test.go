package tests

import (
	"testing"

	"github.com/michaelorina/go-dsa/latest/treesearch"
)

func buildBFSNodeTree() *treesearch.BFSNode {
	//        1
	//       / \
	//      2   3
	//     / \
	//    4   5
	return &treesearch.BFSNode{
		Value: 1,
		Left: &treesearch.BFSNode{
			Value: 2,
			Left:  &treesearch.BFSNode{Value: 4},
			Right: &treesearch.BFSNode{Value: 5},
		},
		Right: &treesearch.BFSNode{Value: 3},
	}
}

func TestBFS(t *testing.T) {
	root := buildBFSNodeTree()
	expected := []int{1, 2, 3, 4, 5}
	got := treesearch.BFS(root)

	if len(got) != len(expected) {
		t.Fatalf("Expected %v, got %v", expected, got)
	}

	for i := range expected {
		if got[i] != expected[i] {
			t.Errorf("BFS[%d] = %d, want %d", i, got[i], expected[i])
		}
	}
}
