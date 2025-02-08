package main

import "fmt"

// Edge represents a connection between two nodes
type Edge struct {
	Source      int
	Destination int
	Weight      int
}

// Graph represents a graph with a list of edges
type Graph struct {
	vertices int
	edges    []Edge
}

const large = 999999

// NewGraph creates a new graph with a given number of vertices
func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		edges:    make([]Edge, 0),
	}
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(source, destination, weight int) {
	g.edges = append(g.edges, Edge{source, destination, weight})
	g.edges = append(g.edges, Edge{destination, source, weight})
}

// Dijkstra calculates the shortest path from a source node to all other nodes
func (g *Graph) Dijkstra(source int) []int {
	distances := make([]int, g.vertices)
	visited := make([]bool, g.vertices)

	for i := range distances {
		distances[i] = large
	}
	distances[source] = 0

	for i := 0; i < g.vertices-1; i++ {
		u := g.minDistance(distances, visited)
		visited[u] = true

		for _, edge := range g.edges {
			if !visited[edge.Destination] && edge.Source == u {
				newDistance := distances[u] + edge.Weight
				if newDistance < distances[edge.Destination] {
					distances[edge.Destination] = newDistance
				}
			}
		}
	}

	return distances
}

func (g *Graph) minDistance(distances []int, visited []bool) int {
	minDist := large
	minIndex := -1

	for v := 0; v < g.vertices; v++ {
		if !visited[v] && distances[v] <= minDist {
			minDist = distances[v]
			minIndex = v
		}
	}

	return minIndex
}

func main() {
	g := NewGraph(9)

	// Add edges to the graph
	g.AddEdge(0, 1, 4)
	g.AddEdge(0, 7, 8)
	g.AddEdge(1, 2, 8)
	g.AddEdge(1, 7, 11)
	g.AddEdge(2, 3, 7)
	g.AddEdge(2, 8, 2)
	g.AddEdge(2, 5, 4)
	g.AddEdge(3, 4, 9)
	g.AddEdge(3, 5, 14)
	g.AddEdge(4, 5, 10)
	g.AddEdge(5, 6, 2)
	g.AddEdge(6, 7, 1)
	g.AddEdge(6, 8, 6)
	g.AddEdge(7, 8, 7)

	// Calculate the shortest path from node 0 to all other nodes
	distances := g.Dijkstra(0)

	// Print the shortest path to all nodes
	fmt.Println("Shortest path from node 0 to all other nodes:")
	for i, distance := range distances {
		fmt.Printf("Node %d: %d\n", i, distance)
	}
}
