package main

func canPartition(nums []int) bool {

	// Key Idea: It is impossible to divide an array with "odd" sum in to only two subsets with equal sum.
	// For ex [1,2,3,5] => Total sum:11, This means ideally both sub-array's should have sum as 5.5 which is
	// not possible since both are int sub-array's. Furthermore, say if `a=b+c`, if `a` is odd, then `b` and `c`
	// must be only two possibilities i.e. (b,c) => (odd,even) or (b,c) => (even, odd). This means we can never
	// have a sub-array with equal sum if the total sum is odd.
	// For array with `even` total sum. We can use a bottom-up dynamic programming approach and build a 2D-grid dp[i][j]
	// to determine if it is possible to achieve a sum of `j` using only FIRST `i` elements. This is the crux of the problem.
	// For ex: dp[0][0] asks: Can I build a target sum j=0 using only i=0 elements ?

	var sum int
	for _, num := range nums {
		sum += num
	}

	// If the sum is odd - return false immediately
	if sum&1 == 1 {
		return false
	}

	// Now, our target sum for each of the two subarray is sum/2
	targetSum := sum / 2
	n, m := len(nums), targetSum
	dp := make([][]bool, n+1)

	for i := 0; i < n+1; i++ {
		dp[i] = make([]bool, m+1)
	}

	// You can build a target sum =0 using 0 elements
	dp[0][0] = true

	// First row asks: Can you build a target sum=j, where j !=0 using first 0 elements? No. Default of bool is false
	// dp[0][j] = false

	// Fill up the first column: - meaning can you build a target sum =0 using first `n` elements. Yes, just dont include them
	for i := 1; i < n+1; i++ {
		dp[i][0] = true
	}

	// Fill-up the DP array
	for i := 1; i < n+1; i++ {
		currentNum := nums[i-1]
		for j := 1; j <= m; j++ {
			// `j` => current target sum

			switch {
			// If the current num is lesser than the current target sum, there is a possibility that
			// we can still match the current target sum if we include the previous sum
			case currentNum <= j:
				// Decision: Don't include this number or include this number
				dp[i][j] = dp[i-1][j] || dp[i-1][j-currentNum]

			// If current num is greater than the current target sum, then we cannot include this number which means our decision
			// is captured entirely by the previous iteration
			default:
				dp[i][j] = dp[i-1][j]

				// MISTAKE: If the current num alone is equal to the current target sum, then that does not mean the num alone is sufficient to decide `true`
				// This is because the question is : Can I achieve the current target sum using the FIRST `i` elements and not just the current element alone?
				// case currentNum == j:
				// 	dp[i][j] = true
			}
		}
	}

	// Returns if you can return a sum of `m` using the first `n` elements in nums
	return dp[n][m]
}
