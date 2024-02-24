package main

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	// Define mapping to convert 2D array to 1D array
	l, r := 0, (m*n - 1)

	// Rem: 1d index = col_index + (row_index * n)

	var midIndex, midValue int
	for l <= r {
		midIndex = l + (r-l)/2

		// Convert 1D mapping back to 2D mapping
		midValue = matrix[midIndex/n][midIndex%n]

		if target == midValue {
			return true
		}

		if target < midValue {
			r = midIndex - 1
		} else {
			l = midIndex + 1
		}
	}

	return false
}
