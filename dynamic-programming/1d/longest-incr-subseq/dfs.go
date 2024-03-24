package main

func LengthOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	// Memoization table to store the length of longest increasing subsequence starting from each index
	dp := make([]int, n)
	dp[n-1] = 1 // Base case: The last element itself is a subsequence of length 1

	// Recursive function to find the length of the longest increasing subsequence
	var dfs func(idx int) int
	dfs = func(idx int) int {
		if dp[idx] != 0 {
			return dp[idx] // If the result is already computed, return it
		}

		maxLen := 1 // Minimum length is 1 (the element itself)
		for i := idx + 1; i < n; i++ {
			if nums[i] > nums[idx] {
				maxLen = max(maxLen, 1+dfs(i)) // Try including the next element
			}
		}
		dp[idx] = maxLen // Store the result in the memoization table
		return maxLen
	}

	// Start the recursion from each index and return the maximum length found
	maxLen := 0
	for i := 0; i < n; i++ {
		maxLen = max(maxLen, dfs(i))
	}
	return maxLen
}
