package main

func longestPalindromeSubseq(s string) int {

	// Key Idea: To solve this problem, we need to think what constitutes a palidrome? The string in forward and reverse order
	// are the same. To find the longest palidromic subsequence is equivalent to finding the longest common subsequence (LCS) between
	// the string and the reverse of its string. We can solve the LCS using BOTTOM-UP Dynamic programming approach

	n := len(s)
	revStr := reverse(s)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}

	// Base cases:
	// Case 1: LCS b/w two empty strings - default to zero
	// case 2: LCS b/w an empty string and a non-empty string - default to zero
	// case 3: LCS b/w an non-empty string and a empty string - default to zero

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == revStr[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[n][n]
}

func reverse(s string) string {
	n := len(s)
	revStr := make([]byte, n)
	for i := 0; i < n; i++ {
		revStr[i] = s[n-i-1]
	}
	return string(revStr)
}
