package tests

import (
  "testing"

  "github.com/michaelorina/go-dsa/latest/search"
)

func TestLinearSearch(t *testing.T) {
  tests := []struct {
    name string
    input []int
    target int
    expected int
    found bool
  }{
    {"Found in the middle", []int{3, 5, 7, 9, 11}, 7, 2, true}
    {"Found at start", []int{4, 6, 8}, 4, 0, true}
    {"Found at end", []int{1, 3, 5}, 5, 2, true}
    {"Not found", []int{1, 2, 3}, 10, -1, false}
    {"Empty slice", []int{}, 1, -1, false}
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      index, found := search.LinearSearch(tt.input, tt.target)
      if index != tt.expected || found != tt.found {
        t.Errorf("Got (%d, %v), Expected (%d, %v)", index, found, tt.expected, tt.found)
      }
    })
  }
}


