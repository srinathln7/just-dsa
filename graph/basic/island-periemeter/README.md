# [Island Perimeter](https://leetcode.com/problems/island-perimeter/)
You are given a 2D grid representing an island. The grid is composed of 0s and 1s, where 0 represents water and 1 represents land. Calculate the perimeter of the island, where the perimeter is defined as the length of the boundary between land and water.

## Intuition
Perform a depth-first search (DFS) traversal of the island grid. For each land cell encountered during the traversal, explore its neighboring cells. Increment the perimeter whenever you encounter a water cell from a land cell or when you reach the boundary of the grid.

## Approach
1. Define a DFS function to explore the island grid recursively.
2. Start DFS traversal from any land cell.
3. For each cell encountered during traversal:
   - If the cell is out of bounds or represents water, increment the perimeter.
   - If the cell has already been visited, return without further exploration.
   - Mark the cell as visited.
   - Explore its neighboring cells recursively.
4. Return the calculated perimeter after DFS traversal.

## Time Complexity
The time complexity of this approach is O(M * N), where M is the number of rows and N is the number of columns in the grid. This is because we traverse each cell of the grid once during DFS.

## Space Complexity
The space complexity of this approach is O(1) since we are not using any additional data structures that grow with the input size. The recursive DFS call stack may require additional space, but it is bounded by the maximum depth of recursion, which is proportional to the size of the grid.