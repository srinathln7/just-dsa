package main

func generate(numRows int) [][]int {

	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)
	}

	result[0][0] = 1
	if numRows == 1 {
		return result
	}

	result[1][0], result[1][1] = 1, 1
	if numRows == 2 {
		return result
	}

	for i := 2; i < numRows; i++ {
		n := len(result[i])
		result[i][0], result[i][n-1] = 1, 1
		for j := 1; j < n-1; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}

	return result
}
