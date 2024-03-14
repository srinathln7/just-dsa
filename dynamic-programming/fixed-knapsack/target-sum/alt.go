package main

func FindTargetSumWays(nums []int, target int) int {

	// Calculate the total sum of nums
	var sum int
	for _, num := range nums {
		sum += num
	}

	// Since the max range we can build using expression `+` is sum, any target greater than that will not be possible
	if sum < abs(target) {
		return 0
	}

	offset := sum

	// Init a dp array
	n := len(nums)
	dp := make([][]int, n) // The no. of elements in the range [-max, max] is 2*max+1
	for i := range dp {
		dp[i] = make([]int, 2*offset+1)
	}

	// Base case - Recall the definition of dp[i][j] from above
	dp[0][offset+nums[0]] += 1 // Using positive sign
	dp[0][offset-nums[0]] += 1 // Using negative sign

	for i := 1; i < n; i++ {
		currentNum := nums[i]
		for j := -offset; j <= offset; j++ {
			// Check if the target index is within the valid range
			if j+currentNum <= offset {
				dp[i][j+offset] += dp[i-1][j+offset+currentNum] // Using positive sign
			}
			if j-currentNum >= -offset {
				dp[i][j+offset] += dp[i-1][j+offset-currentNum] // Using negative sign
			}
		}
	}

	return dp[n-1][target+offset]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
