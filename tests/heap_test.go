package tests

import (
	"testing"

	"github.com/michaelorina/go-dsa/latest/heap"
)

func TestMinHeap_InsertAndPeek(t *testing.T) {
	h := heap.NewMinHeap()
	values := []int{5, 3, 8, 1, 2}

	for _, v := range values {
		h.Insert(v)
	}

	min, ok := h.Peek()
	if !ok || min != 1 {
		t.Errorf("Expected min to be 1, got %d", min)
	}
}

func TestMinHeap_ExtractMin(t *testing.T) {
	h := heap.NewMinHeap()
	values := []int{7, 2, 4, 10, 1}
	for _, v := range values {
		h.Insert(v)
	}

	var extracted []int
	for !h.IsEmpty() {
		val, _ := h.ExtractMin()
		extracted = append(extracted, val)
	}

	expected := []int{1, 2, 4, 7, 10}
	for i := range expected {
		if extracted[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, extracted)
			break
		}
	}
}

func TestMinHeap_Empty(t *testing.T) {
	h := heap.NewMinHeap()

	if _, ok := h.Peek(); ok {
		t.Error("Expected Peek to fail on empty heap")
	}
	if _, ok := h.ExtractMin(); ok {
		t.Error("Expected ExtractMin to fail on empty heap")
	}
	if !h.IsEmpty() {
		t.Error("Expected heap to be empty")
	}
}
