package main

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	// Binary search to find the appropriate row
	left, right := 0, m-1
	for left <= right {
		mid := left + (right-left)/2
		if matrix[mid][0] <= target && target <= matrix[mid][n-1] {
			// Found the row where target might exist
			// Perform binary search within this row
			return binarySearch(matrix[mid], target)
		} else if matrix[mid][0] > target {
			right = mid - 1 // Search in the upper half
		} else {
			left = mid + 1 // Search in the lower half
		}
	}
	return false
}

// Binary search within a sorted row
func binarySearch(row []int, target int) bool {
	left, right := 0, len(row)-1
	for left <= right {
		mid := left + (right-left)/2
		if row[mid] == target {
			return true
		} else if row[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
