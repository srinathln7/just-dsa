package main

func maxProduct(nums []int) int {

	// Key Idea: We can solve this using 1D- dynamic programming. Keep track of max. product sub-array is not a straightforward
	// applic'n of kadence's algorithm since we are also dealing with negative numbers. As multiplying two negative numbers can
	// result in positive number which could result in a max, we also need to keep track of min. product encountered so far. If the
	// incoming number is negative, we SWAP the maxProduct and minProduct to make sure we always get the max. pdt sub-array

	maxSoFar, minSoFar, maxGlobal := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]

		// If the incoming num is negative, then we swap the maxSoFar and minSoFar to ensure we always calculate the max until that subarray
		// Rem. negative * negative -> positive
		if num < 0 {
			maxSoFar, minSoFar = minSoFar, maxSoFar
		}

		maxSoFar = max(num, num*maxSoFar)
		minSoFar = min(num, num*minSoFar)
		maxGlobal = max(maxGlobal, maxSoFar)
	}

	return maxGlobal
}
