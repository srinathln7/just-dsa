package main

func findDisappearedNumbers(nums []int) []int {

	// Key Idea: Since the values in the array range from 1 to n, it is an indication that we might have to use
	// these values itself as an index. To do this We mark all the values we encountered negative and then range
	// from 1 to n and check for positive numbers that represents the missing numbers

	var result []int
	n := len(nums)

	// Mark those numbers
	for i := 0; i < n; i++ {
		idx := abs(nums[i]) - 1
		if nums[idx] > 0 {
			nums[idx] = -nums[idx] // Mark the numbers as negative
		}
	}

	// Check for positive numbers and append its indices
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
