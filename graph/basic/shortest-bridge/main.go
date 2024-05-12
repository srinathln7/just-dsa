package main

var dirs = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func shortestBridge(grid [][]int) int {

	// Key Idea: We essentially have to form the shortest bridge between the two islands. Note that they are in-total only two islands.
	// So, let us first visit the first islands using DFS/BFS and then from all the cells in this visited island, let us run a BFS
	// to the other island. This will determine the shortest path which in essence will give us the number of zeros (water) we need to
	// flip to 1.

	n := len(grid)
	if n == 0 {
		return 0
	}

	// Form a queue to run the BFS first island to second island
	var queue [][2]int

	// First let us run a DFS
	// Label representing the outer loop

Outer:
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 1 {
				dfs(grid, r, c, &queue)
				break Outer // break the outer loop
			}
		}
	}

	// Now, once DFS is run, we have visited our first island whose cells are marked by 2.
	// REM: Slice are references to underlying arrays and hence change made in the other function will be automatically reflected in this function.
	// Only when operation that changes the grid's metadata (such as append etc) will not be reflected. Hence, that is why queue is passed as reference

	// Run BFS to find the shortest path to the second island
	steps := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			x, y := queue[i][0], queue[i][1]
			for _, dir := range dirs {
				nx, ny := x+dir[0], y+dir[1]

				// Out-of-bound or visiting first island again
				if nx >= 0 && nx < n && ny >= 0 && ny < n && grid[nx][ny] != 2 {

					// If you have found the second island from the first island, then return immediately
					if grid[nx][ny] == 1 {
						return steps
					}
					grid[nx][ny] = 2
					queue = append(queue, [2]int{nx, ny})
				}
			}
		}
		queue = queue[size:]
		steps++
	}
	return -1
}

// DFS to mark the cells of the first island
func dfs(grid [][]int, x, y int, queue *[][2]int) {
	n := len(grid)
	if x < 0 || x >= n || y < 0 || y >= n || grid[x][y] != 1 {
		return
	}
	grid[x][y] = 2
	*queue = append(*queue, [2]int{x, y})
	for _, dir := range dirs {
		dfs(grid, x+dir[0], y+dir[1], queue)
	}
}
