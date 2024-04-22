# Buy-Sell Stock 4 

You are given an integer array `prices` where `prices[i]` is the price of a given stock on the ith day, and an integer `k`. Find the maximum profit you can achieve. You may complete at most `k` transactions: i.e., you may buy at most `k` times and sell at most `k` times.

Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).

**Intuition:**

To solve this problem, we can use dynamic programming (DP) with memoization. We need to consider all possible states to determine the maximum profit. We can define states based on the current day, the number of transactions completed, and whether we currently hold a stock or not. By considering all possible states and transitions, we can calculate the maximum profit.

**Approach:**

1. Initialize a two-dimensional array `dp` of size `(k + 1) x n`, where `k` is the maximum number of transactions allowed, and `n` is the length of the `prices` array. `dp[i][j]` represents the maximum profit achievable up to day `j` with at most `i` transactions.
2. Initialize the base cases:
   - `dp[0][j]` (no transactions allowed) is always 0.
   - `dp[i][0]` (on the first day) is always 0.
3. Iterate over `i` from `1` to `k` and `j` from `1` to `n - 1`.
4. For each `(i, j)` pair, calculate the maximum profit considering two scenarios:
   - We don't make any transaction on day `j`, so the maximum profit is `dp[i][j-1]`.
   - We make a transaction on day `j`, so we need to find the maximum profit we can achieve by buying the stock on any day `m` from `0` to `j-1`, and then selling it on day `j`. This can be represented as `prices[j] - prices[m] + dp[i-1][m]` for all `m` from `0` to `j-1`.
5. Update `dp[i][j]` as the maximum profit between the two scenarios calculated in step 4.
6. Return `dp[k][n-1]`, which represents the maximum profit achievable with at most `k` transactions by the end of the last day.

**Time Complexity:**

The time complexity of this approach is O(k * n), where `n` is the length of the `prices` array and `k` is the maximum number of transactions allowed.

**Space Complexity:**

The space complexity is O(k * n), where `n` is the length of the `prices` array and `k` is the maximum number of transactions allowed. This space is used to store the `dp` array.
