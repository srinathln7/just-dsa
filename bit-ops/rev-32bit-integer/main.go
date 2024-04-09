package main

import "math"

func reverse(x int) int {

	// Key Idea: We can determine the overflow conditions by comparing the result in each iteration with the max and min
	// value of 32-bit integer. We then follow the usual process to reverse the numbers
	const INT_MAX, INT_MIN = math.MaxInt32, math.MinInt32

	result := 0
	for x != 0 {
		digit := x % 10

		// Check for overflow conditions

		// Max-out
		if result > INT_MAX/10 || (result == INT_MAX/10 && digit > INT_MAX%10) {
			return 0
		}

		// Min-out
		if result < INT_MIN/10 || (result == INT_MIN/10 && digit < INT_MIN%10) {
			return 0
		}

		result = result*10 + digit
		x /= 10
	}

	return result
}
