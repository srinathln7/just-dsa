# Coin change - Unbounded Knapsack

Given an integer array `coins` representing coins of different denominations and an integer `amount` representing a total amount of money, find the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

## Intuition
The problem can be solved using dynamic programming. We can use a bottom-up approach where we iteratively calculate the fewest number of coins needed to make up each amount from 0 to the target amount.

## Approach
1. Initialize an array `dp` of size `amount + 1` to store the fewest number of coins needed for each amount from 0 to the target amount.
2. Initialize all values in the `dp` array to `amount + 1`, except for `dp[0]` which is set to 0.
3. Iterate through each amount from 1 to the target amount:
   - For each amount, iterate through each coin denomination.
   - If the current coin denomination is less than or equal to the current amount, update `dp[targetAmt]` to be the minimum of its current value and `1 + dp[targetAmt - coin]`.
4. After iterating through all amounts and coin denominations, if `dp[amount]` remains unchanged (i.e., `amount + 1`), return -1 as it indicates that the target amount cannot be made up by any combination of the coins.
5. Otherwise, return `dp[amount]`, which represents the fewest number of coins needed for the target amount.

## Time Complexity
- Time complexity: O(n * amount), where n is the number of coins and amount is the target amount. We iterate through each coin denomination for each amount from 1 to the target amount.

## Space Complexity
- Space complexity: O(amount), as we use an array of size `amount + 1` to store the fewest number of coins needed for each amount.
