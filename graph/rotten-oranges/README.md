# Oranges Rotting

## Problem Statement

You are given a grid of size `m x n` representing the current state of a bunch of oranges in a grid, where `grid[i][j]` represents the status of the orange at position `(i, j)`:

- `0`: The cell is empty.
- `1`: The cell contains a fresh orange.
- `2`: The cell contains a rotten orange.

Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return `-1`.

## Intuition

The problem can be solved using BFS traversal on the grid. We start by identifying the rotten oranges in the initial state and perform BFS from each rotten orange, spreading the rot to its adjacent fresh oranges. We continue BFS until there are no more fresh oranges left or until we cannot find any fresh oranges adjacent to the rotten oranges.

## Approach

1. Check if the grid is empty. If it is, return `-1`.
2. Check if the grid has only empty spaces. If it does, return `0`.
3. Initialize a count variable to track the number of minutes passed.
4. Populate the initial queue with the coordinates of all rotten oranges.
5. If there are no rotten oranges initially, return `-1`.
6. Perform BFS traversal on the grid:
   - For each level of BFS, increment the count by 1.
   - For each rotten orange, explore its adjacent cells in all four directions.
   - If a fresh orange is found adjacent to a rotten orange, mark it as rotten and add its coordinates to the queue.
7. After BFS traversal, check if there are any remaining fresh oranges in the grid. If there are, return `-1`.
8. Return the count, which represents the number of minutes elapsed.

## Time Complexity

The time complexity of BFS traversal is O(V + E), where V is the number of vertices (cells) and E is the number of edges (connections between cells). In the worst case, all cells are traversed, so the time complexity is O(m * n).

## Space Complexity

The space complexity is O(m * n) since we use a queue for BFS traversal and additional space to store the grid and other variables.
