# Best time to buy or sell with cooldown period
You are given an array prices where prices[i] is the price of a given stock on the ith day. Find the maximum profit you can achieve. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times) with the following restrictions:
- After you sell your stock, you cannot buy stock on the next day (i.e., cooldown one day).
- Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).

## Intuition
To maximize the profit with the given constraints, we can use dynamic programming with memoization. We define a state consisting of the current index and whether the stock is bought or not. We recursively explore the possibilities of buying, selling, or cooling down at each index while memoizing the maximum profit obtained.

## Approach
1. Define a state struct to represent the current index and whether the stock is bought or not.
2. Use memoization to store the maximum profit for each state encountered during the recursion to avoid redundant computations.
3. Implement a recursive function `dfs` to calculate the maximum profit recursively for each state.
4. In the `dfs` function:
   - If the maximum profit for the current state is already computed and stored in the memoization map, return the computed value.
   - If the current index is beyond the length of the prices array, return 0.
   - Recursively explore the possibilities of buying, selling, or cooling down at each index and update the maximum profit accordingly.
5. Return the maximum profit obtained.

## Time Complexity
The time complexity of this approach is O(n), where n is the length of the prices array. This is because each state is visited and calculated only once, and the memoization technique helps avoid redundant computations.

## Space Complexity
The space complexity of the provided solution is also O(n), where n is the length of the prices array. This is because we are using memoization to store the maximum profit for each state encountered during the recursion. Since there are potentially n * 2 different states (each index with two possible buy/sell states), the space required to store the memoization table is proportional to the length of the prices array. Therefore, the overall space complexity is O(n).
