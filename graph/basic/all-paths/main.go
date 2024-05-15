package main

func allPathsSourceTarget(graph [][]int) [][]int {

	var result [][]int
	var path []int

	n := len(graph)

	visited := make(map[int]bool)
	var dfs func(src int)
	dfs = func(src int) {

		visited[src] = true
		path = append(path, src)

		if src == n-1 {
			// IMPORTANT: To make a copy of the current path and not directly assign path to it
			// This is because in the subsequent recursive stack frames, path will be modified
			// and in case we are appending only path, then all subsequent modifications will
			// also be reflected in the `results`.
			newPath := make([]int, len(path))
			copy(newPath, path)
			result = append(result, newPath)
		}

		// DFS -Recurse
		for _, neighbor := range graph[src] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}

		// Backtrack
		visited[src] = false
		path = path[:len(path)-1]
	}

	dfs(0)
	return result
}
