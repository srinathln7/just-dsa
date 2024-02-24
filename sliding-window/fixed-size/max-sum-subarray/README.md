# Kadens Algorithm

Problem Statement: The given code consists of two functions: `maxSumSubArray` and `maxSubArray`. `maxSumSubArray` calculates the maximum sum of a contiguous subarray within the given array `nums`. `maxSubArray` extends this functionality to also return the subarray portion containing the maximum sum.

Intuition:
- Both functions utilize Kadane's algorithm, which efficiently finds the maximum sum subarray in a given array.
- `maxSumSubArray` focuses solely on computing the maximum sum, while `maxSubArray` extends this to also return the subarray.

Approach:
1. **maxSumSubArray**:
   - Initialize two variables, `maxSoFar` and `sumEndingHere`, to the first element of the array.
   - Iterate through the array, updating `sumEndingHere` as the maximum of the current element and the sum of the current element and the previous sum (`nums[i]+sumEndingHere`).
   - Update `maxSoFar` as the maximum of the current `maxSoFar` and `sumEndingHere`.
   - Finally, return `maxSoFar`.

2. **maxSubArray**:
   - Initialize variables to track the start and end indices (`startIdx` and `endIdx`) of the subarray with the maximum sum.
   - Similar to `maxSumSubArray`, iterate through the array while updating `sumEndingHere`.
   - Additionally, track the start index of the current subarray (`startTracker`) to handle cases where a new subarray needs to start.
   - Update the `startIdx`, `endIdx`, and `maxSoFar` accordingly whenever a new maximum sum subarray is found.
   - Finally, return both the maximum sum (`maxSoFar`) and the subarray portion containing the maximum sum (`nums[startIdx:endIdx+1]`).

Time Complexity: Both functions iterate through the array once, resulting in a time complexity of O(n), where n is the length of the input array `nums`.

Space Complexity: Both functions use constant extra space, resulting in a space complexity of O(1).

Overall, these functions efficiently find the maximum sum subarray and return both the maximum sum and the subarray portion containing the maximum sum.
