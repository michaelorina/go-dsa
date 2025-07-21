package graph

type Graph struct {
	adj map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		adj: make(map[int][]int),
	}
}

func (g *Graph) AddEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u) // undirected
}

func (g *Graph) GetNeighbors(u int) []int {
	return g.adj[u]
}

func (g *Graph) HasEdge(u, v int) bool {
	for _, neighbor := range g.adj[u] {
		if neighbor == v {
			return true
		}
	}
	return false
}

func (g *Graph) Nodes() []int {
	nodes := make([]int, 0, len(g.adj))
	for node := range g.adj {
		nodes = append(nodes, node)
	}
	return nodes
}
