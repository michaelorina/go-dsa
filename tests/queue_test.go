package tests

import (
	"testing"

	"github.com/michaelorina/go-dsa/latest/sort"
)

func TestQueue(t *testing.T) {
	q := &sort.Queue{}

	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty initially")
	}

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	if q.Len() != 3 {
		t.Errorf("Expected length 3, got %d", q.Len())
	}

	val, ok := q.Peek()
	if !ok || val != 10 {
		t.Errorf("Expected Peek to return 10, got %d", val)
	}

	val, ok = q.Dequeue()
	if !ok || val != 10 {
		t.Errorf("Expected Dequeue to return 10, got %d", val)
	}

	if q.Len() != 2 {
		t.Errorf("Expected length 2 after dequeue, got %d", q.Len())
	}

	q.Dequeue()
	q.Dequeue()

	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty after all dequeues")
	}

	_, ok = q.Dequeue()
	if ok {
		t.Errorf("Expected Dequeue on empty queue to return false")
	}
}
