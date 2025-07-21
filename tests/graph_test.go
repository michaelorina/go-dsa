package tests

import (
	"testing"

	"github.com/michaelorina/go-dsa/latest/graph"
)

func TestGraphOperations(t *testing.T) {
	g := graph.NewGraph()

	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)

	if !g.HasEdge(1, 2) || !g.HasEdge(2, 1) {
		t.Errorf("Edge between 1 and 2 should exist")
	}

	if !g.HasEdge(1, 3) {
		t.Errorf("Edge between 1 and 3 should exist")
	}

	if g.HasEdge(3, 4) {
		t.Errorf("Edge between 3 and 4 should not exist")
	}

	neighbors := g.GetNeighbors(1)
	expected := map[int]bool{2: true, 3: true}
	for _, n := range neighbors {
		if !expected[n] {
			t.Errorf("Unexpected neighbor %d of node 1", n)
		}
	}
}
