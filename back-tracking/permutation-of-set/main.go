package main

func permute(nums []int) [][]int {

	// Key Idea : Once we have used an element in forming the subset, we cannot use it again
	// and the length of the subset must always equal the length of the nums array for a valid permuation
	// To do this, we run a DFS with back tracking algorithm along with a hash map to keep track of the elements
	// we have used in the current recursive stack frame.

	var result [][]int
	used := make(map[int]bool)
	dfsWithBT(nums, []int{}, used, &result)
	return result
}

func dfsWithBT(nums, subset []int, used map[int]bool, result *[][]int) {

	// Valid permuation is only when the subset length matches the original ararys length
	if len(nums) == len(subset) {
		*result = append(*result, append([]int{}, subset...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if !used[nums[i]] {
			// Form the initial subset
			used[nums[i]] = true
			subset = append(subset, nums[i])

			// Recurse
			dfsWithBT(nums, subset, used, result)

			// Backtrack
			used[nums[i]] = false
			subset = subset[:len(subset)-1]
		}
	}
}
