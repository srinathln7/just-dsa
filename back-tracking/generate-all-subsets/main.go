package main

import "fmt"

func subsets(nums []int) [][]int {

	// We use BACKTRACKING approach to solve this problem recursively.
	//  NOTE: we run a recursive function inside a iterative loop

	var result [][]int
	generateSubsets(nums, []int{}, &result, 0)
	return result
}

func generateSubsets(nums, subset []int, result *[][]int, startIdx int) {
	*result = append(*result, append([]int{}, subset...))
	fmt.Printf("result at i= %d => %d\n", startIdx, *result)
	for i := startIdx; i < len(nums); i++ {
		subset = append(subset, nums[i])
		fmt.Printf("subset at i=%d => %d\n", i, subset)
		fmt.Printf("Recurse with start index=%d\n", i+1)
		generateSubsets(nums, subset, result, i+1)

		// Back track by removing the last element of the subset
		subset = subset[0 : len(subset)-1]
		fmt.Printf("Backtracked subset at i=%d => %d \n", i, subset)
	}
}

func main() {
	subsets([]int{1, 2, 3})
}
