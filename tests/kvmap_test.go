package tests

import (
	"testing"

	"github.com/michaelorina/go-dsa/latest/kvmap"
)

func TestMap(t *testing.T) {
	m := kvmap.New[string, int]()

	m.Set("age", 30)
	m.Set("score", 100)

	if val, ok := m.Get("age"); !ok || val != 30 {
		t.Errorf("Expected 30, got %v", val)
	}

	if !m.Has("score") {
		t.Errorf("Expected key 'score' to exist")
	}

	if m.Len() != 2 {
		t.Errorf("Expected length 2, got %d", m.Len())
	}

	m.Delete("score")
	if m.Has("score") {
		t.Errorf("Expected 'score' to be deleted")
	}

	keys := m.Keys()
	if len(keys) != 1 || keys[0] != "age" {
		t.Errorf("Expected keys to contain only 'age', got %v", keys)
	}
}
