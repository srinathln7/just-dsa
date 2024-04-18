package main

func exist(board [][]byte, word string) bool {

	// Key Idea: To find out if a word exists in a board along horizontal and vertical paths, we employ a DFS
	// with backtracking algorithm starting from the list of cells whose starting value matched with the starting
	// character of the word. If a path is not found, we backtrack and explore all paths until all the paths are
	// exhausted.

	m, n := len(board), len(board[0])
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if board[r][c] == word[0] && dfsBT(board, word, r, c, 0) {
				return true
			}
		}
	}

	return false
}

func dfsBT(board [][]byte, word string, row, col, idx int) bool {
	m, n := len(board), len(board[0])

	if idx == len(word) {
		return true
	}

	if row < 0 || row >= m || col < 0 || col >= n || board[row][col] != word[idx] {
		return false
	}

	// Mark the current cell as visited in the current stack frame
	tempCh := board[row][col]
	board[row][col] = '0'

	// Recursion
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, dir := range dirs {
		newRow, newCol := row+dir[0], col+dir[1]
		if dfsBT(board, word, newRow, newCol, idx+1) {
			return true
		}
	}

	// Backtrack - We need backtracking here because still in the current recursive stack frame
	// we can still find a valid possible combination just by moving one-step back and explore
	// all possible paths
	board[row][col] = tempCh

	return false
}
