# 2D-array search

Given a 2D matrix where each row is sorted in ascending order and each column is sorted in ascending order, determine if a target integer exists in the matrix.

## Intuition:
Since the matrix is sorted both row-wise and column-wise, we can exploit this property to perform an efficient search. We can use a binary search algorithm to search for the target in the matrix.

## Approach:
1. Initialize pointers `l` and `r` to the first and last indices of the **flattened** 2D matrix, respectively.
2. While `l` is less than or equal to `r`:
    - Calculate the middle index `midIndex` as `l + (r-l)/2`.
    - Convert `midIndex` back to 2D indices to get the value at the middle of the matrix.
    - Compare the middle value with the target:
        - If they are equal, return true.
        - If the target is less than the middle value, update `r` to `midIndex - 1`.
        - If the target is greater than the middle value, update `l` to `midIndex + 1`.
3. If the loop exits without finding the target, return false.

An alternative approach is to first perform binary search on the row `O(log m)` and then search if a particular element is present in that row `O(log n)`. 

## Time Complexity:
The time complexity of this algorithm is O(log(mn)), where m is the number of rows and n is the number of columns in the matrix. This is because we perform a binary search on a 1D array of size m * n. 

## Space Complexity:
The space complexity of this algorithm is O(1) because we only use a constant amount of extra space for variables regardless of the size of the input matrix.