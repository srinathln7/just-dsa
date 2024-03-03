package main

func PivotIndex(nums []int) int {

	// Key Idea: We construct the prefix sum array and keep iterating the index to check if a left sum and right sum
	// exists

	n := len(nums)
	prefixSum := make([]int, n)
	prefixSum[0] = nums[0]
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}

	// Check for pivot at left most first
	sumExceptFirstIdx := prefixSum[n-1] - prefixSum[0]
	if sumExceptFirstIdx == 0 {
		return 0
	}

	// Traverse and check for any pivot in the array
	for idx := 1; idx < n; idx++ {
		leftSum := prefixSum[idx-1]
		rightSum := prefixSum[n-1] - prefixSum[idx]

		if leftSum == rightSum {
			return idx
		}
	}

	// Finally check for right pivot
	if prefixSum[n-2] == 0 {
		return n - 1
	}

	return -1
}
