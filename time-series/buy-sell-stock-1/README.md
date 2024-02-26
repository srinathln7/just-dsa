**Problem Statement:**
Given an array `prices` representing the price of a stock on each day, we need to find the maximum profit that can be obtained by buying and selling the stock once. If no profit can be made, return 0.

**Intuition:**
To maximize profit, we aim to buy the stock at the lowest possible price and sell it at the highest possible price. We can achieve this by keeping track of the minimum price seen so far and calculating the profit for each subsequent day.

**Approach:**
1. Initialize `minPrice` to the first price in the array and `maxProfit` to 0.
2. Iterate through the prices array starting from the second element.
3. For each price:
   - If the current price is less than `minPrice`, update `minPrice` to the current price.
   - Otherwise, calculate the profit for selling the stock at the current price (current price - `minPrice`), and update `maxProfit` if this profit is greater than the current maximum profit.
4. Return `maxProfit` as the result.

**Time Complexity:**
The time complexity of the solution is O(n), where n is the number of elements in the prices array. This is because we iterate through the prices array only once.

**Space Complexity:**
The space complexity is O(1) since we use a constant amount of extra space regardless of the input size. We only use a few variables to store intermediate results.

