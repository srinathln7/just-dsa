package main

func checkIfPrerequisite(n int, prerequisites [][]int, queries [][]int) []bool {

	// Key Idea: This problem is to check if 2 vertices are connected in directed graph.
	// Floyd-Warshall O(n^3) is an algorithm that will output the minium distance of any vertices.
	// We can modifiy it to output if any vertices is connected or not.

	// Initialize a 2D boolean array to store the connectivity information
	connected := make([][]bool, n)
	for i := range connected {
		connected[i] = make([]bool, n)
	}

	// Mark prerequisites as connected
	for _, p := range prerequisites {
		connected[p[0]][p[1]] = true // p[0] -> p[1]
	}

	// Compute transitive closure using Floyd-Warshall algorithm
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				connected[i][j] = connected[i][j] || (connected[i][k] && connected[k][j])
			}
		}
	}

	// Store the answers for queries
	ans := make([]bool, len(queries))
	for i, q := range queries {
		ans[i] = connected[q[0]][q[1]]
	}
	return ans
}
