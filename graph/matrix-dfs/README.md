# No of unique paths with obstacles using Matrix DFS

You are given a grid representing a maze, where 0 represents an empty cell and 1 represents an obstacle. The task is to find the number of paths from the top-left corner to the bottom-right corner of the maze, avoiding obstacles.

## Intuition
We can solve this problem using depth-first search (DFS), backtracking through the maze from the starting point to the ending point while keeping track of visited cells. At each step, we explore all possible directions (up, down, left, right) from the current cell and continue until we reach the destination cell.

## Approach
1. Define a DFS function that recursively explores all possible paths from a given cell.
2. Start DFS from the top-left corner of the maze.
3. At each step of DFS:
   - Check if the current cell is out of bounds, an obstacle, or already visited. If any of these conditions are met, return 0.
   - If the current cell is the bottom-right corner of the maze, return 1.
   - Otherwise, mark the current cell as visited, and recursively explore all four directions (up, down, left, right).
4. Backtrack by unmarking the current cell as visited before returning from the DFS function.
5. Count and return the total number of paths found.

## Time Complexity
The time complexity of this approach is O(4^nm), where n is the number of rows and m is the number of columns in the maze. This is because, in the worst case, there are four possible directions to explore from each cell, and the DFS function may be called recursively for each cell in the maze.

## Space Complexity
The space complexity of this approach is O(nm) due to the space used by the visited map to track visited cells in the maze. Additionally, recursive function calls consume stack space proportional to the depth of the recursion, which is bounded by the number of cells in the maze.