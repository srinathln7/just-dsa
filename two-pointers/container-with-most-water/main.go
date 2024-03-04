package main

func maxArea(height []int) int {

	// Key Idea: We need to maintain two pointers l, r initially set at the beginning and end of array and calculate the area.
	// We move the shortest height pointer inwards and update the maximum area in each iteration.

	var length, width, area, maxArea int
	l, r := 0, len(height)-1
	for l < r {
		length = r - l
		width = min(height[l], height[r])
		area = length * width
		maxArea = max(maxArea, area)

		// Move the pointer pointing to the lower height of the two containers b/w left and right to increase the chance of finding max area
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxArea
}
