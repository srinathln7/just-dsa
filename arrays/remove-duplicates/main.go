package main

func removeDuplicates(nums []int) int {
	// `j` represents the index where there is no deplicate element
	var j int = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
