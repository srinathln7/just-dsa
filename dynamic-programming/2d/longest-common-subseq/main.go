package main

func longestCommonSubsequence(text1 string, text2 string) int {

	// Edge cases
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}

	if text1 == text2 {
		return len(text1)
	}

	// We adopt a 2D-dynamic programming BOTTOM-UP approach
	m := len(text1)
	n := len(text2)

	// Declare a 2D matrix of size (m+1)*(n+1) where the top row and column would match the length of the common
	// subsequence b/w (text1, empty string) and (empty string, text2) which is zero (default value)

	/*
		dp := make([][]int, m+1)
		for i := 0; i < m+1; i++ {
			dp[i] = make([]int, n+1)
		}

		// Fill-up the 2D-dp array
		for i := 1; i < m+1; i++ {
			for j := 1; j < n+1; j++ {
				if text1[i-1] == text2[j-1] {
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = max(dp[i-1][j], dp[i][j-1])
				}
			}
		}
		return dp[m][n]

	*/

	// FOLLOW UP: A more optimized space solution can come from the observation that at a given time, we only need
	// the previous row to calculate. So space can be brought down to O(n) instead of O(n.m)
	dp1 := make([]int, n+1)
	for i := 1; i < m+1; i++ {
		currentRow := make([]int, n+1)
		for j := 1; j < n+1; j++ {
			if text1[i-1] == text2[j-1] {
				currentRow[j] = dp1[j-1] + 1
			} else {
				currentRow[j] = max(currentRow[j-1], dp1[j])
			}
		}
		dp1 = currentRow
	}
	return dp1[n]
}
