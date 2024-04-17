package main

import "sort"

func subsetsWithDup(nums []int) [][]int {

	// Key Idea: To start detecting the duplicates, first sort and input array.
	// Run DFS with backtracking algorithm for all possible sets. Include the
	// condition to detect duplicates of subsets

	var result [][]int

	// Sort the input array first to start detecting duplicates
	sort.Ints(nums)

	var dfs func(startIdx int, subset []int)
	dfs = func(startIdx int, subset []int) {

		result = append(result, append([]int{}, subset...))

		for i := startIdx; i < len(nums); i++ {

			// In the code if i > start && nums[i] == nums[i-1], the condition i > start ensures that we are not checking for duplicates at the very beginning
			// of the array. This condition helps us skip duplicates when we are not starting from the beginning of the array but rather traversing it recursively.

			// Consider the following example: [1,2,2]
			// When we encounter the first `2`, we don't want to consider it as a duplicate because it's the first occurrence of `2` in the current subset.
			// However, when we encounter the second `2`, we want to skip it because it's a duplicate of the previous `2`.
			// IMPORTANT: By using the condition `i > start`, we ensure that we are only comparing the current element with its previous element within the current subset
			// and not with elements from previous subsets. This helps in skipping duplicates effectively and generating only unique subsets.
			if i > startIdx && nums[i] == nums[i-1] {
				continue
			}

			// Decision to include the current element
			subset = append(subset, nums[i])

			// Recurse to get the subset of the set [i]
			dfs(i+1, subset)

			// Backtrack
			subset = subset[:len(subset)-1]
		}
	}

	dfs(0, []int{})
	return result
}
