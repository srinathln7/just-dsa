# Number of unique paths in a m*n grid from top-left corner to bottom-right corner

## Intuition:
 We use dynamic programming to compute the number of unique paths from the top-left corner to the bottom-right corner of the grid. At each cell, the number of unique paths to reach that cell is the sum of the number of paths from the cell above and the cell to the left.
  
## Approach:
We create a 2D array `dp` of size `m x n` to store the number of unique paths to each cell. We initialize the base cases where there's only one way to reach cells in the first row and first column. Then, we iterate over each cell, filling in the `dp` array based on the recurrence relation mentioned earlier.

## Time Complexity:
The time complexity is O(m*n) because we iterate over each cell in the grid once to compute the number of unique paths.

## Space Complexity: 
The space complexity is also O(m*n) because we use a 2D array of size `m x n` to store the intermediate results.