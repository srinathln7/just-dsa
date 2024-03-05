# No. of Islands

Given an `m x n` 2D binary grid `grid` which represents a map of `'1's` (land) and `'0's` (water), return the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

## Intuition
To solve this problem, we can use either Breadth-First Search (BFS) or Depth-First Search (DFS) to traverse the grid. The idea is to traverse the grid cell by cell. Whenever we encounter a land cell marked by `'1'`, we increment the island count and perform a BFS or DFS starting from that cell to mark all connected land cells as visited.

## Approach
1. **BFS Approach**:
   - We iterate over each cell in the grid.
   - When we encounter a land cell `'1'`, we increment the island count and perform a BFS starting from that cell to mark all connected land cells as visited.
   - We use a queue to perform BFS. We enqueue the current cell and iteratively dequeue each cell, marking its adjacent land cells as visited.
   - We continue this process until the queue is empty.
   - The final count of islands is returned.

2. **DFS Approach**:
   - Similar to the BFS approach, we iterate over each cell in the grid.
   - When we encounter a land cell `'1'`, we increment the island count and perform a DFS starting from that cell to mark all connected land cells as visited.
   - We recursively traverse each cell's neighbors, marking them as visited.
   - The DFS continues until all connected land cells are visited.
   - The final count of islands is returned.

## Time and Space Complexity
Both the BFS and DFS approaches have a time complexity of O(m * n), where m is the number of rows and n is the number of columns in the grid. This is because we traverse each cell in the grid once.

The space complexity of both approaches is O(m * n) as well, considering the space required for the queue in BFS and the recursive stack in DFS.