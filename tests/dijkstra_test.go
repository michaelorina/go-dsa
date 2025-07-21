package tests

import (
	"testing"

	"github.com/michaelorina/go-dsa/latest/graph"
)

func TestDijkstra(t *testing.T) {
	g := graph.NewWeightedGraph()

	g.AddEdge(0, 1, 4)
	g.AddEdge(0, 2, 1)
	g.AddEdge(2, 1, 2)
	g.AddEdge(1, 3, 1)
	g.AddEdge(2, 3, 5)

	dist := g.Dijkstra(0)

	expected := map[int]int{
		0: 0,
		1: 3,
		2: 1,
		3: 4,
	}

	for node, exp := range expected {
		got := dist[node]
		if got != exp {
			t.Errorf("Node %d: expected %d, got %d", node, exp, got)
		}
	}
}
