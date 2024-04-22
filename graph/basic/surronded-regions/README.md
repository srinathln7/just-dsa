# Surronded Region

Given an m x n matrix `board` containing 'X' and 'O', capture all regions that are 4-directionally surrounded by 'X'.

A region is captured by flipping all 'O's into 'X's in that surrounded region.

## Intuition

To solve this problem, we can perform a depth-first search (DFS) starting from the boundary cells that contain 'O'. During the DFS traversal, we mark all connected 'O' cells as visited. After completing the DFS traversal, any remaining 'O' cells are those not connected to the boundary and are therefore surrounded by 'X'. We can then update these cells to 'X' to capture the surrounded regions.

## Approach

1. Perform a depth-first search (DFS) starting from the boundary cells (first and last rows, first and last columns) that contain 'O'.
2. During the DFS traversal, mark all connected 'O' cells as visited by changing them to a temporary marker, e.g., 'T'.
3. After completing the DFS traversal, any remaining 'O' cells are those not connected to the boundary and are surrounded by 'X'. Update these cells to 'X'.
4. Restore boundary-marked cells (temporary marker 'T') back to 'O'.

## Time Complexity

The time complexity of this approach is O(m * n), where m is the number of rows and n is the number of columns in the matrix.

## Space Complexity

The space complexity is O(1) since we are using only a constant amount of extra space for storing directions and variables.

## Remark

In the provided code, the cells that are visited during the depth-first search (DFS) traversal are marked with the temporary marker 'T' (`board[row][col] = 'T'`) instead of 'X' (`board[row][col] = 'X'`). 

The reason for using 'T' instead of 'X' as the marker is to differentiate between the cells that are part of the surrounded regions and those that are genuinely occupied by 'X'. If we mark the visited cells with 'X', we might inadvertently modify cells that were initially 'X' and mistakenly change them to 'O', as part of the surrounded region. This could lead to incorrect results.

By using a temporary marker like 'T', we can ensure that the original 'X' cells remain unchanged, and we only modify the 'O' cells that are part of the surrounded regions to 'X' after the DFS traversal. This helps in accurately capturing the surrounded regions without affecting the genuinely occupied 'X' cells.