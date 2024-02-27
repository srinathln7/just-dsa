package main

func maxProfit(prices []int) int {

	// Key Idea : VISUALISE THE GRAPH. The minute your stock goes up you sell it off.
	// This will give you the max possible profit.
	// It also account for stocks that are monotonically increasing as well.
	// Buy/Sell on the same day is allowed

	var maxProfit, profit int
	for i := 1; i < len(prices); i++ {
		profit = prices[i] - prices[i-1]
		// If its a profitable trade, add to the maxProfit
		if profit > 0 {
			maxProfit += profit
		}
	}

	return maxProfit
}
