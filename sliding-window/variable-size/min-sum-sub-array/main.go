package main

import "math"

func minSubArrayLen(target int, nums []int) int {

	// Simplest Base Case
	for _, val := range nums {
		if val >= target {
			return 1
		}
	}

	var left, windowSum int

	length := math.MaxInt32
	for right := range nums {
		windowSum += nums[right]

		// Whenever the condition is met, we start becoming "GREEDY" and try to
		// shrink the window and see if the condition still holds true. If it holds true,
		// we can minimize further otherwise we can move the right pointer.
		// SECOND LOOP: Does not run for every `r` value. It runs only when the condition is met
		// and the the "amortized" time complexity is O(n)
		for windowSum >= target {
			length = min(length, right-left+1)
			windowSum -= nums[left]
			left += 1
		}
	}

	// If the condition was never ever found then set the minimum length to zero
	if length == math.MaxInt32 {
		length = 0
	}

	return length
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
