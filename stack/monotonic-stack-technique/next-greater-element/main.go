package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	// Key Idea: We use a monotonic stack technique to solve this problem.

	result := make([]int, len(nums1))
	n1Set := make(map[int]int)
	for i, num := range nums1 {
		n1Set[num] = i
		result[i] = -1 // Default value
	}

	var stack []int
	for _, num := range nums2 {

		// For a non-empty stack and if the incoming value is greater than the top of the stack
		// it means we found the element in the right that is just greater than value
		for len(stack) > 0 && num > stack[len(stack)-1] {

			top := stack[len(stack)-1]

			// Pop from the stack
			stack = stack[:len(stack)-1]

			// Get the corresponding index of it in the `n1Set`
			idx := n1Set[top]

			// Form the result
			result[idx] = num
		}

		// If the incoming num is not present in the subset array, then we skip appending to the stack
		// since we dont have to find the max. for it
		if _, exists := n1Set[num]; !exists {
			continue
		}

		stack = append(stack, num)
	}

	return result
}
