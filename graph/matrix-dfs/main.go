package main

import "fmt"

var grid = [][]int{
	{0, 0, 0, 0},
	{1, 1, 0, 0},
	{0, 0, 0, 1},
	{0, 1, 0, 0},
}

// Key Idea: To get the number of unique paths, we deploy a DFS + Backtracking algorithm
func dfs(grid [][]int, r, c int, visit map[[2]int]bool) int {
	rows, cols := len(grid), len(grid[0])
	fmt.Printf("(row,col) => (%d,%d)\n", r, c)
	if r < 0 || c < 0 || r == rows || c == cols || visit[[2]int{r, c}] || grid[r][c] == 1 {
		return 0
	}

	if r == rows-1 && c == cols-1 {
		return 1
	}

	visit[[2]int{r, c}] = true

	count := 0
	count += dfs(grid, r+1, c, visit)
	count += dfs(grid, r-1, c, visit)
	count += dfs(grid, r, c+1, visit)
	count += dfs(grid, r, c-1, visit)

	// Back track the visited path to find more unique paths
	delete(visit, [2]int{r, c})
	return count
}

func main() {
	visit := make(map[[2]int]bool)
	result := dfs(grid, 0, 0, visit)
	println(result)
}
