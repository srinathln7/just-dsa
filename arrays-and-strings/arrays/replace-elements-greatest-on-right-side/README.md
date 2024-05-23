# Replace Elements with Greatest Element on Right Side 
Given an array of integers `arr`, replace each element with the greatest element among the elements to its right, and replace the last element with `-1`.

## Intuition
To solve this problem efficiently, we can traverse the array from right to left. By keeping track of the maximum element seen so far to the right, we can replace each element with this maximum. This approach ensures that we only pass through the array once, maintaining a time complexity of O(n).

## Approach
1. Initialize `rightMax` to `-1` since the last element should be replaced with `-1`.
2. Iterate through the array from the end to the beginning.
3. For each element, store the current `rightMax` in a temporary variable.
4. Update `rightMax` to be the maximum of the current element and `rightMax`.
5. Replace the current element with the temporary variable (the previous `rightMax`).
6. Continue until all elements have been processed.

## Time Complexity
- **O(n)**: We traverse the array once, where `n` is the length of the array.

## Space Complexity
- **O(1)**: We use a constant amount of extra space for the variables.
