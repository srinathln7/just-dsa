package main

type State struct {
	idx    int
	canBuy bool
	k      int // IMPORTANT: No. of transactions completed
}

func MaxProfit(prices []int) int {

	n := len(prices)
	if n == 0 || n == 1 {
		return 0
	}

	initState := State{canBuy: true}

	// Declare a dp map for memoization
	dp := make(map[State]int)

	var dfs func(s State) int
	dfs = func(s State) int {

		currIdx := s.idx

		// We do s.k > 4 because atmost two transactions are only allowed where
		// each transaction represents a buy and sell operation. Here we count
		// `buy` as one single transaction and `sell` as another single transaction.
		// Hence we do `s.k >4`
		if currIdx >= n || s.k > 4 {
			return 0
		}

		// Memoization
		if val, exists := dp[s]; exists {
			return val
		}

		// Decision to cooldown
		coolDownProfit := dfs(State{idx: currIdx + 1, canBuy: s.canBuy, k: s.k})

		switch s.canBuy {

		// Decision to buy/cooldown
		case true:
			buyProfit := dfs(State{idx: currIdx + 1, canBuy: false, k: s.k + 1}) - prices[currIdx]
			dp[s] = max(buyProfit, coolDownProfit)

		// Decision to sell/down
		case false:
			sellProfit := dfs(State{idx: currIdx + 1, canBuy: true, k: s.k + 1}) + prices[currIdx]
			dp[s] = max(sellProfit, coolDownProfit)
		}

		return dp[s]
	}

	return dfs(initState)
}
