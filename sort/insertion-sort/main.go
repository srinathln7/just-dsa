package main

func insertionSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	// Key Idea: To move all elements behind index `i` (nums[0..n-1]) that are greater than the element at index `i` by position 1
	for i := 1; i < len(nums); i++ {
		key := nums[i] // Element at index i
		j := i - 1
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}

	return nums
}
