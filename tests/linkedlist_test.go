package tests

import (
	"testing"

	"github.com/michaelorina/go-dsa/latest/sort"
)

func TestLinkedList(t *testing.T) {
	list := &linkedlist.LinkedList{}

	list.Insert(10)
	list.Insert(20)
	list.Insert(30)

	if !list.Search(20) {
		t.Errorf("Expected to find 20 in list")
	}

	if list.Length() != 3 {
		t.Errorf("Expected length 3, got %d", list.Length())
	}

	if !list.Delete(20) {
		t.Errorf("Expected to delete 20 from list")
	}

	if list.Search(20) {
		t.Errorf("Did not expect to find 20 after deletion")
	}

	if list.Length() != 2 {
		t.Errorf("Expected length 2 after deletion, got %d", list.Length())
	}
}
