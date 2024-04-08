# Spiral Matrix
Given an m x n matrix, return all elements of the matrix in spiral order.

## Intuition
To traverse the matrix in spiral order, we can simulate the movement of a spiral by keeping track of four boundaries: top, bottom, left, and right. We start by moving from left to right along the top row, then from top to bottom along the rightmost column, then from right to left along the bottom row, and finally from bottom to top along the leftmost column. We repeat this process until we have visited all elements in the matrix.

## Approach
1. Initialize four variables `top`, `bottom`, `left`, and `right` to represent the boundaries of the matrix.
2. Initialize an empty array `result` to store the elements visited in spiral order.
3. Repeat the following steps while `top` is less than or equal to `bottom` and `left` is less than or equal to `right`:
    - Traverse from left to right along the top row and add each element to `result`. Increment `top`.
    - Traverse from top to bottom along the rightmost column and add each element to `result`. Decrement `right`.
    - Traverse from right to left along the bottom row and add each element to `result`. Decrement `bottom`.
    - Traverse from bottom to top along the leftmost column and add each element to `result`. Increment `left`.
4. Return `result`.

## Time Complexity
The time complexity of this approach is O(m * n), where m is the number of rows and n is the number of columns in the matrix.

## Space Complexity
The space complexity is O(1) since we use only a constant amount of extra space for storing variables and the result array.
