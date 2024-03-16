package main

func minDistance(word1 string, word2 string) int {

	// Key Idea: We can use a bottom-up dynamic programming approach where we split our big problem into smaller sub-problems
	// and answer basic questions: We do this by building an array dp[i][j] which answers the question "what is the minimum
	// number of operations needed to convert the first `i` characters in word1 to the first `j` characters in word2". We build
	// on top of that.

	if len(word1) == 0 {
		return len(word2)
	}

	if len(word2) == 0 {
		return len(word1)
	}

	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	// Base case: Min ops. req'd to convert one empty string to another empty string is zero
	dp[0][0] = 0

	// Fill up the first column dp[i][0] - what is the minimum number of ops. requ'd to convert the first `i` characters in
	// word1 to first `0` characters in word2 i.e. empty string? - It is equal to the length of the prefix in word1.
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}

	// Fill up the first row dp[0][j] - what is the minimum number of ops. requ'd to convert the first `0` characters in
	// word1 (i.e. empty string) to first `j` characters in word2 ? - It is equal to the length of the prefix in word2.
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	// Fill up the DP-array
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// Happy path - No. ops are needed
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {

				// IMPORTANT: This is the crux of the problem. Until now, we have determined the min. number of operations req'd
				// to match first `i-1` chars in word1 to the first `j-1` chars in word2.

				// Now, we can make three operations

				// 1. **Insertion**:
				// If we insert a character into `word1` to match the character at index `j` in `word2`, we would have to look at `dp[i][j-1]`.
				// Why `j-1`? Because we've already matched up to `j-1` in `word2`, and we're considering adding a character to `word1` to match `word2[j]`.
				// So, the minimum number of operations required is `dp[i][j-1] + 1`, where we add 1 because we're inserting a character.

				minInsertOps := dp[i][j-1]

				// 2. **Deletion**:
				// If we delete a character from `word1`, we effectively skip that character and match the rest of `word1` with `word2[j]`.
				// So, we would consider `dp[i-1][j]`. Here, we subtract 1 because we're deleting a character.

				minDeleteOps := dp[i-1][j]

				// 3. **Replacement**:
				// If we replace the character at index `i` in `word1` with the character at index `j` in `word2`, we would need to look at `dp[i-1][j-1]`.
				// The reason for `i-1` and `j-1` is because we're matching the characters at the current positions. Like the other operations, we add 1 because
				//we're performing a replacement.

				minReplaceOps := dp[i-1][j-1]

				// After considering these three operations, we take the minimum among them to get the minimum number of operations required to convert
				// the substring of `word1` ending at index `i` to the substring of `word2` ending at index `j`.
				// Since we are deinitely going to perform 1 operation atleast insertion, deletion or replacement, we increment the min of 3 ops by 1.
				dp[i][j] = 1 + min(minInsertOps, minDeleteOps, minReplaceOps)
			}
		}
	}

	return dp[m][n]
}
