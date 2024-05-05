package main

func numTrees(n int) int {

	// Key Idea: If we form the num array [1...n], each of these elements can be the root at some point in time.
	// We can use bottom-up DP approach where dp[i] represents the total number of nodes in the current tree

	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	// Calculate dp[i] which represents the number of unique BST's with i nodes
	for i := 2; i <= n; i++ {

		// For `i` number of nodes [1...i] => can be root of those nodes

		// For each node as root, calculate the number of unique BST's  by considering the number of left and right subtrees
		// Calculate the left and right subtrees with `j` as the root. `j-1` represents the number of nodes in left sub_tree
		// and `i-j` represents the number of nodes in right subtree
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}
