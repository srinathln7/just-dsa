package main

var dirs [][2]int = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

type Trie struct {
	children [26]*Trie
	word     string
}

func buildTrie(words []string) *Trie {

	root := &Trie{}
	for _, word := range words {
		curr := root

		for i := 0; i < len(word); i++ {
			idx := word[i] - 'a'

			if curr.children[idx] == nil {
				curr.children[idx] = &Trie{}
			}

			// Traverse the trie
			curr = curr.children[idx]
		}

		curr.word = word
	}

	return root
}

func findWords(board [][]byte, words []string) []string {

	root := buildTrie(words)
	if root == nil {
		return nil
	}

	var result []string
	m, n := len(board), len(board[0])

	if m == 0 || n == 0 {
		return nil
	}

	var dfs func(node *Trie, row, col int)
	dfs = func(node *Trie, row, col int) {

		// Out-of-bound or already visited cell
		if row < 0 || row >= m || col < 0 || col >= n || board[row][col] == '#' {
			return
		}

		// Current character
		ch := board[row][col]

		// End of trie - then no use proceeding further
		if node.children[ch-'a'] == nil {
			return
		}

		// Traverse the node
		node = node.children[ch-'a']

		// If the word exists in the board
		if node.word != "" {
			result = append(result, node.word)

			// Mark word as found to avoid duplicates
			node.word = ""
		}

		// Mark the current cell as visited
		board[row][col] = '#'

		// Travere the neighbours in all 4 directions
		for _, dir := range dirs {
			newRow, newCol := row+dir[0], col+dir[1]
			dfs(node, newRow, newCol)
		}

		// Backtrack
		board[row][col] = ch
	}

	// Run the DFS throughout the board to capture all words present in the board from the list of words
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			dfs(root, r, c)
		}
	}

	return result
}
