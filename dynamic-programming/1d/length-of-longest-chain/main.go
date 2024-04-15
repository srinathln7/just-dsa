package main

import (
	"fmt"
)

func longestIncreasingSubsequenceLength(blocks [][2]int) int {

	// Key Idea: We use 1D-dynamic programming. To start simple, lets start from the reverse and work our way backwards.
	// This is because from the last index there is no element to the right of it and hence the LIS (longest increasing sub-seq) of only 1
	// element is always 1. Then when we move backwards where we can add the element to our sub-seq only when the number to the right of
	// the current index is strictly greater than the number at the current index. We store the maxLen of our computation at every
	// index in our DP array.

	n := len(blocks)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	for i := n - 1; i >= 0; i-- {
		// All elements to the right of the current element (idx `j`) should be strictly increasing
		for j := i + 1; j < n; j++ {
			if blocks[j][0] > blocks[i][1] {
				// You can make a decision to exclude the current element or include the current element along with the element to the right
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

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {

	var n, a, b int
	fmt.Scan(&n)

	blocks := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b)
		blocks[i] = [2]int{a, b}
	}

	fmt.Println(longestIncreasingSubsequenceLength(blocks))
}
