package main

func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	var minPrice int = prices[0]
	var maxProfit int

	// Key Idea: Is that the loop moves only forward
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
