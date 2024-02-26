package main

func uniquePaths(m int, n int) int {

	// We use 2D-dynamic programming BOTTOM-UP approach method
	// We start small by filling up the number of unique ways we can reach the top row and first column
	// from the top-left corner (0,0) which is only one way. Since we can move only to the right or bottom
	// the number of unique ways we can reach a point (i,j) = #(i,0) + #(0,j)

	// Declare a 2D-DP array
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// Base cases
	dp[0][0] = 1

	// Unique ways to reach rows from top-left corner
	for i := 1; i < m; i++ {
		dp[i][0] = 1
	}

	// Unique ways to reach columns from top-left corner
	for j := 1; j < n; j++ {
		dp[0][j] = 1
	}

	// For a generic point (i,j), i,j != 0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
