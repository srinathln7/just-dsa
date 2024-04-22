package main

var dirs = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func solve(board [][]byte) {

	// Key Idea: To solve this problem, we use reverse thinking approach where to find the regions that are 4-directionally surronded
	// by `X`, we try to find all regions that are not 4-directionally surronded by `X`. This region is easier to get since we can
	// find this easily by running DFS along the borders. We mark this regions by a temporary marker `T` and later once the DFS is
	// complete we change this to `0`

	m, n := len(board), len(board[0])
	if m == 0 || n == 0 {
		return
	}

	var dfs func(int, int)
	dfs = func(row, col int) {
		if row < 0 || row >= m || col < 0 || col >= n || board[row][col] != 'O' {
			return
		}

		// Mark the unsurronded region with a temporary marker
		board[row][col] = 'T'

		for _, dir := range dirs {
			dfs(row+dir[0], col+dir[1])
		}
	}

	// Run DFS on top and bottom rows
	for i := 0; i < n; i++ {
		// Top row
		dfs(0, i)

		// Bottom row
		dfs(m-1, i)
	}

	// Run DFS on left and right column

	for i := 0; i < m; i++ {
		// Left column
		dfs(i, 0)

		// Right column
		dfs(i, n-1)
	}

	// Change all temporary markers to `O` and all other to `X`
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			switch board[r][c] {
			case 'T':
				board[r][c] = 'O'
			default:
				board[r][c] = 'X'
			}
		}
	}

}
