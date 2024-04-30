# Word Search 2 

Given an \(m \times n\) board of characters and a list of strings `words`, return all words on the board.

Each word must be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.

## Intuition

To find all words on the board, we can perform a depth-first search (DFS) starting from each cell on the board. At each cell, we explore all possible directions to form words. We can optimize the search process using a trie data structure to efficiently check if a prefix or word exists in the list of words.

## Approach

1. **Construct Trie**: Build a trie data structure using the list of words. Each node in the trie represents a character in a word, and the edges represent transitions between characters.

2. **DFS on Board**: Perform a depth-first search (DFS) starting from each cell on the board. At each cell, explore all possible directions (up, down, left, right) to form words.

3. **Backtracking**: During DFS, maintain a visited set to track cells that have been visited. Append the characters of the current cell to form a prefix and check if it exists in the trie. If the prefix exists, continue DFS exploration. If the prefix forms a word, add it to the result list and remove it from the trie to avoid duplicate words.

4. **Explore All Cells**: Iterate over all cells on the board and start DFS from each cell.

## Time Complexity

- Construct Trie: \(O(WL)\), where \(W\) is the total number of words and \(L\) is the average length of words. Constructing the trie requires iterating over all words and characters.
- DFS on Board: \(O(M \times N \times 4^L)\), where \(M\) is the number of rows, \(N\) is the number of columns, and \(L\) is the maximum length of words. For each cell on the board, we explore up to 4 directions (up, down, left, right) and perform DFS, which has a time complexity of \(O(4^L)\).
- Total: \(O(WL + MN \times 4^L)\)

## Space Complexity

- Trie: \(O(WL)\), where \(W\) is the total number of words and \(L\) is the average length of words. The space complexity of the trie depends on the number of words and characters.
- DFS Stack: \(O(L)\), where \(L\) is the maximum length of words. The maximum depth of the DFS recursion stack is equal to the length of the longest word.
- Visited Set: \(O(M \times N)\), where \(M\) is the number of rows and \(N\) is the number of columns. We use a visited set to track visited cells during DFS traversal.
- Total: \(O(WL + MN)\)


