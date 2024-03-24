package main

func lengthOfLIS(nums []int) int {

	// Key Idea: We use 1D- dynamic programming. To start simple, lets start from the reverse and work our way backwards.
	// This is because from the last index there is no element to the right of it and hence the LIS (longest increasing sub-seq) of only 1
	// element is always 1. Then when we move backwards where we can add the element to our sub-seq only when the number to the right of
	// the current index is strictly greater than the number at the current index. We store the maxLen of our computation at every
	// index in our DP array.

	n := len(nums)

	// Declare and init dp (LIS)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := n - 1; i >= 0; i-- {
		currElement := nums[i]
		// All elements to the right of the current element (idx `j`) should be strictly increasing
		for j := i + 1; j < n; j++ {
			if nums[j] > currElement {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
	}

	// Find the max element among all starting at all indexes
	maxLen := dp[0]
	for i := 1; i < n; i++ {
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}

	return maxLen
}
