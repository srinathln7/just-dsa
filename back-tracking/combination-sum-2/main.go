package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {

	// Key Idea: Use DFS with BT
	var result [][]int

	// To avoid duplicates
	sort.Ints(candidates)

	dfsWithBT(candidates, target, 0, []int{}, &result)

	return result
}

func dfsWithBT(candidates []int, target int, startIdx int, comb []int, result *[][]int) {

	if target == 0 {
		*result = append(*result, append([]int{}, comb...))
	}

	if target < 0 {
		return
	}

	for i := startIdx; i < len(candidates); i++ {

		// In the code if i > start && nums[i] == nums[i-1], the condition i > start ensures that we are not checking for duplicates at the very beginning
		// of the array. This condition helps us skip duplicates when we are not starting from the beginning of the array but rather traversing it recursively.

		// Consider the following example: [1,2,2]
		// When we encounter the first `2`, we don't want to consider it as a duplicate because it's the first occurrence of `2` in the current subset.
		// However, when we encounter the second `2`, we want to skip it because it's a duplicate of the previous `2`.
		// IMPORTANT: By using the condition `i > start`, we ensure that we are only comparing the current element with its previous element within the current subset
		// and not with elements from previous subsets. This helps in skipping duplicates effectively and generating only unique subsets.
		if i > startIdx && candidates[i] == candidates[i-1] {
			continue
		}

		// Decision to include current num
		comb = append(comb, candidates[i])

		// Run DFS with `target = target-candidates[i]`
		// Since we can use the number only once in the combination -> we specify the start index for this recursion
		// as `i+1`
		dfsWithBT(candidates, target-candidates[i], i+1, comb, result)

		// Back track
		comb = comb[:len(comb)-1]
	}

}
