package main

func combine(n int, k int) [][]int {
	var result [][]int
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	dfsWithBT(nums, []int{}, k, 0, &result)
	return result
}

func dfsWithBT(nums, comb []int, k, startIdx int, result *[][]int) {

	if len(comb) == k {
		*result = append(*result, append([]int{}, comb...))
	}

	for i := startIdx; i < len(nums); i++ {

		// Form the combination slice
		comb = append(comb, nums[i])

		// Recurse
		dfsWithBT(nums, comb, k, i+1, result)

		// Backtrack
		comb = comb[:len(comb)-1]
	}

}
