package main

func search(nums []int, target int) int {
	var mid int
	l, r := 0, len(nums)-1

	for l <= r {
		mid = l + (r-l)/2 // Acccounts for integet overflow
		if nums[mid] == target {
			return mid
		}

		if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}
