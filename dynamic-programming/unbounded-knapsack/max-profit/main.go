package main

import "fmt"

func calcMaxProfit(weight, profit []int, capacity int) int {

	// Key Idea: Similar to fixed-knapsack problem, we build a 2D-grid of size n * (m+1). dp[i][j]
	// represents the max. profit obtained by including or not including the ith item. To optimize space
	// we can do this problem in 1D since at a time only the previous row is required for the calculation.

	n, m := len(profit), capacity
	dp := make([]int, m+1)

	// Range over items from the very first item
	// NOTICE how it is different in fixed knapsack problem
	for i := 0; i < n; i++ {
		currentRow := make([]int, m+1)

		// Range over capacity
		for c := 0; c <= m; c++ {

			// Decision not to include the item at all
			exProfit := dp[c]

			var inclProfit int
			if weight[i] <= c {
				// Decision to include the item unboundedly till max. capacity
				// NOTICE how this line varies in fixed knapsack problem
				inclProfit = profit[i] + currentRow[c-weight[i]]
			}

			currentRow[c] = max(exProfit, inclProfit)
		}
		dp = currentRow
	}

	return dp[m]
}

func main() {
	profit := []int{4, 4, 7, 1}
	weight := []int{5, 2, 3, 1}
	capacity := 8

	fmt.Println("Maximum profit through bottom-up dynamic programming", calcMaxProfit(weight, profit, capacity)) // Output: Maximum profit: 18
}
