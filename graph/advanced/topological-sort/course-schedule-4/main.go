package main

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {

	// Key Idea: We form the adjacency list based on the prerequisties i.e. For a given prequesite
	// we are able to potentially finish the list of courses. MOST IMPORTANT: To remember how to
	// form the adjancey list based on the question (query).  After that, if we are able to traverse
	// from course1 to course2 through DFS, then we can return true. Otherwise, false.
	// NOTE: Since Prerequisites can also be indirect. If course a is a prerequisite of course b, and
	//course b is a prerequisite of course c, then course a is a prerequisite of course c, we run a DFS.
	// If not, we can directly  determine the relationship just by looking at the adjaceny list.

	// IMPORTANT: Form the adjacency list based on prerequisites and NOT courses
	adj := make(map[int][]int, numCourses)
	for _, preReq := range prerequisites {
		prereq, course := preReq[0], preReq[1]
		adj[prereq] = append(adj[prereq], course)
	}

	// Let's build the results bool
	n := len(queries)
	result := make([]bool, n)

	// If there are no prerequisites, we simply return false
	if len(adj) == 0 {
		return result
	}

	for i, query := range queries {
		visited := make(map[int]bool, numCourses)
		course1, course2 := query[0], query[1]
		result[i] = dfs(adj, visited, course1, course2)
	}

	return result
}

func dfs(adj map[int][]int, visited map[int]bool, src, dst int) bool {
	if visited[src] {
		return false
	}

	// When we are able to traverse from src to dst, then the two nodes are connected
	// i.e. If we take course1, we are able to finish course2
	if src == dst {
		return true
	}

	visited[src] = true
	for _, neighbour := range adj[src] {
		if dfs(adj, visited, neighbour, dst) {
			return true
		}
	}

	return false
}
