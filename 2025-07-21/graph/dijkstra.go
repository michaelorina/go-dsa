package graph

import (
	"container/heap"
	"math"
)

type WeightedGraph struct {
	adj map[int][]Edge
}

type Edge struct {
	To     int
	Weight int
}

func NewWeightedGraph() *WeightedGraph {
	return &WeightedGraph{
		adj: make(map[int][]Edge),
	}
}

func (g *WeightedGraph) AddEdge(u, v, weight int) {
	g.adj[u] = append(g.adj[u], Edge{To: v, Weight: weight})
	g.adj[v] = append(g.adj[v], Edge{To: u, Weight: weight}) // remove for directed
}

func (g *WeightedGraph) Dijkstra(start int) map[int]int {
	dist := make(map[int]int)
	for node := range g.adj {
		dist[node] = math.MaxInt
	}
	dist[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{Node: start, Priority: 0})

	visited := make(map[int]bool)

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*Item)
		if visited[curr.Node] {
			continue
		}
		visited[curr.Node] = true

		for _, edge := range g.adj[curr.Node] {
			if dist[edge.To] > dist[curr.Node]+edge.Weight {
				dist[edge.To] = dist[curr.Node] + edge.Weight
				heap.Push(pq, &Item{Node: edge.To, Priority: dist[edge.To]})
			}
		}
	}

	return dist
}
