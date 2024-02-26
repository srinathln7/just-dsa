**Problem Statement:**
Given an array `prices` representing the price of a stock on each day, we aim to maximize profit by buying and selling the stock multiple times. However, we can only hold one share of the stock at a time. 

**Intuition:**
The key idea is to sell the stock whenever there is a profit available. By doing so, we can capture all the possible profit from each upward movement in stock prices.

**Approach:**
1. Initialize `maxProfit` and `profit` variables to 0.
2. Iterate through the prices array starting from the second day.
3. For each day, calculate the profit by subtracting the price of the previous day from the price of the current day.
4. If the profit is positive (indicating an increase in stock price), add it to the `maxProfit`.
5. Continue this process for all days.
6. Return `maxProfit` as the final result.

**Time Complexity:**
The time complexity of the solution is O(n), where n is the number of elements in the prices array. This is because we iterate through the prices array only once.

**Space Complexity:**
The space complexity is O(1) since we use a constant amount of extra space regardless of the input size. We only use a few variables to store intermediate results.

Overall, the provided code efficiently solves the problem with linear time complexity and constant space complexity, allowing us to maximize profit by buying and selling stocks multiple times.