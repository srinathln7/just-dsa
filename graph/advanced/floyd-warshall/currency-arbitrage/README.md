# Currency Arbitrage
The code is designed to find arbitrage opportunities in currency exchange rates. Arbitrage is the practice of taking advantage of a price difference between two or more markets, exploiting the difference to make a profit.

## Intuition:
The code checks all possible combinations of three different currencies (i, j, k) and calculates the product of the exchange rates between these currencies. If the product is greater than 1, it indicates an arbitrage opportunity.

## Approach:
1. Define a constant `exchange` which stores exchange rates between different currencies in a nested map.
2. Iterate over all combinations of three different currencies (i, j, k) using nested loops.
3. For each combination, calculate the product of exchange rates between these currencies.
4. If the product is greater than 1, print the currencies involved in the arbitrage opportunity along with the calculated product.

## Time Complexity:
- The time complexity of this algorithm is O(n^3), where n is the number of currencies. This is because we iterate over all possible combinations of three different currencies.

## Space Complexity:
- The space complexity is O(1) since we only use a constant amount of extra space for storing temporary variables and currency symbols.