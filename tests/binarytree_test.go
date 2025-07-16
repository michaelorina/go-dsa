package tests

import (
	"testing"

	"github.com/michaelorina/go-dsa/latest/tree"
)

func buildTestTree() *tree.TreeNode {
	//        1
	//       / \
	//      2   3
	//     / \
	//    4   5
	return &tree.TreeNode{
		Value: 1,
		Left: &tree.TreeNode{
			Value: 2,
			Left:  &tree.TreeNode{Value: 4},
			Right: &tree.TreeNode{Value: 5},
		},
		Right: &tree.TreeNode{Value: 3},
	}
}

func TestPreOrder(t *testing.T) {
	root := buildTestTree()
	expected := []int{1, 2, 4, 5, 3}
	got := tree.PreOrder(root)

	if len(got) != len(expected) {
		t.Fatalf("Expected %v, got %v", expected, got)
	}

	for i := range expected {
		if got[i] != expected[i] {
			t.Errorf("PreOrder[%d] = %d, want %d", i, got[i], expected[i])
		}
	}
}

func TestInOrder(t *testing.T) {
	root := buildTestTree()
	expected := []int{4, 2, 5, 1, 3}
	got := tree.InOrder(root)

	if len(got) != len(expected) {
		t.Fatalf("Expected %v, got %v", expected, got)
	}

	for i := range expected {
		if got[i] != expected[i] {
			t.Errorf("InOrder[%d] = %d, want %d", i, got[i], expected[i])
		}
	}
}

func TestPostOrder(t *testing.T) {
	root := buildTestTree()
	expected := []int{4, 5, 2, 3, 1}
	got := tree.PostOrder(root)

	if len(got) != len(expected) {
		t.Fatalf("Expected %v, got %v", expected, got)
	}

	for i := range expected {
		if got[i] != expected[i] {
			t.Errorf("PostOrder[%d] = %d, want %d", i, got[i], expected[i])
		}
	}
}
