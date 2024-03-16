# Min.cost stair-case problem

You are given an integer array `cost` where `cost[i]` is the cost of the ith step on a staircase. Once you pay the cost, you can either climb one or two steps.

You can either start from the step with index 0 or the step with index 1.

Return the minimum cost to reach the top of the floor.

## Intuition

This problem can be solved using dynamic programming. We can define a `dp` array where `dp[i]` represents the minimum cost to reach the ith step. The minimum cost to reach the current step `i` can be calculated by taking the minimum of the costs of reaching the previous two steps (`i-1` and `i-2`) and adding the cost of the current step `cost[i]`. We can fill up the `dp` array iteratively, starting from the third step (since the minimum cost to reach the first and second steps is given). Finally, we return the minimum of the costs of reaching the last two steps, which represents the minimum cost to reach the top of the floor.

## Approach
1. Initialize variables `prev1` and `prev2` to store the minimum costs for the previous two steps.
2. Iterate through the steps starting from the third step:
   - Calculate the minimum cost to reach the current step by adding the cost of the current step with the minimum of the previous two steps.
   - Update `prev1` and `prev2` with the new minimum costs for the previous two steps.
3. Return the minimum cost of reaching the last two steps, which represents the minimum cost to reach the top of the floor.


## Time Complexity

The time complexity of this approach is O(n), where `n` is the length of the `cost` array, since we iterate through the array once to fill up the `dp` array.

## Space Complexity

The space complexity is O(1) 





