# Triangle Min Path Sum 
Given a triangle array, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.

## Intuition
This problem can be solved using dynamic programming, specifically memoization. We can recursively traverse the triangle from top to bottom, memoizing the minimum path sum for each position. 

## Approach
1. Start from the top of the triangle and define a recursive function to calculate the minimum path sum.
2. Use memoization to store the minimum path sum for each position.
3. In the recursive function, calculate the minimum path sum for the current position by considering the two possible paths to the next row: one to the left and one to the right.
4. Return the minimum path sum for the current position.
5. Use the memoization table to avoid redundant calculations.

## Time Complexity
The time complexity of this approach is O(n^2), where n is the number of rows in the triangle. This is because we traverse each position in the triangle only once and perform constant-time operations for each position.

## Space Complexity
The space complexity of this approach is also O(n^2), where n is the number of rows in the triangle. This is because we use a memoization table to store the minimum path sum for each position in the triangle.