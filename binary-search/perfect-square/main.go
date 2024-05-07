package main

func isPerfectSquare(num int) bool {

	// Key Idea: The lowest possible positive integer solution is 1 and the highest possible integer solution is `n`.
	// Lets us run a binary search and narrow down the search space until the whole search space is exhausted

	l, r := 1, num
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid == num {
			return true
		}
		switch {
		// If the sq. of the middle number is lesser than the number we can narrow down the search space to the right
		case mid*mid < num:
			l = mid + 1

			// If the sq. of the middle number is greater than the number we can narrow down the search space to the left
		default:
			r = mid - 1
		}
	}

	return false
}
