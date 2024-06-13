package main

import (
	"fmt"
)

const (
	USD int = iota
	JPY
	GBP
	AUD
	CHN
	THAI
)

func main() {
	currencySet := make(map[string]int)
	currencySet["USD"] = USD
	currencySet["JPY"] = JPY
	currencySet["GBP"] = GBP
	currencySet["AUD"] = AUD
	currencySet["CHN"] = CHN
	currencySet["THAI"] = THAI

	graph := [][3]int{
		{0, 1, 4}, {1, 2, 1}, {0, 3, 8},
	}

	fmt.Println(getCurrencyConversion(graph, 2, 3))

	graph1 := [][3]int{
		{0, 1, 100},
		{1, 4, 20},
		{4, 5, 200},
	}

	queries := [][2]string{
		{"USD", "CHN"},
		{"JPY", "THAI"},
		{"USD", "INR"},
	}

	result := make([]float32, len(queries))
	for i, query := range queries {
		src, dst := query[0], query[1]
		from, srcExists := currencySet[src]
		to, dstExists := currencySet[dst]

		if !srcExists || !dstExists {
			result[i] = -1
			continue
		}

		result[i] = getCurrencyConversion(graph1, from, to)
	}

	fmt.Println(result)
}

type Node struct {
	to   int
	rate float32
}

func getCurrencyConversion(graph [][3]int, src, dst int) float32 {
	// Form the adjacency list
	adj := make(map[int][]Node)
	for _, element := range graph {
		src, dst, rate := element[0], element[1], element[2]
		adj[src] = append(adj[src], Node{to: dst, rate: float32(rate)})
		adj[dst] = append(adj[dst], Node{to: src, rate: 1 / float32(rate)})

	}

	visited := make(map[int]bool)
	queue := []Node{{to: src, rate: 1.0}}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.to == dst {
			return curr.rate
		}

		if visited[curr.to] {
			continue
		}

		visited[curr.to] = true

		for _, neighbor := range adj[curr.to] {
			if !visited[neighbor.to] {
				queue = append(queue, Node{to: neighbor.to, rate: curr.rate * neighbor.rate})
			}
		}
	}

	return -1.0
}
