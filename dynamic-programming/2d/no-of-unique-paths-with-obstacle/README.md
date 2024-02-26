# Number of unique paths with an obstacle

The problem involves a robot located at the top-left corner of an `m x n` grid, aiming to reach the bottom-right corner. The robot can only move either down or right at any given point. However, the grid contains obstacles marked as `1`, while empty spaces are marked as `0`. The task is to determine the number of unique paths the robot can take from the top-left corner to the bottom-right corner, while avoiding obstacles.

## Intuition:
The problem can be solved by using dynamic programming, where we iteratively calculate the number of unique paths to each cell in the grid, considering the presence of obstacles. By breaking down the problem into smaller subproblems and leveraging the optimal substructure property, we can efficiently compute the total number of paths.

## Approach:
1. **Initialize DP Array:** Create a 2D array `dp` to store the number of unique paths to each cell.
2. **Handle Base Cases:** If there's an obstacle at the starting point or the destination, return `0`. Otherwise, initialize the top-left corner of the `dp` array to `1`.
3. **Iterate Through Grid:** Traverse the grid, breaking out of the loop when encountering obstacles in the first row or column. For each cell:
    - If it's an obstacle, set the corresponding `dp` value to `0`.
    - Otherwise, update `dp[i][j]` based on the number of paths from the cells above and to the left, considering obstacles.
4. **Return Result:** Finally, return the value stored in `dp[m-1][n-1]`, representing the number of unique paths to the bottom-right corner.

## Time Complexity:
The time complexity of the solution is O(m * n), where `m` and `n` are the dimensions of the grid. This complexity arises from iterating through each cell of the grid once to calculate the number of unique paths.

## Space Complexity:
The space complexity is O(m * n) as well. It's due to the usage of a 2D array `dp` with dimensions corresponding to the grid's dimensions, storing intermediate results for dynamic programming.