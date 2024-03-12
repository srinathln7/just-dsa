package main

func findOrder(numCourses int, prerequisites [][]int) []int {

	// Key Idea: We form the graph of each course using adjacency list and then perform a DFS to detect cycles. For this, we keep track of two maps
	// `visited` and `visting` to detect the vistited vertex and visiting vertex in each iteration.
	// If the graph is found to be cyclic, we cannot run topological sort on that graph and hence return nil.
	// If the graph is accyle, we run  topological sort and return the slice. NOTE: In this specific example,
	// we do not have to reverse the array once we build the topo. array. Refer further down for explanation.

	var topo []int

	// IMPORTANT: Form the adjacency list based on courses sicne we have to return the ordering of the courses
	graph := make(map[int][]int)
	for _, prerequisite := range prerequisites {
		course, prereqs := prerequisite[0], prerequisite[1]
		graph[course] = append(graph[course], prereqs)
	}

	visited := make(map[int]bool)
	visiting := make(map[int]bool)

	// Define a DFS function to detect cycles
	var hasCycle func(course int) bool
	hasCycle = func(course int) bool {

		// Already detected in current recursion => cycle detected
		if visiting[course] {
			return true
		}

		// Already checked and detected no cycles - no need to repeat the work
		if visited[course] {
			return false
		}

		// Make the course as visited
		visited[course] = true

		// Mark it the in the current visiting map
		visiting[course] = true
		for _, prereq := range graph[course] {
			if hasCycle(prereq) {
				return true
			}
		}
		// Backtrack: Mark visiting back to false after finished traversing the graph for that course
		visiting[course] = false

		topo = append(topo, course)
		return false
	}

	// First, ensure, the graph has no cycles
	for course := 0; course < numCourses; course++ {
		if hasCycle(course) {
			return nil
		}
	}

	// Once, the graph is confirmed to be a DAG, we can run a topo ordering on it
	// No need to reverse the ordering here. This is because of the nature of the problem
	// For example, the pair [1, 0], indicates that to take course 1 you have to first take course 0.
	// The graph we form by itself is in reverse order
	// This means in the actual topological ordering would be [0, 1].

	return topo
}
