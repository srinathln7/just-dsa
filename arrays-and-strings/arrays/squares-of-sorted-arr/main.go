package main

func sortedSquares(nums []int) []int {
	left := 0
	right := len(nums) - 1
	res := make([]int, 0, right+1)

	for left <= right {
		if nums[left]*nums[left] < nums[right]*nums[right] {
			res = append(res, nums[right]*nums[right])
			right--
		} else {
			res = append(res, nums[left]*nums[left])
			left++
		}
	}
	reverseSlice(res)
	return res
}

func reverseSlice(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
