package main

func maxAreaOfIsland(grid [][]int) int {
	// To find the max area of an island in the grid, whenever we first encouter a land marked by "1"
	// we run a BFS on that land cell and keep incrementing the area
	var maxArea int

	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 1 {
				var count int
				// It is important to notice here that although we are passing `grid` by value to the bfs function
				// the BFS changes are indeed reflected in the grid because it is represented as a slice of bytes.
				// Slice in Go are nothing but a reference to the underlying arrays => The pointer is implicitly passed
				bfs(grid, r, c, &count)
				maxArea = max(maxArea, count)
			}
		}
	}

	return maxArea
}

func bfs(grid [][]int, row, col int, count *int) {
	m, n := len(grid), len(grid[0])

	// Enqueue the current cell
	var queue [][2]int
	queue = append(queue, [2]int{row, col})

	for len(queue) > 0 {
		// Dequeue the current cell in the grid
		// If you pop the last element out instead of the first element in the arrat, then it would be a DFS
		cell := queue[0]
		row, col = cell[0], cell[1]
		queue = queue[1:]

		// Edge cases - If index is out-of-bound or if the grid is already visited
		if row < 0 || col < 0 || row >= m || col >= n || grid[row][col] == 0 {
			continue
		}

		// Increment the count and mark the current cell as visited
		*count++
		grid[row][col] = 0

		// Enqueue the neighbours
		queue = append(queue, [2]int{row + 1, col})
		queue = append(queue, [2]int{row - 1, col})
		queue = append(queue, [2]int{row, col + 1})
		queue = append(queue, [2]int{row, col - 1})
	}
}
