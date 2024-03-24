# Max. Product Sub-array
Given an integer array nums, find a subarray that has the largest product, and return the product. The test cases are generated so that the answer will fit in a 32-bit integer.

## Intuition
To find the maximum product subarray, we can use dynamic programming. However, since we are dealing with both positive and negative numbers, a simple Kadane's algorithm won't work. We need to keep track of both the maximum and minimum products encountered so far.

## Approach
1. Initialize three variables: `maxSoFar`, `minSoFar`, and `maxGlobal`, all set to the first element of the input array `nums`.
2. Iterate through the input array `nums` starting from the second element.
3. For each element `num` in the array:
   - If `num` is negative, swap `maxSoFar` and `minSoFar`. This swap ensures that we handle negative numbers correctly.
   - Update `maxSoFar` to be the maximum of the current element `num` and the product of `num` and `maxSoFar`.
   - Update `minSoFar` to be the minimum of the current element `num` and the product of `num` and `minSoFar`.
   - Update `maxGlobal` to be the maximum of `maxGlobal` and `maxSoFar`. This keeps track of the maximum product encountered so far.
4. After iterating through the entire array, return `maxGlobal`, which represents the maximum product of a contiguous subarray within the input array.

## Time Complexity
The time complexity of this approach is O(n), where n is the length of the input array `nums`. We iterate through the array once.

## Space Complexity
The space complexity is O(1) since we use only a constant amount of extra space for storing variables regardless of the input size.

