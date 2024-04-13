package main

// import (
// 	"fmt"
// )

// // Structure to represent a weighted edge in graph
// type Edge struct {
// 	src, dest, weight int
// }

// // Structure to represent a directed and weighted graph
// type Graph struct {
// 	// V. Number of vertices, E. Number of edges
// 	V, E int
// 	// Graph is represented as an array of edges.
// 	edge []Edge
// }

// // Creates a new graph with V vertices and E edges
// func createGraph(V, E int) Graph {
// 	graph := Graph{V: V, E: E}
// 	graph.edge = make([]Edge, E)
// 	return graph
// }

// // Function runs Bellman-Ford algorithm and prints negative cycle(if present)
// func negCycleBellmanFord(graph Graph, src int) {
// 	V := graph.V
// 	E := graph.E
// 	dist := make([]int, V)
// 	parent := make([]int, V)

// 	// Initialize distances from src to all other vertices as INFINITE and all parent as -1
// 	for i := 0; i < V; i++ {
// 		dist[i] = 1000000
// 		parent[i] = -1
// 	}
// 	dist[src] = 0

// 	// Relax all edges |V| - 1 times.
// 	for i := 1; i <= V-1; i++ {
// 		for j := 0; j < E; j++ {
// 			u := graph.edge[j].src
// 			v := graph.edge[j].dest
// 			weight := graph.edge[j].weight

// 			if dist[u] != 1000000 && dist[u]+weight < dist[v] {
// 				dist[v] = dist[u] + weight
// 				parent[v] = u
// 			}
// 		}
// 	}

// 	// Check for negative-weight cycles
// 	C := -1
// 	for i := 0; i < E; i++ {
// 		u := graph.edge[i].src
// 		v := graph.edge[i].dest
// 		weight := graph.edge[i].weight

// 		if dist[u] != 1000000 && dist[u]+weight < dist[v] {
// 			// Store one of the vertex of the negative weight cycle
// 			C = v
// 			break
// 		}
// 	}

// 	if C != -1 {
// 		for i := 0; i < V; i++ {
// 			C = parent[C]
// 		}

// 		// To store the cycle vertex
// 		cycle := []int{}
// 		for v := C; ; v = parent[v] {
// 			cycle = append(cycle, v)

// 			if v == C && len(cycle) > 1 {
// 				break
// 			}
// 		}

// 		// Reverse cycle[]
// 		for i, j := 0, len(cycle)-1; i < j; i, j = i+1, j-1 {
// 			cycle[i], cycle[j] = cycle[j], cycle[i]
// 		}

// 		// Printing the negative cycle
// 		for _, v := range cycle {
// 			fmt.Print(v, " ")
// 		}

// 		fmt.Println()
// 	} else {
// 		fmt.Println(-1)
// 	}
// }

// // Driver Code
// func Main() {
// 	// Number of vertices in graph
// 	V := 5

// 	// Number of edges in graph
// 	E := 5

// 	graph := createGraph(V, E)

// 	// Given Graph
// 	graph.edge[0] = Edge{src: 0, dest: 1, weight: 1}
// 	graph.edge[1] = Edge{src: 1, dest: 2, weight: 2}
// 	graph.edge[2] = Edge{src: 2, dest: 3, weight: 3}
// 	graph.edge[3] = Edge{src: 3, dest: 4, weight: -3}
// 	graph.edge[4] = Edge{src: 4, dest: 1, weight: -3}

// 	// Function Call
// 	negCycleBellmanFord(graph, 0)

// 	// Output -> 1 2 3 4 1
// }
