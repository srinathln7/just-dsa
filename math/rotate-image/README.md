# Rotate Image
You are given an n x n 2D matrix representing an image. Rotate the image by 90 degrees (clockwise).
You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

## Intuition
To rotate the image by 90 degrees clockwise in-place, we can perform the following steps:
1. Transpose the matrix by swapping matrix[i][j] with matrix[j][i] for all valid i and j.
2. Reverse each row of the transposed matrix to complete the rotation.

## Approach
1. Transpose the matrix by iterating over the upper triangular portion of the matrix and swapping elements.
2. Reverse each row of the transposed matrix to complete the rotation.

## Time Complexity
The time complexity of this approach is O(n^2), where n is the number of rows (or columns) in the matrix.

## Space Complexity
The space complexity is O(1) since we perform the rotation in-place without using any extra space.
