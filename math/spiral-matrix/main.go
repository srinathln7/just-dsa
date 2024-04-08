package main

func spiralOrder(matrix [][]int) []int {

	// Key Idea: Keep track of all the four boundaries: top, bottom, left, and right.
	// Move in a spiral fashion

	m, n := len(matrix), len(matrix[0])
	top, bottom, left, right := 0, m-1, 0, n-1

	idx, result := 0, make([]int, m*n)
	for top <= bottom && left <= right {

		// Move along the top to the right
		for j := left; j <= right; j++ {
			result[idx] = matrix[top][j]
			idx++
		}
		top++

		// Move from top to bottom downwards along the right most column
		for j := top; j <= bottom; j++ {
			result[idx] = matrix[j][right]
			idx++
		}
		right--

		// Move from right to left along the bottom row
		if top <= bottom {
			for j := right; j >= left; j-- {
				result[idx] = matrix[bottom][j]
				idx++
			}
			bottom--
		}

		// Move from bottom to top along the left most column
		if left <= right {
			for j := bottom; j >= top; j-- {
				result[idx] = matrix[j][left]
				idx++
			}
			left++
		}

	}

	return result
}
