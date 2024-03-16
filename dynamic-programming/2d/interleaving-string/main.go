package main

func isInterleave(s1 string, s2 string, s3 string) bool {

	// Key Idea: There are a couple of important observations in this problem.
	// Obsv1: The total no. of characters must always be the sum of the two strings.
	// Obsv2: The relative ordering of the two sub-strings is always maintained.
	// Obsv3: Since the total. no of characters in a string is the sum of the two strings,
	// at any given time, s3's pointer will be sum of s1's ptr (i) + s2's ptr (j)

	// BOTTOM-UP DYNAMIC PROGRAMMING: dp[i][j] asks if we can build first `i+j` characters of s3 by
	// interleaving a substring of s1 up to first `i` chars and a substring of s2 up to first `j` chars.

	m, n := len(s1), len(s2)
	if len(s3) != m+n {
		return false
	}

	// We build a (m+1)*(n+1) DP grid
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}

	// Base case: Two empty strings can be interleaved to form another empty sub-string
	dp[0][0] = true

	// Fill up the first column - Answer the question if you can form s3 purely by using s1 alone
	// dp[i][0] - asks can you form s3 up to index (i+0) by interleaving s1 up to index i
	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}

	// Fill up the first row - Answer the question if you can form s3 purely by using s2 alone
	// dp[0][j] - asks can you form s3 up to index (0+j) by interleaving s2 up to index j
	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}

	// Fill up the rest of the array
	for i := 1; i <= m; i++ {

		for j := 1; j <= n; j++ {

			// Check if an element to your left or above is true

			// IMPORTANT: If the element to your left is true, we proceed along the right dir'n i.e. compare s3 with s2
			// If the element to your top is true, we proceed along the downwards dir'n i.e. compare s3 with s1

			// Left true or Top true
			if (dp[i][j-1] && s2[j-1] == s3[i+j-1]) || (dp[i-1][j] && s1[i-1] == s3[i+j-1]) {
				dp[i][j] = true
			}
		}
	}

	return dp[m][n]
}
