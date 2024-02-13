package main

func removeElement(nums []int, val int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		// retreive all the indices which are not equal to val
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}
