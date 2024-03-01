package main

func orangesRotting(grid [][]int) int {
	// Key Idea: Run a BFS on the grid whenever you encounter a rotten orange (2)
	// Continue progressing on the path only when you encounter fresh oranges (1)
	// And do not continue further when there is empty space (0)
	// Once BFS is done, count the no. of 1s in the grid if found return -1 otherwise return the total number of minutes
	// IMPORTANT: It is important to the populate the initial queue with the grids of all rotten oranges

	if len(grid) == 0 {
		return -1
	}

	// Check for grids with all empty spaces
	if isEmptyGrid(grid) {
		return 0
	}

	var count int = -1
	var queue [][2]int

	// Populate the initial queue with rotten oranges
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 2 {
				queue = append(queue, [2]int{r, c})
			}
		}
	}

	// If no rotten oranges are found
	if len(queue) == 0 {
		return count
	}

	bfs(grid, queue, &count)

	// Check if still any fresh oranges remain after BFS
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 1 {
				return -1
			}
		}
	}

	return count
}

func isEmptyGrid(grid [][]int) bool {
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] != 0 {
				return false
			}
		}
	}
	return true
}

func bfs(grid [][]int, queue [][2]int, count *int) {
	m, n := len(grid), len(grid[0])

	// 4-directionally adjacent
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(queue) > 0 {
		*count++
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			rottenOrange := queue[0]
			row, col := rottenOrange[0], rottenOrange[1]
			queue = queue[1:]
			for _, dir := range dirs {
				newRow, newCol := row+dir[0], col+dir[1]
				// Traverse the grid only if it is a fresh orange and the index are in-bound
				if newRow >= 0 && newCol >= 0 && newRow < m && newCol < n && grid[newRow][newCol] == 1 {
					queue = append(queue, [2]int{newRow, newCol})
					// Mark the oranges as ROTTEN
					grid[newRow][newCol] = 2
				}
			}
		}
	}
}
