package tests

import (
	"testing"

	"github.com/michaelorina/go-dsa/latest/doublylinkedlist"
)

func TestDoublyLinkedList(t *testing.T) {
	list := doublylinkedlist.New()

	if !list.IsEmpty() {
		t.Errorf("Expected empty list")
	}

	list.InsertFront(10)
	list.InsertFront(5)
	list.InsertBack(20)

	if list.Len() != 3 {
		t.Errorf("Expected length 3, got %d", list.Len())
	}

	forward := list.ToSlice()
	reverse := list.ToReverseSlice()

	expectedForward := []int{5, 10, 20}
	expectedReverse := []int{20, 10, 5}

	for i := range expectedForward {
		if forward[i] != expectedForward[i] {
			t.Errorf("ToSlice[%d] = %d, want %d", i, forward[i], expectedForward[i])
		}
		if reverse[i] != expectedReverse[i] {
			t.Errorf("ToReverseSlice[%d] = %d, want %d", i, reverse[i], expectedReverse[i])
		}
	}

	if list.Search(10) == nil {
		t.Errorf("Expected to find 10")
	}

	if !list.Delete(10) {
		t.Errorf("Expected to delete 10")
	}

	if list.Search(10) != nil {
		t.Errorf("Expected 10 to be gone")
	}

	if list.Len() != 2 {
		t.Errorf("Expected length 2 after delete, got %d", list.Len())
	}

	if list.Delete(999) {
		t.Errorf("Expected delete to fail on missing value")
	}
}
