package main

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	// Recursively sort the left and right tree and then merge back finally
	mid := len(nums) / 2
	left := sortArray(nums[:mid])
	right := sortArray(nums[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	// Declare two pointers
	i, j := 0, 0

	results := make([]int, 0, len(left)+len(right))
	for i < len(left) && j < len(right) {

		if left[i] <= right[j] {
			results = append(results, left[i])
			i++
		} else {
			results = append(results, right[j])
			j++
		}
	}

	// Deal with left overs
	results = append(results, left[i:]...)
	results = append(results, right[j:]...)

	return results
}
