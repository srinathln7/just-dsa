package main

func trap(height []int) int {

	// Key Idea: Let us just visualize only one unit of block in the middle of the container and assume we have a left most edge
	// with height `l` and rightmost edge with height `r`. The amount of water this container can hold = min(l,r) - height of unit block (= 1).
	// Now here `l` and `r` represent the max height of its `left` and `right` side. In this case, we assumed the block is at the middle of the
	// index. Now since these blocks can be any where, we take the sum of this calculation across all blocks to capture the entire water trapped.

	n := len(height)

	// Base case: No water can be trapped
	if n <= 1 {
		return 0
	}

	// Calculcate the left max and right max at each index
	leftMax, rightMax := make([]int, n), make([]int, n)

	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	var totalWater int
	for i := 0; i < n; i++ {
		totalWater += min(leftMax[i], rightMax[i]) - height[i]
	}

	return totalWater
}
