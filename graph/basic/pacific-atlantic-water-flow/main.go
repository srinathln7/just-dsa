package main

// Capture all 4 dirs
var dirs [][2]int = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func pacificAtlantic(heights [][]int) [][]int {

	// Key Idea: To find the list of co-ordinates denotes that rain water can flow from cell (ri, ci) to both the Pacific and Atlantic oceans
	// we turn things other way around and try to find all cells that are reachable from both pacific and atlantic ocean. We form a boolean
	// array to cupture the cells reachable from pacific ocean and atlantic ocean respectively and then look for cells that are reachable from
	// both the cells

	m, n := len(heights), len(heights[0])
	if m == 0 || n == 0 {
		return nil
	}

	pacificReachable := make([][]bool, m)
	for i := 0; i < m; i++ {
		pacificReachable[i] = make([]bool, n)
	}

	atlanticReachable := make([][]bool, m)
	for i := 0; i < m; i++ {
		atlanticReachable[i] = make([]bool, n)
	}

	// Runs DFS on the left and right edges
	for i := 0; i < m; i++ {

		// Left Edge - Check how many cells are reachable from pacific ocean
		dfs(heights, i, 0, pacificReachable)

		// Right Edge - Check how many cells are reachable from atlantic ocean
		dfs(heights, i, n-1, atlanticReachable)
	}

	// Run DFS on the top and bottom edges
	for i := 0; i < n; i++ {

		// Top Edge - Check how many cells are reachable from pacific ocean
		dfs(heights, 0, i, pacificReachable)

		// Bottom Edge - Check how many cells are reachable from atlantic ocean
		dfs(heights, m-1, i, atlanticReachable)
	}

	var result [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacificReachable[i][j] && atlanticReachable[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

func dfs(heights [][]int, r, c int, reachable [][]bool) {

	m, n := len(heights), len(heights[0])

	if reachable[r][c] {
		return
	}

	reachable[r][c] = true
	for _, dir := range dirs {
		nr, nc := r+dir[0], c+dir[1]

		// IMPORTANT: Water can flow from ocean into the neighbouring cells only if
		// the height of the current cell is lesser than or equal to the neighbouring cells
		// Note, we reverse the condition here since we are trying to find the cells that are
		// reachable from ocean to the islands.
		if nr >= 0 && nr < m && nc >= 0 && nc < n && heights[r][c] <= heights[nr][nc] {
			dfs(heights, nr, nc, reachable)
		}
	}

}
