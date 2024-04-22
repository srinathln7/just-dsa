package main

func maxProfit(prices []int) int {

	// To calculate the max. profit that can be achieved with atmost two transactions
	// we calculate max profit achieveable before day `i` and max profit achieveable
	// after `day i`

	n := len(prices)
	maxProfitBefore := make([]int, n)
	maxProfitAfter := make([]int, n)

	// Calculate max. profit achieveableby completing atmost one transaction before day `i`
	minPrice := prices[0]
	for i := 1; i < n; i++ {
		maxProfitBefore[i] = max(maxProfitBefore[i-1], prices[i]-minPrice)
		minPrice = min(minPrice, prices[i])
	}

	// Calculate max. profit achievable by completing one transaction after day `i`
	maxPrice := prices[n-1]
	for i := n - 2; i >= 0; i-- {
		maxProfitAfter[i] = max(maxProfitAfter[i+1], maxPrice-prices[i])
		maxPrice = max(maxPrice, prices[i])
	}

	var result int
	for i := 0; i < n; i++ {
		result = max(result, maxProfitBefore[i]+maxProfitAfter[i])
	}

	return result
}
