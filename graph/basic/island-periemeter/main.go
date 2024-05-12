package main

var dirs [][2]int = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func islandPerimeter(grid [][]int) int {

	// Key Idea: Perform a simple DFS. Whenever we encounter a water from land or reach out-of-bound condition
	// then we increment the perimeter

	m, n := len(grid), len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}

	var perimeter int
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 1 {
				dfs(grid, r, c, &perimeter)
				return perimeter
			}
		}
	}

	return perimeter
}

func dfs(grid [][]int, row, col int, perimeter *int) {

	m, n := len(grid), len(grid[0])

	// Out of bound or encountered water
	// Increment the perimeter and then return
	if row < 0 || row >= m || col < 0 || col >= n || grid[row][col] == 0 {
		*perimeter++
		return
	}

	// Already visited
	if grid[row][col] == 2 {
		return
	}

	// Mark the cell as visited
	grid[row][col] = 2

	for _, dir := range dirs {
		newRow, newCol := row+dir[0], col+dir[1]
		dfs(grid, newRow, newCol, perimeter)
	}
}
