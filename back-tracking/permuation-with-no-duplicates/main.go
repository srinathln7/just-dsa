package main

func permuteUnique(nums []int) [][]int {

	// Key Idea: Build a freq table to count the number of characters. Run DFS with backtracking.
	// Whenever the freq. count becomes zero, skip the iteration

	var result [][]int
	freq := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
	}

	dfsWithBT(nums, []int{}, freq, &result)
	return result
}

func dfsWithBT(nums, perm []int, freq map[int]int, result *[][]int) {

	if len(perm) == len(nums) {
		*result = append(*result, append([]int{}, perm...))
		return
	}

	// Without duplicates loop over the map
	for num := range freq {
		if freq[num] == 0 {
			continue
		}

		freq[num]--
		perm = append(perm, num)

		// Recurse
		dfsWithBT(nums, perm, freq, result)

		// Backtrack
		freq[num]++
		perm = perm[:len(perm)-1]
	}

	// If you want with duplicates - Loop over the array
	// for _, num := range nums {
	// 	if freq[num] == 0 {
	// 		continue
	// 	}

	// 	freq[num]--
	// 	perm = append(perm, num)

	// 	// Recurse
	// 	dfsWithBT(nums, perm, freq, result)

	// 	// Backtrack
	// 	freq[num]++

	// }

}
