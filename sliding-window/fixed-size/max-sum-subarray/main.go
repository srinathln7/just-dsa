package main

func maxSumSubArray(nums []int) int {
	maxSoFar := nums[0]
	sumEndingHere := nums[0]

	for i := 1; i < len(nums); i++ {
		sumEndingHere = max(nums[i], nums[i]+sumEndingHere)
		maxSoFar = max(maxSoFar, sumEndingHere)
	}

	return maxSoFar
}

func maxSubArray(nums []int) (int, []int) {

	maxSoFar := nums[0]
	sumEndingHere := nums[0]
	startIdx, endIdx := 0, 0
	startTracker := 0

	for i, num := range nums {

		// If the current num is greater than the sum of all previous nums + current num
		// then we start a new sub-array
		if num > num+sumEndingHere {
			sumEndingHere = num
			startTracker = i
		} else {
			// extend the sub array
			sumEndingHere = sumEndingHere + num
		}

		if sumEndingHere > maxSoFar {
			maxSoFar = sumEndingHere
			startIdx = startTracker
			endIdx = i
		}
	}

	return maxSoFar, nums[startIdx : endIdx+1]

}
