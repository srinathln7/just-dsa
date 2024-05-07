package main

func searchInsert(nums []int, target int) int {

	// Key Idea: Run standard binary search. If target is not found, right pointer `r` represents the closest index
	// to the target while the left pointer `l` will point to the corresponding insert position.

	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if target == nums[mid] {
			return mid
		}

		switch {
		case target < nums[mid]:
			r = mid - 1
		default:
			l = mid + 1
		}
	}

	return l
}
