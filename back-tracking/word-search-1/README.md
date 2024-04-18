#  Word Search

Given a 2D board of characters and a word, determine if the word exists in the board. The word can be constructed from letters of sequentially adjacent cells, where "adjacent" cells are horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.

## Intuition

To solve this problem, we can use a depth-first search (DFS) algorithm with backtracking. We start by iterating through each cell of the board. For each cell, if the character matches the first character of the word, we initiate a DFS search to explore all possible paths from that cell.

## Approach

1. Iterate through each cell of the board.
2. For each cell, if the character matches the first character of the word, initiate a DFS search.
3. In the DFS function, check if the current index of the word matches the character in the current cell. If it does, explore the neighboring cells recursively.
4. If a complete word is found, return true.
5. Otherwise, backtrack by restoring the character of the current cell and continue exploring other paths.
6. Repeat steps 3-5 until all paths are exhausted.
7. If no path results in finding the word, return false.

## Time Complexity

Let \( m \) be the number of rows and \( n \) be the number of columns in the board, and \( k \) be the length of the word.

The time complexity of this approach is \( O(m \times n \times 4^k) \) due to the DFS search, where \( 4^k \) is the number of possible directions to explore at each step.

## Space Complexity

The space complexity of this approach is \( O(k) \), where \( k \) is the length of the word. This space is used for the recursive call stack during the DFS traversal.