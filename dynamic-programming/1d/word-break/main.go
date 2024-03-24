package main

func wordBreak(s string, wordDict []string) bool {

	// Key Idea: We use 1D-dynamic programming where dp[i] represents if we can segment the first `i` characters in  string `s` into a
	// space-seperated seq. of one or more dict. words. We use a set (i.e map in Go) for an efficient look up of words in wordDict
	// We iterate through all possible sub-strings at index `i` and check if we can segment the words
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	// Init a dp slice of length `n+1` - plus for base case empty string
	n := len(s)
	dp := make([]bool, n+1)

	// Base case: The first `0` chars i.e. empty string can always be segmented
	dp[0] = true
	for i := 1; i <= n; i++ {
		// Iterate through all possible sub-strings ending at index `i`
		for j := 0; j < i; j++ {
			// IMPORTANT: If the substring is contained in the word dict. and substring from 0 to j is segmented
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}
