package main

func subsets(nums []int) [][]int {
	var res [][]int
	var subset []int

	var dfs func(int)
	dfs = func(i int) {
		if i == len(nums) {
			// MISTAKE : Refer pre-req.go
			// res = append(res1, subset)
			res = append(res, append([]int{}, subset...))
			return
		}

		// Decision to include nums[i]
		subset = append(subset, nums[i])
		dfs(i + 1)

		// Decision NOT to include nums[i] - pop the element out of the stack and run dfs
		subset = subset[:len(subset)-1]
		dfs(i + 1)
	}

	dfs(0)
	return res
}
