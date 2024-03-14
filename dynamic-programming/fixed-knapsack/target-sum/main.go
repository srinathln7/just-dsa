package main

func findTargetSumWays(nums []int, target int) int {

	dp := make(map[[2]int]int) // (index, total) -> # of ways

	var dfsBT func(int, int) int
	dfsBT = func(idx, total int) int {

		// Base case: end of array
		if idx == len(nums) {
			if total == target {
				return 1
			}
			return 0
		}

		if val, isExists := dp[[2]int{idx, total}]; isExists {
			return val
		}

		// Decision tree: To either use `+` sign or `-` sign
		dp[[2]int{idx, total}] = dfsBT(idx+1, total+nums[idx]) + dfsBT(idx+1, total-nums[idx])
		return dp[[2]int{idx, total}]
	}

	// Start at index 0 with a total 0
	return dfsBT(0, 0)
}
