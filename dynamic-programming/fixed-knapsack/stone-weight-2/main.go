package main

import (
	"fmt"
	"math"
)

func lastStoneWeightII(stones []int) int {

	// Key Idea: We can reframe this problem into a fixed knapsack problem. If we observe in-depth, all we care about
	// is to assemble two piles of stone bags with (almost) similar weight. Then, if we smash those two bags, we can
	// find the difference to be our minimum. To accomplish this, we calculate the sum of these stones and divide it
	// by 2 (two bags) as our target. And we brute force through DFS, we will try to achieve a combination that is
	// closest to the target.

	var sum int
	for _, weight := range stones {
		sum += weight
	}

	target := int(math.Ceil(float64(sum) / 2))
	dp := make(map[[2]int]int)
	var dfs func(idx, total int) int
	dfs = func(idx, total int) int {
		if total >= target || idx == len(stones) {
			return abs(total - (sum - total))
		}

		// Memoization
		if val, isExists := dp[[2]int{idx, total}]; isExists {
			return val
		}

		// Decision tree
		dp[[2]int{idx, total}] = min(dfs(idx+1, total), dfs(idx+1, total+stones[idx]))
		return dp[[2]int{idx, total}]
	}

	result := dfs(0, 0)
	fmt.Println("DP Map", dp)
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func main() {

	stones := []int{31, 26, 33, 21, 40}
	fmt.Println(lastStoneWeightII(stones))
}
