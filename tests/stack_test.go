package tests

import (
	"testing"

	"github.com/michaelorina/go-dsa/latest/sort"
)

func TestStack(t *testing.T) {
	s := &sort.Stack{}

	if !s.IsEmpty() {
		t.Errorf("Expected stack to be empty initially")
	}

	s.Push(10)
	s.Push(20)
	s.Push(30)

	if s.Len() != 3 {
		t.Errorf("Expected length 3, got %d", s.Len())
	}

	val, ok := s.Peek()
	if !ok || val != 30 {
		t.Errorf("Expected Peek to return 30, got %d", val)
	}

	val, ok = s.Pop()
	if !ok || val != 30 {
		t.Errorf("Expected Pop to return 30, got %d", val)
	}

	if s.Len() != 2 {
		t.Errorf("Expected length 2 after Pop, got %d", s.Len())
	}

	_, ok = s.Pop()
	_, ok = s.Pop()

	if !s.IsEmpty() {
		t.Errorf("Expected stack to be empty after all Pops")
	}

	_, ok = s.Pop()
	if ok {
		t.Errorf("Expected Pop on empty stack to return false")
	}
}
