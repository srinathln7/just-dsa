# Max Area of an Island
Given a 2D grid `grid` containing `0`s (water) and `1`s (land), find the maximum area of an island. An island is a group of connected land cells (horizontal or vertical). You may assume all four edges of the grid are surrounded by water.

## Intuition
To find the maximum area of an island, we iterate through each cell in the grid. Whenever we encounter a land cell marked by `1`, we perform a BFS on that cell to calculate the area of the island connected to it. We maintain the maximum area encountered so far.

## Approach
1. Define a function `maxAreaOfIsland(grid [][]int) int` to find the maximum area of an island.
2. Initialize `maxArea` to 0.
3. Iterate through each cell in the grid:
   - If the cell is land (`grid[r][c] == 1`), initialize a counter `count` to 0.
   - Perform a BFS starting from the current cell to calculate the area of the island connected to it. Update `count` with the area.
   - Update `maxArea` with the maximum of `maxArea` and `count`.
4. Return `maxArea`.

## Time Complexity
The time complexity of this solution is O(m * n), where m is the number of rows and n is the number of columns in the grid. We visit each cell once.

## Space Complexity
The space complexity is O(m * n) in the worst case, where m is the number of rows and n is the number of columns. This is due to the space used by the BFS queue.
