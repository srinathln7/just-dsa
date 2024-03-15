# Coin change 2

You are given an integer amount representing a total amount of money and an integer array coins representing coins of different denominations. Write a function to compute the number of combinations that make up that amount. If that amount of money cannot be made up by any combination of the coins, return 0.

## Intuition

This problem can be solved using a bottom-up dynamic programming approach. We can define a `dp` array where `dp[i]` represents the number of ways to make up the amount `i` using the given coins. We can build the `dp` array iteratively, starting from `dp[0]`, which represents the base case of having one way to make up the amount 0 (by choosing no coins). Then, for each coin denomination, we iterate through the `dp` array and update the number of ways to make up each amount using the current coin denomination.

## Approach

1. Initialize a `dp` array of size `amount + 1` to store the number of ways to make up each amount.
2. Set `dp[0] = 1` to represent the base case where there is one way to make up the amount 0 (by choosing no coins).
3. Iterate through each coin denomination in the `coins` array.
4. For each coin denomination, iterate through the `dp` array from 1 to `amount`.
5. For each amount `targetAmt`, update `dp[targetAmt]` by adding the value of `dp[targetAmt - coin]` if `coin` is less than or equal to `targetAmt`. This represents the number of ways to make up the amount `targetAmt` using the current coin denomination.
6. After updating all values in the `dp` array, return `dp[amount]`, which represents the number of ways to make up the given amount using the given coins.

## Time Complexity

The time complexity of this approach is O(amount * n), where `amount` is the target amount and `n` is the number of coins.

## Space Complexity

The space complexity is O(amount) since we are using a `dp` array of size `amount + 1` to store the number of ways for each amount.
