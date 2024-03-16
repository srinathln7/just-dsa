package main

func rob(nums []int) int {

	// Key Idea: When we reach say a house `i`, we keep track of the maximum amount of money
	// that we would have had if we had stolen until house `i-1` and `i-2`

	// We adopt a BOTTOM-UP approach
	// For example: -> Only one house => max = nums[0]
	// Two houses => max = max(nums[0], nums[1])
	// Three houses => max = max(nums[0] + nums[2] , nums[1])
	// Four houses => max = max(nums[0]+ nums[2], nums[0]+nums[3], nums[1]+nums[3])
	// ...

	// Edge cases
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	}

	var prevMax, currMax int
	prevMax = nums[0]
	currMax = max(nums[0], nums[1])

	n := len(nums)
	for i := 2; i < n; i++ {
		tmp := currMax
		currMax = max(currMax, nums[i]+prevMax)
		prevMax = tmp
	}
	return currMax
}
