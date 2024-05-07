package main

func arrangeCoins(n int) int {

	// Key Idea: To make `k` complete rows, we need `1+2+3...k` coins. The lowest possible row is 1 (when n=1)
	// and highest possible row (r=n) when n=1. We run a binary search narrowing down the range.

	l, r := 1, n
	for l <= r {
		mid := l + (r-l)/2
		if n == sumTillK(mid) {
			return mid
		}

		switch {
		// If the number of coins we have is lesser than the no.of coins to build `mid` complete rows
		// we narrow down the range to the left
		case n < sumTillK(mid):
			r = mid - 1

		// Opposite to above
		default:
			l = mid + 1
		}
	}

	// Represents the max. no of rows where sum of `k` rows is lesser than or equal to `n`
	// `l`: would be pointing to the first value where the sum exceeds n which is not the max. valid of k
	return r
}

func sumTillK(k int) int {
	return (k * (k + 1)) / 2
}
