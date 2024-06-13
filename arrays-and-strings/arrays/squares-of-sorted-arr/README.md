# [Squares of Sorted Array](https://leetcode.com/problems/squares-of-a-sorted-array/description/)
Given a sorted array of integers, return a new array containing the squares of each number, also sorted in non-decreasing order.

## Approach in `Alt.go`
1. **Find Index of Last Negative Number**: Iterate through the array until we find the last negative number (or zero), if any. This index marks the boundary between negative and non-negative numbers in the sorted array.
2. **Square and Reverse Negative Numbers**: Square the numbers in the range [0, idx] (inclusive) and reverse the array. This ensures the squared negative numbers are in non-decreasing order.
3. **Square Positive Numbers**: Square the numbers in the range [idx+1, n-1] (inclusive).
4. **Merge Two Sorted Arrays**: Merge the two sorted arrays (squared negatives and squared positives) into one sorted array using the two-pointer approach.

## Time Complexity
- **O(n)**: We iterate through the input array once to find the last negative number and then again to square the numbers. Merging the sorted arrays also takes linear time.

## Space Complexity
- **O(n)**: We create three new arrays: one for squared negatives, one for squared positives, and one for the merged result.
