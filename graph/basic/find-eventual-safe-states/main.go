package main

func eventualSafeNodes(graph [][]int) []int {

	// Key Idea: We run DFS on every nodes in the given graph. Nodes are marked to be ""safe" when they dont
	// lead to cycle and its chidren nodes are safe.

	n := len(graph)
	safe := make([]bool, n)
	visited := make([]bool, n)

	var result []int
	var isSafe func(node int) bool
	isSafe = func(node int) bool {

		// If the node is already visited: return its safety status
		if visited[node] {
			return safe[node]
		}

		// Mark the node as visited
		visited[node] = true

		for _, child := range graph[node] {
			// Since every possible path from the node should lead to anothe safe/terminal node
			// so if even one of them is false, then that node is not a safe node.
			if !isSafe(child) {
				return false
			}
		}

		safe[node] = true
		return true
	}

	for i := 0; i < n; i++ {
		if isSafe(i) {
			result = append(result, i)
		}
	}

	return result
}
