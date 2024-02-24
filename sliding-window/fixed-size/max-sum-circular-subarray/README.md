# Max-sum circular subarray

Given a circular array `nums` of integers, the task is to find the maximum possible sum of a non-empty subarray of the array.

## Intuition
To find the maximum subarray sum in a circular array, we can consider two cases:
1. The maximum subarray sum is within the array (non-circular).
2. The maximum subarray sum wraps around the circular boundary.

For the first case, we can utilize Kadane's algorithm to find the maximum subarray sum in the non-circular array.

For the second case, we can find the minimum subarray sum using a modified Kadane's algorithm and subtract it from the total sum of the array. The result would be the maximum subarray sum wrapping around the circular boundary.

Finally, we return the maximum of these two sums as the maximum subarray sum for the circular array.

## Approach
1. Define two helper functions `getMaxSumSubarray` and `getMinSumSubarray` to find the maximum and minimum subarray sums using Kadane's algorithm.
2. Initialize variables `minSumForward`, `maxSumForward`, and `maxSumCircular` to store the minimum sum, maximum sum, and maximum circular sum, respectively.
3. Calculate `maxSumForward` using `getMaxSumSubarray` to find the maximum subarray sum within the array.
4. Calculate `minSumForward` using `getMinSumSubarray` to find the minimum subarray sum within the array.
5. Calculate `sumTotal` as the sum of all elements in the array.
6. Calculate `maxSumCircular` as `sumTotal - minSumForward` to find the maximum subarray sum wrapping around the circular boundary.
7. Determine `result` as the maximum of `maxSumForward` and `maxSumCircular`.
8. Return `result` as the maximum subarray sum for the circular array.

## Time Complexity
The time complexity of this algorithm is O(n), where n is the length of the input array `nums`. This is because we traverse the array only once to calculate the maximum and minimum subarray sums.

## Space Complexity
The space complexity of this algorithm is O(1) since it uses a constant amount of extra space regardless of the input size.

## Remark:
Account for the case with only negative elements. Watch out for the edge case at the last.