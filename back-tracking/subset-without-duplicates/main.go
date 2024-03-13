func subsetsWithDup(nums []int) [][]int {

	var result [][]int

	// Lets create a frequency map
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	dfsWithBT(nums, []int{}, freq, &result)
	return result
}

func dfsWithBT(nums, subset []int, freq map[int]int, result *[][]int) {
	*result = append(*result, append([]int{}, subset...))
	for num := range freq {
		// Skip forming the subset if you have vanished all the elements with in the current subset
		// or the last element in your formed subset is greater than current num in the map since you would
		// have already formed those subsets. BUT WHY we do we skip if the last element in our current subset is
		// greater than the num and HOW does it ensure those subsets are already formed?
		// This is because, no matter how unsorted your array is, we will always form  all subsets in SORTED order.
		// This means if there is ever a case, where the last element in our subset is greater than the current number
		// in the map, those subsets are already been formed. Hence, we skip the iteration to avoid duplicates.
		// And avoiding to form subsets when `freq[num] == 0` in the current recursive stack frame is STRAIGHT FORWARD.
		if freq[num] == 0 || len(subset) > 0 && subset[len(subset)-1] > num {
			continue
		}

		// Append the current num to the subset and decrement the freq. count
		subset = append(subset, num)
		freq[num]--

		dfsWithBT(nums, subset, freq, result)

		// Backtrack by popping the last element of the subset and increment the freq. count
		subset = subset[:len(subset)-1]
		freq[num]++
	}
}