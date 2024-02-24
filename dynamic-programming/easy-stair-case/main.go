package main

func climbStairs(n int) int {

	// Since we are given only two options, the number of unique ways we can climb `n` steps
	// would be the sum of the number of unique ways of climbing `n-1` and `n-2` steps.
	// In other words, f[n] = f[n-1] + f[n-2]

	// f := make([]int, n+1)
	// // Store number off unique ways in f[i] array where index `i` represents the `i+1`th step
	// f[0]= 1
	// f[1] =2

	// for i := 2; i < n; i++ {
	//         f[i] = f[i-1] + f[i-2]
	// }

	// Space complexity can be further optimized
	current := 1
	next := 2

	for i := 2; i <= n; i++ {
		current, next = next, current+next
	}

	return current
}
