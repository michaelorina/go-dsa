package tests

import (
	"testing"

	"github.com/michaelorina/go-dsa/latest/quicksort"
)

func TestQuickSort(t *testing.T) {
	cases := []struct {
		input		[]int
		expected []int
	}{
		{input: []int{3, 1, 4, 1, 5, 9}, expected: []int{1, 1, 3, 4, 5, 9}},
		{input: []int{5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{10}, expected: []int{10}},
		{input: []int{}, expected: []int{}},
		{input: []int{2, 3, 2, 1}, expected: []int{1, 2, 2, 3}},
	}

	for _, c := range cases {
		arr := append([]int(nil), c.input...)
		quicksort.QuickSort(arr, 0, len(arr)-1)
		for i := range arr {
			if arr[i] != c.expected[i] {
				t.Errorf("QuickSort(%v) = %v; want %v", c.input, arr, c.expected)
				break
			}
		} 
	}
}
