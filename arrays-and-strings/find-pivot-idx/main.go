package main

func pivotIndex(nums []int) int {
	var sum, leftSum int
	for _, value := range nums {
		sum += value
	}
	for i := range nums {
		if leftSum == sum-leftSum-nums[i] {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}
