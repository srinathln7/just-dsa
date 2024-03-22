package main

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	// Define mapping to convert 2D array to 1D array
	l, r := 0, (m*n - 1)

	// Rem: 1d index = col_index + (row_index * n)

	var midIdx, midValue int
	for l <= r {
		midIdx = l + (r-l)/2

		// Convert 1D mapping back to 2D mapping
		midValue = matrix[midIdx/n][midIdx%n]

		if target == midValue {
			return true
		}

		if target < midValue {
			r = midIdx - 1
		} else {
			l = midIdx + 1
		}
	}

	return false
}
