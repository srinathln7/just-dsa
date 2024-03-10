package main

import "fmt"

func subsetsWithDup(nums []int) [][]int {

	var result [][]int

	// Lets create a frequency map
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	fmt.Println("Frequency map", freq)
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
		// This is because, no matter how unsorted your array is, when you store its freq. in the hash map and then
		// loop over it, it will always be printed in the SORTED order. For eg: nums = [4,4,4,1,4,1,2] has a
		// frequency map map[1:2 2:1 4:4] => This means all subsets we form will always be in SORTED order.
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
