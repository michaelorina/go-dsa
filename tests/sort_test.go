package tests

import (
	"reflect"
	"testing"

	"github.com/michaelorina/go-dsa/latest/sort"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Already sorted", []int{1, 2, 3}, []int{1, 2, 3}},
		{"Reverse order", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"Random order", []int{3, 1, 4, 2}, []int{1, 2, 3, 4}},
		{"All same", []int{7, 7, 7}, []int{7, 7, 7}},
		{"Single element", []int{1}, []int{1}},
		{"Empty slice", nil, nil}, // <- use nil for both to avoid reflect.DeepEqual mismatch
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy so original input remains unchanged
			inputCopy := append([]int(nil), tt.input...)
			sort.BubbleSort(inputCopy)

			// Treat nil and []int{} as equivalent for comparison
			if len(inputCopy) == 0 && len(tt.expected) == 0 {
				return
			}

			if !reflect.DeepEqual(inputCopy, tt.expected) {
				t.Errorf("got %v, want %v", inputCopy, tt.expected)
			}
		})
	}
}
