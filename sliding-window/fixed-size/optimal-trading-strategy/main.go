package main

import "fmt"

func calcProfit(rates, strategy []int) int {
	var currentProfit int
	for i, action := range strategy {
		switch action {
		// Buy Operation
		case -1:
			currentProfit -= rates[i]
		// Sell operation
		case 1:
			currentProfit += rates[i]
		}
	}
	return currentProfit
}

func calcMaxProfit(rates, strategy []int, k int) ([]int, int) {

	// Declare and init variables
	var originalProfit int = calcProfit(rates, strategy)

	modifyStrategyWindow := make([]int, k)
	for i := k / 2; i < k; i++ {
		modifyStrategyWindow[k-i-1] = 0
		modifyStrategyWindow[i] = 1
	}

	// Calculate profit for each window with the current strategy and modified strategy and update the max difference accodingly
	var maxDiff, currProfitWindow, modifyProfitWindow int

	// First Window
	r := 0
	for i := 0; i < k; i++ {
		currProfitWindow = calcProfit(rates[0:k], strategy[0:k])
		modifyProfitWindow = calcProfit(rates[k/2:k], modifyStrategyWindow[k/2:k]) // The first `k/2` values are set to 0
		if modifyProfitWindow-currProfitWindow > maxDiff {
			maxDiff = modifyProfitWindow - currProfitWindow
			r = i
		}
		// maxDiff = max(maxDiff, modifyProfitWindow-currProfitWindow)
	}

	// Slide the window and update the max profit accordingly
	n := len(strategy)
	for i := k; i < n; i++ {
		currProfitWindow = calcProfit(rates[i-k+1:i+1], strategy[i-k+1:i+1])
		modifyProfitWindow = calcProfit(rates[i-k+1:i+1], modifyStrategyWindow[0:k])

		if modifyProfitWindow-currProfitWindow > maxDiff {
			maxDiff = modifyProfitWindow - currProfitWindow
			r = i
		}
		//maxDiff = max(maxDiff, modifyProfitWindow-currProfitWindow)
	}

	l := r - k + 1
	// Get optimal strategy
	optimalStrategy := make([]int, n)
	copy(optimalStrategy, strategy)

	maxProfit := originalProfit
	if originalProfit+maxDiff > originalProfit {
		maxProfit = originalProfit + maxDiff
		for i, j := l, 0; i < r+1 && j < k; i, j = i+1, j+1 {
			optimalStrategy[i] = modifyStrategyWindow[j]
		}
	}

	return optimalStrategy, maxProfit
}

func main() {
	rates := []int{2, 4, 1, 5, 10, 6}
	strategy := []int{-1, 1, 0, 1, -1, 0}
	k := 4

	optimalStrategy, maxProfit := calcMaxProfit(rates, strategy, k)
	fmt.Println("optimal strategy", optimalStrategy)
	fmt.Println("max profit", maxProfit)
}
