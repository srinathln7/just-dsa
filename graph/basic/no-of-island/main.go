package main

func numIslands(grid [][]byte) int {

	// To find the number of islands in the grid, we increment the count by 1
	// whenever we encouter a land marked by "1" and run a BFS on that land cell
	var count int
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == '1' {
				count++

				// It is important to notice here that although we are passing `grid` by value to the bfs function
				// the BFS changes are indeed reflected in the grid because it is represented as a slice of bytes.
				// Slice in Go are nothing but a reference to the underlying arrays => The pointer is implicitly passed
				bfs(grid, r, c)
			}
		}
	}

	return count
}

func bfs(grid [][]byte, row, col int) {
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
		if row < 0 || col < 0 || row >= m || col >= n || grid[row][col] == '0' {
			continue
		}

		// Mark the current cell as visited
		grid[row][col] = '0'

		// Enqueue the neighbours
		queue = append(queue, [2]int{row + 1, col})
		queue = append(queue, [2]int{row - 1, col})
		queue = append(queue, [2]int{row, col + 1})
		queue = append(queue, [2]int{row, col - 1})
	}
}
