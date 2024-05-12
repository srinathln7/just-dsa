# [Shortest Bridge](https://leetcode.com/problems/shortest-bridge/)
- Given a binary matrix where '1' represents land and '0' represents water, there are exactly two islands in the grid.
- An island is a connected group of '1's.
- The task is to find the shortest bridge to connect the two islands by changing '0's to '1's.

## Approach:
1. **DFS to Mark First Island:**
   - The first step is to find and mark all the cells of the first island.
   - This is achieved by traversing the grid using DFS starting from any cell with a value of '1'.
   - The cells of the first island are marked with a value of '2', and their positions are stored in a queue.

2. **BFS to Find Second Island:**
   - After marking the first island, perform BFS from the cells of the first island to find the shortest path to the second island.
   - Explore the neighboring cells in all four directions (up, down, left, right) from the current cell.
   - If a neighboring cell is part of the second island ('1'), return the number of steps taken to reach it.
   - If a neighboring cell is a valid land cell ('1') that has not been visited yet, mark it as visited ('2') and add it to the queue for further exploration.

3. **Return Steps or -1:**
   - If the BFS traversal successfully finds the second island, return the number of steps taken to reach it.
   - If no connection is found, return -1 to indicate that it's impossible to connect the two islands.

## Code Explanation:
- The `shortestBridge` function starts by performing DFS to mark the cells of the first island and store their positions in the `queue`.
- Then, it executes BFS from the cells in the `queue` to find the shortest path to the second island.
- During BFS, it explores neighboring cells, checks if they belong to the second island, and updates their values if necessary.
- The `dfs` function is a helper function to perform DFS traversal to mark the cells of the first island.

## Time Complexity:
- The time complexity of this solution is O(n^2), where n is the size of the grid. This is because both DFS and BFS visit each cell of the grid once.

## Space Complexity:
- The space complexity is also O(n^2) due to the queue used to store cell positions.

