package tests

import (
	"testing"

	"github.com/michaelorina/go-dsa/latest/heap"
)

func TestTrie_InsertAndSearch(t *testing.T) {
	trie := heap.NewTrie()

	trie.Insert("go")
	trie.Insert("golang")
	trie.Insert("gone")

	tests := map[string]bool{
		"go":     true,
		"golang": true,
		"gone":   true,
		"g":      false,
		"god":    false,
	}

	for word, expected := range tests {
		got := trie.Search(word)
		if got != expected {
			t.Errorf("Search(%q) = %v, want %v", word, got, expected)
		}
	}
}

func TestTrie_StartsWith(t *testing.T) {
	trie := heap.NewTrie()

	trie.Insert("go")
	trie.Insert("golang")

	tests := map[string]bool{
		"g":      true,
		"go":     true,
		"gol":    true,
		"gola":   true,
		"golang": true,
		"gone":   false,
	}

	for prefix, expected := range tests {
		got := trie.StartsWith(prefix)
		if got != expected {
			t.Errorf("StartsWith(%q) = %v, want %v", prefix, got, expected)
		}
	}
}
