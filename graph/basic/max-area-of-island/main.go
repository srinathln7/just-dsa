package main

func maxAreaOfIsland(grid [][]int) int {

	m, n := len(grid), len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}

	var bfs func(int, int) int
	bfs = func(r, c int) int {
		area := 0
		dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

		queue := [][2]int{{r, c}}

		// Mark the current cell as visited
		grid[r][c] = 0

		for len(queue) > 0 {

			// Currently visiting cell
			cell := queue[0]
			row, col := cell[0], cell[1]

			// Dequeue
			queue = queue[1:]

			area++

			for _, dir := range dirs {
				newRow, newCol := row+dir[0], col+dir[1]

				if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && grid[newRow][newCol] == 1 {

					// Mark the cell as visited
					grid[newRow][newCol] = 0

					// Enqueue
					queue = append(queue, [2]int{newRow, newCol})
				}
			}
		}

		return area
	}

	var area, maxArea int
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 1 {
				area = bfs(r, c)
			}
			maxArea = max(maxArea, area)
		}
	}

	return maxArea
}
