package main

import "fmt"

// topologicalSort: Runs a topological ordering of a DIRECTED ACYCLIC GRAPH (DAG's)
func topologicalSort(edges [][]int, n int) []int {
	var topoSort []int

	// First form the adjacency graph
	adj := make(map[int][]int)
	for _, edge := range edges {
		src, dst := edge[0], edge[1]
		adj[src] = append(adj[src], dst)
	}

	visited := make(map[int]bool)
	var dfs func(node int)
	dfs = func(node int) {
		// If the node has already been visited, no need to perform DFS further
		if visited[node] {
			return
		}

		visited[node] = true

		// DFS on the node's neighbour
		for _, neighbour := range adj[node] {
			dfs(neighbour)
		}

		// Build the topological order
		topoSort = append(topoSort, node)
	}

	// Run a DFS on every node
	for i := 1; i <= n; i++ {
		dfs(i)
	}

	// Reverse once the full topological ordering once the slice is fully built
	reverse(topoSort)
	return topoSort
}

func reverse(nums []int) []int {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
	return nums
}

func main() {
	edges := [][]int{{1, 3}, {1, 4}, {2, 4}, {3, 5}, {4, 6}, {5, 6}}
	n := 6
	fmt.Println("Topological Sort:", topologicalSort(edges, n))
	// Output: Topological Sort: [2 1 4 3 5 6]
}
