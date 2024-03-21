# Minimum in Rotated Sorted Array

Given an array `nums` of unique elements sorted in ascending order, which is rotated between 1 and n times, where `n` is the length of the array, return the minimum element of this array.

You must write an algorithm that runs in O(log n) time.

## Intuition

The key idea is to utilize binary search since the original array is sorted. We can determine the minimum element by checking if the middle element is less than the rightmost element of the array. If the middle element is less than the rightmost element, it implies that the right half of the array is already sorted, so we move our search to the left half. Otherwise, we move our search to the right half of the array.

## Approach

1. Initialize two pointers `l` and `r` to represent the start and end of the array respectively.
2. While `l` is less than `r`, perform the following steps:
   - Calculate the middle index as `mid = l + (r - l) / 2`.
   - Check if `nums[mid]` is less than `nums[r]`. If so, it means the right half of the array is sorted, so set `r = mid`.
   - Otherwise, move the search to the right half by setting `l = mid + 1`.
3. After the loop ends, `r` will be pointing to the minimum element.
4. Return `nums[r]`.

## Time Complexity

The time complexity of this algorithm is O(log n) since binary search is used to find the minimum element. In each iteration, the search space is halved, leading to logarithmic time complexity.

## Space Complexity

The space complexity is O(1) since only a constant amount of extra space is used for storing variables `l`, `r`, `mid`, and the input array `nums`.