package main

func productExceptSelf(nums []int) []int {

	// Key Idea: Is to maintain two pointers `prefix` and `suffix` that keeps track of the product of all
	// elements to the left and right of it except itself

	n := len(nums)
	result := make([]int, n)

	// First , calculate the pdt of all the elements to the left of the index
	prefix := 1
	for i := 0; i < n; i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	// Calculate the pdt of all the elements to the right of the index
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}
