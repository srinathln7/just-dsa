package main

type State struct {
	idx    int
	canBuy bool
	K      int // Number of transactions completed
}

func MaxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 || n == 1 {
		return 0
	}

	initState := State{idx: 0, canBuy: true, K: 0}

	// Declare a dp map for memoization
	dp := make(map[State]int)

	var dfs func(s State) int
	dfs = func(s State) int {
		if s.idx >= n || s.K > 2*k {
			return 0
		}

		// Memoization
		if val, exists := dp[s]; exists {
			return val
		}

		// Decision to cooldown i.e. not do anything on that day
		coolDownProfit := dfs(State{idx: s.idx + 1, canBuy: s.canBuy, K: s.K})

		// Calculate buy or sell profit based on current state
		var profit int
		if s.canBuy {
			// Decision to buy
			buyProfit := dfs(State{idx: s.idx + 1, canBuy: false, K: s.K + 1}) - prices[s.idx]
			profit = max(buyProfit, coolDownProfit)
		} else {
			// Decision to sell
			sellProfit := dfs(State{idx: s.idx + 1, canBuy: true, K: s.K + 1}) + prices[s.idx]
			profit = max(sellProfit, coolDownProfit)
		}

		// Store the calculated profit in the memoization map
		dp[s] = profit

		return profit
	}

	return dfs(initState)
}
