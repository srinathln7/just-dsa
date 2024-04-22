# Pacific Atlantic Waterflow 
Given an m x n matrix representing the heights of cells, determine the cells from which rainwater can flow to both the Pacific and Atlantic oceans.

## Input
- An m x n integer matrix `heights` representing the heights of cells on the island.

## Output
- A 2D list of grid coordinates `result`, where each element `result[i] = [ri, ci]` denotes that rainwater can flow from cell `(ri, ci)` to both the Pacific and Atlantic oceans.

## Intuition
To solve this problem, we can perform a depth-first search (DFS) from the boundaries of the matrix (Pacific and Atlantic oceans) to identify all cells reachable from each ocean. Then, we can find the cells reachable from both oceans and return their coordinates.

## Approach
1. Create two boolean matrices `pacificReachable` and `atlanticReachable` of size m x n to store whether each cell is reachable from the Pacific and Atlantic oceans, respectively.
2. Perform DFS from all boundary cells along the left and top edges to mark cells reachable from the Pacific Ocean in `pacificReachable`.
3. Perform DFS from all boundary cells along the right and bottom edges to mark cells reachable from the Atlantic Ocean in `atlanticReachable`.
4. Iterate through all cells in the matrix:
   - If a cell is reachable from both oceans (i.e., `pacificReachable` and `atlanticReachable` are both true), add its coordinates to the result.
5. Return the result.

## Time Complexity
The time complexity is O(m * n), where m is the number of rows and n is the number of columns in the matrix. We perform DFS once for each ocean, which takes O(m * n) time, and then iterate through the entire matrix once.

## Space Complexity
The space complexity is O(m * n) for storing the boolean matrices `pacificReachable` and `atlanticReachable`.
