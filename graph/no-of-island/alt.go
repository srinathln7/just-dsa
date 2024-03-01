package main

func numIslandsDFS(grid [][]byte) int {

	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	// To find the number of islands in the grid, we increment the count by 1
	// whenever we encouter a land marked by "1" and run a DFS on that land cell

	// We either have to run BFS or DFS because adjacent cells that are `1` still are a part of the same island
	// and not a different one. So every time we encounter neighbours with `1` we mark them as visited and then
	// no longer run DFS/BFS on them.

	var count int
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == '1' {
				count++

				// It is important to notice here that although we are passing `grid` by value to the bfs function
				// the BFS changes are indeed reflected in the grid because it is represented as a slice of bytes.
				// Slice in Go are nothing but a reference to the underlying arrays => The pointer is implicitly passed
				dfs(grid, r, c)
			}
		}
	}

	return count
}

func dfs(grid [][]byte, row, col int) {

	m := len(grid)
	n := len(grid[0])

	if row < 0 || col < 0 || row >= m || col >= n || grid[row][col] == '0' {
		return
	}

	// Mark the current cell as visited
	grid[row][col] = '0'

	// Recurse
	dfs(grid, row+1, col)
	dfs(grid, row-1, col)
	dfs(grid, row, col+1)
	dfs(grid, row, col-1)
}
