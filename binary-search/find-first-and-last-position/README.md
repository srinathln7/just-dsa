# Find first and last position
Given an array of integers `nums` sorted in non-decreasing order, find the starting and ending position of a given `target` value. If `target` is not found in the array, return `[-1, -1]`.

You must write an algorithm with O(log n) runtime complexity.

## Intuition
To find the starting and ending positions of the target value in the array, we can use a modified binary search algorithm. We need to run the binary search algorithm twice: once to find the leftmost occurrence of the target and once to find the rightmost occurrence of the target.

## Approach
1. Implement a binary search function that takes the array `nums`, the `target` value, and a boolean flag `isLeft` indicating whether to search for the leftmost or rightmost occurrence of the target.
2. Initialize two pointers `l` and `r` to the start and end of the array, respectively.
3. Use a loop to perform binary search until `l` is less than or equal to `r`.
4. Calculate the middle index `mid` using the formula `mid := l + (r-l)/2`.
5. If the target value is found at index `mid`:
   - Update the result `res` to the current index `mid`.
   - If `isLeft` is `true`, update `r` to `mid - 1` to search for the leftmost occurrence.
   - If `isLeft` is `false`, update `l` to `mid + 1` to search for the rightmost occurrence.
6. If the target value is less than the value at index `mid`, update `r` to `mid - 1`.
7. If the target value is greater than the value at index `mid`, update `l` to `mid + 1`.
8. Return the result `res`.

## Time Complexity
The time complexity of this approach is O(log n), where n is the number of elements in the array `nums`.

## Space Complexity
The space complexity of this approach is O(1) as it uses only a constant amount of extra space.