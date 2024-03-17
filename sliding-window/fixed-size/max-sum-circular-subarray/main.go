package main

func maxSubarraySumCircular(nums []int) int {
	// Key Idea: Since the array is circular which means we can essentially stitch both ends of the corner
	// we can also keep track of the minimum sum and see if stiching the two ends indeed result in a
	// greater sum than the one we obtain from the kaden's algorithm

	var minSumForward, maxSumForward, maxSumCircular int
	maxSumForward = getMaxSumSubarray(nums)
	minSumForward = getMinSumSubarray(nums)

	var sumTotal int
	for i := 0; i < len(nums); i++ {
		sumTotal += nums[i]
	}
	maxSumCircular = sumTotal - minSumForward

	// Account for the array with only negative values
	if maxSumForward < 0 {
		return maxSumForward
	}

	return max(maxSumForward, maxSumCircular)
}

func getMaxSumSubarray(nums []int) int {
	var maxSoFar, sumEndingHere int = nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		sumEndingHere = max(nums[i], nums[i]+sumEndingHere)
		maxSoFar = max(maxSoFar, sumEndingHere)
	}

	return maxSoFar
}

func getMinSumSubarray(nums []int) int {
	var minSoFar, sumEndingHere int = nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		sumEndingHere = min(nums[i], nums[i]+sumEndingHere)
		minSoFar = min(minSoFar, sumEndingHere)
	}
	return minSoFar
}
