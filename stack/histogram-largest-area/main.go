package main

type Stack struct {
	height int
	idx    int
}

func largestRectangleArea(heights []int) int {

	// Key Idea: Use a bottom-up approach. There are in total 3 sceanarios - Histograms in increasing order of size which is extendable
	// decresing order of size - not extendable and same size which is extendable

	var maxArea int
	var stack []Stack
	n := len(heights)
	for i, height := range heights {

		start := i

		// Non-Extendable case - Pop out from stack
		for len(stack) > 0 && stack[len(stack)-1].height > height {
			idx, height := stack[len(stack)-1].idx, stack[len(stack)-1].height
			stack = stack[:len(stack)-1]
			maxArea = max(maxArea, height*(i-idx))
			start = idx // Extend the min. width to the left
		}

		stack = append(stack, Stack{height: height, idx: start})
	}

	// Compare current maxArea with the rest of the areas in the stack that we extended all the way till the last index
	for _, element := range stack {
		maxArea = max(maxArea, element.height*(n-element.idx))
	}

	return maxArea
}
