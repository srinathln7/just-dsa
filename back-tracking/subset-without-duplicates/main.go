package main

import "fmt"

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
	fmt.Println("RESULT", *result)

	for num := range freq {
		// Skip forming the subset if you have vanished all the elements with in the current subset
		// or the last element in your formed subset is greater than current num in the map since
		// you would have already formed those subsets
		if freq[num] == 0 || len(subset) > 0 && subset[len(subset)-1] > num {
			fmt.Printf("Duplicate subset %d occured when num= %d\n", subset, num)
			continue
		}

		// Append the current num to the subset and decrement the count in the freq. table
		subset = append(subset, num)
		fmt.Printf("Subset %d at num=%d\n", subset, num)
		freq[num]--

		dfsWithBT(nums, subset, freq, result)

		// Backtrack
		subset = subset[:len(subset)-1]
		freq[num]++
	}
}

func main() {

	nums := []int{4, 4, 4, 1, 4}
	fmt.Println(nums)
	fmt.Println(subsetsWithDup(nums))
}
