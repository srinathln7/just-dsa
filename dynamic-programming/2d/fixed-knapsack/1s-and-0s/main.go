package main

import "strings"

func findMaxForm(strs []string, m int, n int) int {

	// Key Idea: Our subset should contain atmost `m` 0s and `n` 1s. To find teh largest subset number, we form a decision tree, where we go down two paths
	// of including the current element or not including it. We use memoziation to make sure we dont do any repeated work by stroing (m, n, idx). If we include
	// the current element in the subset, we then increment the count by 1. Finally, since we have to find the largest subset of the two, we find the max between
	// two decisions and return this value.

	dp := make(map[[3]int]int)
	var dfs func(m, n, idx int) int
	dfs = func(m, n, idx int) int {

		// Base case: end of slice - No further subsets can be added
		if idx == len(strs) {
			return 0
		}

		if val, isExists := dp[[3]int{m, n, idx}]; isExists {
			return val
		}

		// Decision Tree

		// Decision 1: - To skip the current element in the subset
		dp[[3]int{m, n, idx}] = dfs(m, n, idx+1)

		// Decision 2: - To include (+1) the current element in the subset
		n1s, n0s := strings.Count(strs[idx], "1"), strings.Count(strs[idx], "0")

		// As long as the we have atmost m 0s and n 1s
		if n0s <= m && n1s <= n {
			dp[[3]int{m, n, idx}] = max(dp[[3]int{m, n, idx}], 1+dfs(m-n0s, n-n1s, idx+1))
		}

		return dp[[3]int{m, n, idx}]
	}

	return dfs(m, n, 0)
}
