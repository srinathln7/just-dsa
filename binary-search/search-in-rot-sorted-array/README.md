# Search in Rotated sorted array

Given an array `nums` of distinct integers sorted in ascending order, and a target value, return the index of the target value in the array. If the target is not found, return -1.

## Intuition

Since the array is sorted in ascending order, we can apply binary search to efficiently find the target element. However, the array might be rotated at an unknown pivot index. To handle this rotation, we can still use binary search but need to modify the search conditions based on whether the left or right half of the array is sorted.

## Approach

1. Initialize two pointers `l` and `r` representing the start and end of the search range, respectively.
2. Use a loop to iteratively narrow down the search range until `l` is less than or equal to `r`.
3. Inside the loop:
   - Calculate the middle index `mid` using the formula `mid := l + (r-l)/2`.
   - Check if the middle element `nums[mid]` is equal to the target. If it is, return the index `mid`.
   - Determine whether the right half or left half of the array is sorted using a `switch` statement.
   - Adjust the pointers `l` and `r` based on whether the right half or left half is sorted to continue the search.
4. If the target is not found after the loop ends, return -1.

## Time Complexity

The time complexity of the binary search algorithm is O(log n), where n is the number of elements in the array. This is because the search space is halved in each iteration of the loop, leading to a logarithmic time complexity.

## Space Complexity
O(1)
