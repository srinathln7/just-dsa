package main

// Brute force Solution:  Time: O(2^n), Space: O(n) where n is the number of items.
func dfs(weight, profit []int, capacity int) int {
	return dfsHelper(weight, profit, capacity, 0)
}

func dfsHelper(weight, profit []int, capacity, idx int) int {
	var exProfit, inProfit, maxProfit int

	// Base case - When we have added all the items and still trying to add items
	if idx == len(profit) {
		return 0
	}

	// Form the decision tree:

	// Decision1: Skip the current item - excluded profit
	exProfit = dfsHelper(weight, profit, capacity, idx+1)

	// Decision2: Include the current item
	newCapacity := capacity - weight[idx]

	// Ensure the after adding the current item - our capacity does not exceed total capacity
	if newCapacity >= 0 {
		// included. profit will be the sum of the profit of the current item and the profit of the following decision trees
		inProfit = profit[idx] + dfsHelper(weight, profit, newCapacity, idx+1)
	}

	maxProfit = max(inProfit, exProfit)
	return maxProfit
}
