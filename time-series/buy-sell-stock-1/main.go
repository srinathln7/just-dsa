package main

func maxProfit(prices []int) int {

	// Key Idea: We keep track of the minPrice and maxProfit that can obtained after every iteration
	// Whenever you dont encounter a price that is lesser than the minimum try selling and checking
	// if the resulting profit is greater than the current known max profit.

	if len(prices) == 0 {
		return 0
	}

	var minPrice int = prices[0] // Buy Price
	var maxProfit int

	// IMPORTANT: Is that the loop moves only forward
	// Buy before you can Sell idea is taken care by the loop
	// Aim to buy at the lowest possible price and sell at the highest possible price
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			maxProfit = max(maxProfit, prices[i]-minPrice)
		}
	}

	return maxProfit
}
