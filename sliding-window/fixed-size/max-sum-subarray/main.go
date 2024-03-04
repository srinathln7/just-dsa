package main

func maxSubArraySum(nums []int) int {
	maxG := nums[0] // Global maximum
	maxSoFar := nums[0]
	for i := 1; i < len(nums); i++ {
		maxSoFar = max(nums[i], nums[i]+maxSoFar)
		maxG = max(maxG, maxSoFar)
	}
	return maxG
}

func maxSubArray(nums []int) (int, []int) {

	maxG := nums[0]
	maxSoFar := nums[0]
	startIdx, endIdx := 0, 0
	startTracker := 0

	for i, num := range nums {

		// If the current num is greater than the sum of all previous nums + current num
		// then we start a new sub-array
		if num > num+maxSoFar {
			maxSoFar = num
			startTracker = i
		} else {
			// extend the sub array
			maxSoFar = maxSoFar + num
		}

		if maxSoFar > maxG {
			maxG = maxSoFar
			startIdx = startTracker
			endIdx = i
		}
	}

	return maxG, nums[startIdx : endIdx+1]
}
