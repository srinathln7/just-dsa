package main

import "fmt"

func GetCurrencyConversion(graph [][3]int, src, dst int, isBidirectional bool) float32 {

	// Form the adjacency list
	adj := make(map[int][]Node)
	for _, element := range graph {
		src, dst, rate := element[0], element[1], element[2]
		adj[src] = append(adj[src], Node{to: dst, rate: float32(rate)})

		if isBidirectional {
			adj[dst] = append(adj[dst], Node{to: src, rate: 1 / float32(rate)})
		}
	}

	fmt.Println(adj)
	visited := make(map[int]bool)
	var result float32 = 1.0
	var dfs func(src int)
	dfs = func(src int) {
		if src == dst {
			return
		}

		visited[src] = true
		for _, node := range adj[src] {
			if !visited[node.to] {
				result *= node.rate
				dfs(node.to)
			}
		}
	}

	dfs(src)
	return result
}
