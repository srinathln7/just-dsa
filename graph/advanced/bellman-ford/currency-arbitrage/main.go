package main

import (
	"fmt"
	"math"
)

type Currency int

type Node struct {
	id   Currency
	rate float64
}

const (
	USD Currency = iota
	EUR
	GBP
	CHF
)

var exchange = map[Currency]map[Currency]float64{
	USD: {USD: 1.00, EUR: 0.92, GBP: 0.79, CHF: 0.89},
	EUR: {USD: 1.09, EUR: 1.00, GBP: 0.86, CHF: 0.96},
	GBP: {USD: 1.28, EUR: 1.19, GBP: 1.00, CHF: 1.13},
	CHF: {USD: 1.13, EUR: 1.04, GBP: 0.88, CHF: 1.00},
}

type Edge struct {
	from, to int
	weight   float64
}

type Graph struct {
	V, E int // #Vertices = `V` #Edges = `E`

	edges []Edge
}

func getGraph(exchange map[Currency]map[Currency]float64) *Graph {
	var edges []Edge
	for key, valueMap := range exchange {
		for neighbour, weight := range valueMap {
			if neighbour == key {
				continue
			}
			edges = append(edges, Edge{from: int(key), to: int(neighbour), weight: -math.Log10(weight)})
		}
	}

	return &Graph{V: len(exchange), E: len(edges), edges: edges}
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// Outputs all the nodes that are part of the negative weight cycle
func getNegativeCycle(graph Graph, src int) []int {

	n, m := graph.V, graph.E

	dist := make([]float64, n)
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = math.Inf(1)
		parent[i] = -1
	}

	dist[src] = 0

	// Relax the edges `n-1` times
	for i := 0; i < n-1; i++ {
		temp := make([]float64, n)
		copy(temp, dist)

		for _, edge := range graph.edges {
			from, to, weight := edge.from, edge.to, edge.weight
			temp[to] = min(temp[to], dist[from]+weight)
			parent[to] = from
		}
		dist = temp
	}

	// Check for negative cycles - Relax further and check if the values change
	var negativeCycle []int
	c := -1
	for i := 0; i < m; i++ {
		for _, edge := range graph.edges {
			src, dst, weight := edge.from, edge.to, edge.weight
			if dist[src] != math.Inf(1) && dist[src]+weight < dist[dst] {
				c = dst
				break
			}
		}
	}

	if c != -1 {
		for i := 0; i < n; i++ {
			c = parent[c]
		}

		// To store the cycle vertex

		for v := c; ; v = parent[v] {
			negativeCycle = append(negativeCycle, v)

			if v == c && len(negativeCycle) > 1 {
				break
			}
		}

		// Reverse cycle[]
		for i, j := 0, len(negativeCycle)-1; i < j; i, j = i+1, j-1 {
			negativeCycle[i], negativeCycle[j] = negativeCycle[j], negativeCycle[i]
		}

		// Printing the negative cycle
		for _, v := range negativeCycle {
			fmt.Print(v, " ")
		}

		fmt.Println()
	}

	return negativeCycle
}

func main() {

	graph := getGraph(exchange)
	var arbitrages [][]int
	for i := 0; i < graph.V; i++ {
		cycles := getNegativeCycle(*graph, i)
		if len(cycles) != 0 {
			arbitrages = append(arbitrages, cycles)
		}
	}

	fmt.Println(arbitrages)
}
