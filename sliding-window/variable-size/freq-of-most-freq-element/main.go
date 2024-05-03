package main

import "sort"

func maxFrequency(nums []int, k int) int {

	// Key Idea: If we sort the array and initialize two pointers `l` and `r`. As we move the right pointer `r` forward,
	// the max. achievable sum is equal to sum of the window + k. If that is not attainable in that window, then we start
	// shrinking the window. How to decide that it is not attainable? Basically, we can try to increment all the elements in
	// the window to its highest value and check if that exceeds max. achieveable sum in that window. If it does, then we start
	// shrinking the window

	sort.Ints(nums)
	var maxFreq int

	l, sum := 0, 0
	for r := 0; r < len(nums); r++ {
		sum += nums[r]

		for (r-l+1)*nums[r] > sum+k {
			sum -= nums[l]
			l++
		}

		maxFreq = max(maxFreq, r-l+1)
	}

	return maxFreq
}
