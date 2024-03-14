# Unbounded Knapsack Problem

Given weights and profits of items, and a capacity of a knapsack, determine the maximum profit that can be obtained by selecting any number of items (including repetitions) such that the total weight does not exceed the capacity of the knapsack.

## Intuition
This problem is a variation of the classical knapsack problem where an item can be selected multiple times. We can solve it using dynamic programming by considering each item and capacity combination and selecting the maximum profit.

## Approach
1. Initialize a 1D array `dp` of size `capacity + 1` to store the maximum profit achievable for different capacities.
2. Iterate over each item in the profit array and update the `dp` array:
   - For each capacity `c` from 0 to the knapsack capacity `capacity`:
     - Calculate the maximum profit that can be achieved by including the current item, considering the weight constraints.
     - Update the `dp` array at index `c` with the maximum profit achieved.
3. Return `dp[capacity]` as the maximum profit.

## Time Complexity
The time complexity of this approach is O(n * m), where `n` is the number of items and `m` is the capacity of the knapsack.

## Space Complexity
The space complexity is O(m), where `m` is the capacity of the knapsack, as we only use a 1D array to store the maximum profit values.
