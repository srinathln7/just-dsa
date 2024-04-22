# Max Area of an Island

You are given an `m x n` binary matrix `grid`. An island is a group of `1`'s (representing land) connected 4-directionally (horizontal or vertical). You may assume all four edges of the grid are surrounded by water. The area of an island is the number of cells with a value `1` in the island. Return the maximum area of an island in the grid. If there is no island, return `0`.

## Intuition:

To find the maximum area of an island, we can use a breadth-first search (BFS) approach. We'll traverse the grid, and whenever we encounter a cell with a value of `1`, we'll perform BFS from that cell to explore the connected island cells. During BFS, we'll mark visited cells as `0` to ensure we don't count them multiple times. We'll keep track of the area of each island during BFS and return the maximum area found.

## Approach:

1. Define a BFS function to explore an island starting from a given cell `(r, c)`. This function will return the area of the island.
2. Iterate through each cell in the grid:
   - If the cell value is `1`, perform BFS from that cell and update the maximum area found.
3. During BFS:
   - Initialize a queue and enqueue the current cell `(r, c)`.
   - Mark the current cell as visited by setting its value to `0`.
   - While the queue is not empty, dequeue a cell and increment the area.
   - Explore neighboring cells (up, down, left, right) and enqueue them if they are valid (`1`) and mark them as visited.
4. Return the maximum area found.

## Time Complexity:

Let `m` be the number of rows and `n` be the number of columns in the grid.
- The BFS function takes `O(m*n)` time to traverse the entire island.
- We perform BFS for each cell in the grid, so the overall time complexity is `O(m*n*m*n)`, which simplifies to `O(m^2 * n^2)`.

## Space Complexity:

- The space complexity is `O(m*n)` for the queue used in BFS and other variables, where `m` and `n` are the dimensions of the grid.



## Alt.go

### Intuition
To find the maximum area of an island, we iterate through each cell in the grid. Whenever we encounter a land cell marked by `1`, we perform a BFS on that cell to calculate the area of the island connected to it. We maintain the maximum area encountered so far.

### Approach
1. Define a function `maxAreaOfIsland(grid [][]int) int` to find the maximum area of an island.
2. Initialize `maxArea` to 0.
3. Iterate through each cell in the grid:
   - If the cell is land (`grid[r][c] == 1`), initialize a counter `count` to 0.
   - Perform a BFS starting from the current cell to calculate the area of the island connected to it. Update `count` with the area.
   - Update `maxArea` with the maximum of `maxArea` and `count`.
4. Return `maxArea`.

### Time Complexity
The time complexity of this solution is O(m * n), where m is the number of rows and n is the number of columns in the grid. We visit each cell once.

### Space Complexity
The space complexity is O(m * n) in the worst case, where m is the number of rows and n is the number of columns. This is due to the space used by the BFS queue.
