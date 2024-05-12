package main

func minReorder(n int, connections [][]int) int {

	// Key Idea: The roads are too narrow and can always be only one way. Intutively, if all cities need to
	// reach city 0 for the big event, it means there should be no outgoing path(edges) from city 0. So
	// to find the minimum number of flips required, we run a dfs starting from city 0 and count the total number of
	// outgoing paths starting from city 0. This result represents the number of directions that needs to be reversed.
	// To account for dir'n, we treat this graph as a bi-directional or undirected graph.

	adj := make(map[int][]int)
	for _, edge := range connections {
		src, dst := edge[0], edge[1]

		// Signifies an outgoing edge from `src` to `dst`
		adj[src] = append(adj[src], dst)

		// Reverse the direction => Signifies an incoming edge from `dst` to `src`
		adj[dst] = append(adj[dst], -src)
	}

	var count int
	visited := make(map[int]bool, n)
	var dfs func(int)
	dfs = func(node int) {
		if visited[node] {
			return
		}

		visited[node] = true
		for _, neighbor := range adj[node] {

			if !visited[abs(neighbor)] {

				// Forward dir'n i.e. outgoing edge that needs to be reversed
				if neighbor > 0 {
					count++
					dfs(neighbor)
				} else {
					// Run a normal DFS since these directions need not be reversed
					dfs(abs(neighbor))
				}
			}
		}
	}

	dfs(0)
	return count
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
