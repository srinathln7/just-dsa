package main

func dfs(node int, target int, adjList map[int][]int, visit map[int]bool) int {

	// Avoid infinite looping - on self loops
	if visit[node] {
		return 0
	}
	if node == target {
		return 1
	}

	count := 0
	visit[node] = true
	for _, neighbor := range adjList[node] {
		count += dfs(neighbor, target, adjList, visit)
	}

	// Backtrack
	delete(visit, node)

	return count
}
