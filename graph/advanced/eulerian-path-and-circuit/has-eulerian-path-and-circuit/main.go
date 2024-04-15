package main

import (
	"fmt"
)

type Graph struct {
	adjacencyList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[int][]int),
	}
}

func (g *Graph) addEdge(u, v int) {
	g.adjacencyList[u] = append(g.adjacencyList[u], v)
	g.adjacencyList[v] = append(g.adjacencyList[v], u) // For undirected graph
}

func (g *Graph) eulerianPathCircuitExists() (bool, bool) {
	oddDegreeCount := 0
	for _, neighbours := range g.adjacencyList {
		if len(neighbours)%2 != 0 {
			oddDegreeCount++
		}
	}

	hasEulerianPath := oddDegreeCount == 2 || oddDegreeCount == 0
	hasEulerianCircuit := oddDegreeCount == 0

	return hasEulerianPath, hasEulerianCircuit
}

func main() {
	graph := NewGraph()
	graph.addEdge(0, 1)
	graph.addEdge(0, 2)
	graph.addEdge(1, 3)
	graph.addEdge(1, 4)
	graph.addEdge(2, 3)
	graph.addEdge(2, 4)

	hasPath, hasCircuit := graph.eulerianPathCircuitExists()
	fmt.Println("Graph has Eulerian Path:", hasPath)
	fmt.Println("Graph has Eulerian Circuit:", hasCircuit)

	// OUTPUT
	// Graph has Eulerian Path: true
	// Graph has Eulerian Circuit: false
}
