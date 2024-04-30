# Range Sum 2D- Immutuable
Implement a `NumMatrix` class that supports the `SumRegion` method, which is designed to sum up the elements of an array within the rectangle defined by its upper left corner `(row1, col1)` and lower right corner `(row2, col2)`.

## Intuition
The problem can be efficiently solved using prefix sums. By precomputing the cumulative sum for each cell of the matrix, we can calculate the sum of any rectangular region by taking the difference between the cumulative sums of its corners.

## Approach
1. Initialize a `NumMatrix` struct with a 2D slice `prefixSum` to store the cumulative sums.
2. In the constructor `Constructor`, iterate through each cell of the input matrix and compute its cumulative sum.
3. Define the `SumRegion` method to return the sum of elements within the specified rectangle.
4. To calculate the sum of the region `(row1, col1)` to `(row2, col2)`, use the formula:

## Time Complexity
- Constructor:
- Time complexity: O(m*n), where m and n are the dimensions of the matrix.
- `SumRegion` method:
- Time complexity: O(1)

## Space Complexity
- Space complexity: O(m*n), where m and n are the dimensions of the matrix.


