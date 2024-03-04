package main

func removeDuplicates(nums []int) int {

	// Key Idea: Atmost `k` elements can be unique => We have to compare elements only starting from index `i=k` to previous window
	// element at `i-k`. We develop this solution for a generic `k`. In this case `k=2`

	k := 2
	if len(nums) <= k {
		return len(nums)
	}

	for i := k; i < len(nums); i++ {
		if nums[i] != nums[k-2] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
