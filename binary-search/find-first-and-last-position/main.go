package main

func searchRange(nums []int, target int) []int {

	// Key Idea: Since we need to output the start and end position of the target value in the index and do this in O(log.n) complexity
	// we need to modify the classic binary search algorithm. We need to run the binary search algorithm twice, one to find the left-most
	// and the other to find the right-most index

	leftIdx, rightIdx := binarySearch(nums, target, true), binarySearch(nums, target, false)
	return []int{leftIdx, rightIdx}
}

func binarySearch(nums []int, target int, isLeft bool) int {

	res := -1
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if target == nums[mid] {

			// Capture the index first
			res = mid

			switch isLeft {
			// Check for the left-most position i.e. first position of target in the array
			// So we trim the right-most portion of the array
			case true:
				r = mid - 1

				// Check for the right-most position i.e. last position of target in the array
				// So we trim the left-most portion of the array
			default:
				l = mid + 1
			}
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}

		// MISTAKE : `default` here also covers the equality case which we already addressed above. Hence
		// better to do it via else if as above or mention `case target < nums[mid]` seperately below
		// switch {
		//     case target > nums[mid]: l = mid+1
		//     default: r = mid-1
		// }

	}

	return res
}
