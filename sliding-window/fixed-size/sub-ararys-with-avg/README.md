# Number of Sub-arrays of Size K and Average Greater than or Equal to Threshold

Given an integer array `arr`, an integer `k`, and a threshold `threshold`, the task is to count the number of subarrays of length `k` in `arr` such that the average of elements in each subarray is greater than or equal to `threshold`.

## Intuition:
The solution involves using a sliding window approach to efficiently calculate the sum of elements within each subarray window and then determine if the average meets the threshold condition.

## Approach:
1. Initialize variables `count`, `avg`, and `windowSum`.
2. Iterate through the first `k` elements of the array to calculate the initial `windowSum`.
3. Iterate through the array starting from the `(k+1)th` element:
   - Calculate the average (`avg`) using the current `windowSum` divided by `k`.
   - If `avg` is greater than or equal to `threshold`, increment the `count`.
   - Slide the window by adding the current element and subtracting the leftmost element from `windowSum`.
4. After the loop exits, check the average of the last window and increment the `count` if it satisfies the threshold condition.
5. Return the final `count`.

## Time Complexity:
The time complexity of this solution is O(n), where n is the length of the input array `arr`. This is because we iterate through the array only once.

## Space Complexity:
The space complexity is O(1) since we are using only a constant amount of extra space for storing variables regardless of the input size.

