package main

func setZeroes(matrix [][]int) {
	var firstRowHasZero, firstColHasZero bool
	m := len(matrix)
	n := len(matrix[0])

	// Check if first row has zero
	for i := 0; i < n; i++ {
		if matrix[0][i] == 0 {
			firstRowHasZero = true
			break
		}
	}

	// Check if first col has zero
	for j := 0; j < m; j++ {
		if matrix[j][0] == 0 {
			firstColHasZero = true
			break
		}
	}

	// Keep track of the remaining zeros in the first row and col of the matrix
	// First row keeps track of the cols to be made zero
	// First col keeps track of the rows to be made zero
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// ith Row and jth Col
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
				continue
			}
		}
	}

	// Visit first Row which contains the corresponding cols to be made zero
	for i := 1; i < n; i++ {
		if matrix[0][i] == 0 {
			nullifyCol(matrix, i)
			continue
		}
	}

	// Virst first col which contains the corresponding rows to be made zero
	for j := 1; j < m; j++ {
		if matrix[j][0] == 0 {
			nullifyRow(matrix, j)
			continue
		}
	}

	if firstRowHasZero {
		nullifyFirstRow(matrix)
	}

	if firstColHasZero {
		nullifyFirstCol(matrix)
	}

}

func nullifyRow(matrix [][]int, row int) {
	n := len(matrix[0])
	// Nullify row
	for j := 0; j < n; j++ {
		matrix[row][j] = 0
	}
}

func nullifyCol(matrix [][]int, col int) {
	m := len(matrix)
	// Nullify col
	for i := 0; i < m; i++ {
		matrix[i][col] = 0
	}
}

func nullifyFirstRow(matrix [][]int) {
	for j := 0; j < len(matrix[0]); j++ {
		matrix[0][j] = 0
	}
}

func nullifyFirstCol(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		matrix[i][0] = 0
	}
}
