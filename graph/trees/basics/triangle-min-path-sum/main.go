package main

func minimumTotal(triangle [][]int) int {

	m := len(triangle)
	if m == 1 {
		return triangle[0][0]
	}

	// Memoization Table
	dp := make(map[[2]int]int)

	var dfs func(int, int) int
	dfs = func(rowIdx, colIdx int) int {
		if rowIdx >= m || colIdx > rowIdx {
			return 0
		}

		if val, exists := dp[[2]int{rowIdx, colIdx}]; exists {
			return val
		}

		// Decision tree with (rowIdx, colIdx) as its root
		leftSum := triangle[rowIdx][colIdx] + dfs(rowIdx+1, colIdx)    // Move to idx `i` on the next row
		rightSum := triangle[rowIdx][colIdx] + dfs(rowIdx+1, colIdx+1) // Move to idx `i+1` on the next row

		dp[[2]int{rowIdx, colIdx}] = min(leftSum, rightSum)
		return dp[[2]int{rowIdx, colIdx}]
	}

	return dfs(0, 0)
}
