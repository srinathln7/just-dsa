package main

func isValidSudoku(board [][]byte) bool {

	// Key Idea: We need to validate every row and every column for a possible repitition.
	// We can use a hash map for this and reset it after every iteration

	// Validate rows
	if !isValidBoard(board) {
		return false
	}

	// Validate all 3*3 subboxes
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			if !isValidSubBox(board, i, j) {
				return false
			}
		}
	}

	// Validate all cols. To do this, we first transpose the board
	for i := 0; i < 9; i++ {
		for j := i; j < 9; j++ {
			board[i][j], board[j][i] = board[j][i], board[i][j]
		}
	}

	return isValidBoard(board)
}

func isValidBoard(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		freq := make(map[byte]struct{})
		for j := 0; j < 9; j++ {
			ch := board[i][j]
			if ch == '.' {
				continue
			}

			if _, exists := freq[ch]; exists {
				return false
			}
			freq[ch] = struct{}{}
		}
	}

	return true
}

func isValidSubBox(board [][]byte, row, col int) bool {
	freq := make(map[byte]struct{})
	for i := row; i < row+3; i++ {
		for j := col; j < col+3; j++ {
			ch := board[i][j]
			if ch == '.' {
				continue
			}

			if _, exists := freq[ch]; exists {
				return false
			}
			freq[ch] = struct{}{}
		}
	}
	return true
}
