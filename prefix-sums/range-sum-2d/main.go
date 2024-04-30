package main

type NumMatrix struct {
	prefixSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	prefixSum := make([][]int, m+1)
	for i := range prefixSum {
		prefixSum[i] = make([]int, n+1)
	}

	// Compute cumulative sum for each cell in prefixSum assuming the cell (i,j) is the bottom right of the grid
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			prefixSum[i][j] = prefixSum[i-1][j] + prefixSum[i][j-1] - prefixSum[i-1][j-1] + matrix[i-1][j-1]
		}
	}

	return NumMatrix{prefixSum: prefixSum}
}

func (nm *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return nm.prefixSum[row2+1][col2+1] - nm.prefixSum[row2+1][col1] - nm.prefixSum[row1][col2+1] + nm.prefixSum[row1][col1]
}
