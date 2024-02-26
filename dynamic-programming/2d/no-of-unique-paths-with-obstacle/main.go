package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// We adopt a 2D-dynamic programming bottom-up approach.
	// We use the given obstacle grid to fill up our dp array
	// We set the number of paths to zero if we have an obstacle else
	// we set it to the sum of #unique ways we can reach that row and column

	// Declare and init a 2D-dp array
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// Base cases

	// If there is no obstacle we init the top-left corner to 1
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	// First row: When you encounter an obstacle you just BREAK
	// This will set the current column and all subsequent columns in the first row to zero (default value)
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			break
		}
		dp[0][j] = 1
	}

	// First column: When you encounter an obstacle you just BREAK
	// This will set the current row and all subsequent rows in the first column to zero (default value)
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			switch obstacleGrid[i][j] {
			// Obstacle
			case 1:
				dp[i][j] = 0
			// No obstacle
			case 0:
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}
