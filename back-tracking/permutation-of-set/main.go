package main

func permute(nums []int) [][]int {
	var result [][]int
	used := make(map[int]bool)
	dfs(nums, used, []int{}, &result)
	return result
}

func dfs(nums []int, used map[int]bool, perm []int, result *[][]int) {

	if len(nums) == len(perm) {
		*result = append(*result, append([]int{}, perm...))
		return
	}

	for _, num := range nums {
		if !used[num] {
			perm = append(perm, num)
			used[num] = true

			// Recurse
			dfs(nums, used, perm, result)

			// Backtrack to explore all the other possibilities
			perm = perm[:len(perm)-1]
			used[num] = false
		}
	}
}
