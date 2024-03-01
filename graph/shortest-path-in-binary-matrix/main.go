package main

func shortestPathBinaryMatrix(grid [][]int) int {

	if len(grid) == 0 {
		return -1
	}

	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}

	var pathlength int
	var queue [][2]int
	queue = append(queue, [2]int{0, 0})

	// Since we can move in 8 directions
	dirs := [8][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

	for len(queue) > 0 {

		// Increment path length after every level i.e. trying to move in all possible directions
		pathlength++
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			row, col := queue[0][0], queue[0][1]
			queue = queue[1:]

			// If the destination is reached
			if row == n-1 && col == n-1 {
				return pathlength
			}

			// Try moving in all dirs
			for _, dir := range dirs {
				newRow, newCol := row+dir[0], col+dir[1]

				// If out-of-bound or grid already visited
				if newRow < 0 || newCol < 0 || newRow >= n || newCol >= n || grid[newRow][newCol] == 1 {
					continue
				}

				queue = append(queue, [2]int{newRow, newCol})

				// IMPORTANT: Mark the grid as visited here immediately in the current iteration
				// This will prevent visting grids that have already been visited and ensure we do not run out of memory
				grid[newRow][newCol] = 1
			}
		}
	}

	// No clear path has been found
	return -1
}
