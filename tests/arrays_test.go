package tests

import (
	"testing"

	"github.com/michaelorina/go-dsa/latest/arrays"
)

func TestStaticArray(t *testing.T) {
	expected := [5]int{0, 0, 0, 0, 0}
	result := arrays.StaticArray()

	if result != expected {
		t.Errorf("Expected %v got %v", expected, result)
	}
}

func TestDynamicArray(t *testing.T) {
	expected := []int{}
	result := arrays.DynamicArray()

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}
}
