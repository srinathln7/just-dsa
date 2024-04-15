package main

func numDecodings(s string) int {

	// Key Idea: We use bottom-up dynamic programming approach where we build the `dp` array
	// where dp[i] calculates the number of ways the number of ways we can decode a string
	// formed by the first `i` chars.

	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)

	// Number of ways to decode an empty string and a string with just one character is only 1 way
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		// If current char is zero then that char cannot be mapped by itself
		if s[i-1] == '0' {
			// For example: S -> 10, 20 => 1, 30 => 0,  | Eg: 110 => 1, 130 => 0
			if s[i-2] == '1' || s[i-2] == '2' {
				dp[i] = dp[i-2]
			} else {
				return 0
			}
		} else if s[i-2] == '1' || (s[i-2] == '2' && s[i-1] >= '1' && s[i-1] <= '6') {
			// For eg: S -> 19, 21, 26 => 2
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			// For eg: S -> 333 => 1 , 263 => 2
			dp[i] = dp[i-1]
		}
	}

	return dp[n]
}
