package main

import (
	"fmt"
)

// TimeComplexity : O(n.m) where `n` is the number of items and `m` is the capacity
// SpaceComplexity : O(m)
func calculateMaxProfit(weight, profit []int, capacity int) int {

	// Key Idea: We adopt a bottom-up dynamic programming approach. Visualize a n * (m+1) - 2D grid where
	// the rows represent the items and cols. represent the capacity from 0 to the given capacity. We can
	// determine the max profit every item-by-item and capacity-by-capacity and fill up this 2-D grid.
	// IMPORTANT: As we progress along, we will observe that the max-profit in the current cell is dependent only
	// on the previous row and hence, we can optimize the space from O(nm) to O(m). However, conceptually, we
	// can visualize it as a 2D-grid.

	n, m := len(profit), capacity
	dp := make([]int, m+1)

	// Fill-up the first row i.e. how would the grid look like if u add the first item
	// `c` -> capacity iterator
	for c := 0; c < m+1; c++ {
		// We can add the item to the grid only if it's weight is lesser than or equal to the current grid's capacity
		if weight[0] <= c {
			dp[c] = profit[0]
		}
	}

	// Fill up the entire grid for ith item and cth capacity
	// NOTE: Although we are operating only on 1D-grid, conceptually it is still a 2-D grid
	for i := 1; i < n; i++ {
		currentRow := make([]int, m+1)
		for c := 1; c < m+1; c++ {
			// Decision1: To skip (exclude) the current item, then profit essentially is the same as adding only the previous item
			// This is captured in the previous iteration and for i=1, we filled it up with first item's weight above.
			exProfit := dp[c]

			// Decision2: To include the current item if the current cell capacity can accomodate the current item
			var inclProfit int
			if weight[i] <= c {
				prevProfit := dp[c-weight[i]] // Represents the profit obtained from adding the previous item
				inclProfit = profit[i] + prevProfit
			}

			// Fill-up the current row
			currentRow[c] = max(exProfit, inclProfit)
		}

		// Replace previous row with current row and continue the process
		dp = currentRow
	}

	return dp[m]
}

func main() {
	profit := []int{4, 4, 7, 1}
	weight := []int{5, 2, 3, 1}
	capacity := 8

	fmt.Println("Maximum profit through brute-force DFS:", dfs(weight, profit, capacity))                             // Output: Maximum profit: 12
	fmt.Println("Maximum profit through bottom-up dynamic programming", calculateMaxProfit(weight, profit, capacity)) // Output: Maximum profit: 12
}
