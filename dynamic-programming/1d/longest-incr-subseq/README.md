# Longest Increasing Subsequence

Given an integer array `nums`, find the length of the longest strictly increasing subsequence.

## Intution
The intuition behind finding the longest increasing subsequence (LIS) can be summarized as follows:

1. **Dynamic Programming (First Code):**
   - The idea is to build a dynamic programming (DP) array, where `dp[i]` represents the length of the LIS starting at index `i`.
   - We iterate through the array from right to left and compare each element with all elements to its right.
   - If an element to the right is greater than the current element, we update `dp[i]` with the maximum length found so far.
   - Finally, we find the maximum value in the `dp` array, which represents the length of the longest increasing subsequence.

2. **Recursive Approach with Memoization (Second Code):**
   - In this approach, we use recursion with memoization (top-down dynamic programming).
   - We initialize a memoization table `dp` to store the length of the LIS starting from each index.
   - We implement a recursive function to find the length of the LIS. For each index, we iterate through the array to its right.
   - If an element to the right is greater than the current element, we recursively find the length of the LIS starting from that index.
   - We store the result in the `dp` table to avoid redundant calculations.
   - Finally, we start the recursion from each index and return the maximum length found.

The intuition behind both approaches is to identify the optimal substructure of the problem and use dynamic programming techniques to efficiently solve it. By breaking down the problem into smaller subproblems and memoizing the results, we can avoid redundant calculations and optimize the overall runtime complexity.

## Approach
### First Code
- Initialize a dynamic programming (DP) array `dp`, where `dp[i]` represents the length of the LIS starting at index `i`.
- Iterate through the array from right to left.
  - For each element at index `i`, compare it with all elements to its right.
  - If an element to the right is greater than the current element, update `dp[i]` with the maximum length found so far.
- Find the maximum value in the `dp` array, representing the length of the longest increasing subsequence.
- Time complexity: O(n^2), where n is the length of the input array.

### Second Code
- Use a recursive approach with memoization (top-down dynamic programming).
- Initialize a memoization table `dp` to store the length of the LIS starting from each index.
- Implement a recursive function to find the length of the LIS.
  - For each index, iterate through the array to its right.
  - If an element to the right is greater than the current element, recursively find the length of the LIS starting from that index.
  - Store the result in the `dp` table.
- Start recursion from each index and return the maximum length found.
 
## Time complexity: 
O(n^2), where n is the length of the input array.

## Spacec complexity
O(n)
