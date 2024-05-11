package main

func removeElement(nums []int, val int) int {

	// Key Idea: Use two pointer index approach on the same array

	var count int
	for i := 0; i < len(nums); i++ {
		// Retreive all the indices which are not equal to val
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}
