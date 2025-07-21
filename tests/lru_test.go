package tests

import (
	"testing"

	"github.com/michaelorina/go-dsa/latest/kvmap"
)

func TestLRUCache(t *testing.T) {
	cache := kvmap.NewLRU(2)

	cache.Put("a", 1)
	cache.Put("b", 2)

	if val, ok := cache.Get("a"); !ok || val != 1 {
		t.Errorf("Expected Get('a') = 1, got %v", val)
	}

	// This should evict key "b" because "a" was recently accessed
	cache.Put("c", 3)

	if _, ok := cache.Get("b"); ok {
		t.Error("Expected 'b' to be evicted")
	}

	if val, ok := cache.Get("a"); !ok || val != 1 {
		t.Errorf("Expected Get('a') = 1, got %v", val)
	}

	if val, ok := cache.Get("c"); !ok || val != 3 {
		t.Errorf("Expected Get('c') = 3, got %v", val)
	}
}
