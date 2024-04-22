package main

type state struct {
	idx   int
	isBuy bool
}

func maxProfit(prices []int) int {

	// Key Idea: Draw the logical decision tree and we see the most intutive way to solve this problem is to
	// use dynamic programming with memoization technique. At any given time, we keep track of state consisting of two variables
	// current idx and state at the current idx (buy/sell/cooldown).

	n := len(prices)
	switch {
	case n == 0 || n == 1:
		return 0
	}

	initState := state{isBuy: true}
	dp := make(map[state]int)
	var dfs func(state) int
	dfs = func(s state) int {

		if val, exists := dp[s]; exists {
			return val
		}

		if s.idx >= n {
			return 0
		}

		// Decision tree
		currentIdx := s.idx

		// For cooldown, we have no profit/loss. We just update the index
		cooldown := dfs(state{idx: currentIdx + 1, isBuy: s.isBuy})
		switch s.isBuy {

		// Decision to either buy or cooldown
		// Recursively calculate the profit that can be obtained by buying the stock at this index.
		// We set the next state by incrementing the index and changing the buy state to false
		// Set max. profit to be max b/w profit obtained from buying or cooling down
		case true:
			currProfit := dfs(state{idx: currentIdx + 1, isBuy: false}) - prices[currentIdx]
			dp[s] = max(currProfit, cooldown)

		// Decision to either sell or cooldown
		// Recursively calculate the profit that can be obtained by selling the stock at this index.
		// Set max. profit to be max b/w profit obtained from selling the stock or cooling down. We set the next state
		// by double incrementing the index and changing the buy state to true
		// IMPORTANT: we update the idx to `currentIdx+2` because of the cooldown period i.e. we cannot buy the next day due to
		// the forced cool down constraint
		case false:
			currProfit := dfs(state{idx: currentIdx + 2, isBuy: true}) + prices[currentIdx]
			dp[s] = max(currProfit, cooldown)
		}

		return dp[s]
	}

	// Trigger the dfs with initstate isBuy as true
	return dfs(initState)
}
