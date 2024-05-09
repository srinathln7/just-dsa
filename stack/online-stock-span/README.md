# Online Stock Span

Design a class to calculate the span of stock prices. The span of the stock's price today is defined as the maximum number of consecutive days (starting from today and going backward) for which the stock price was less than or equal to today's price.

## Intuition

We can solve this problem efficiently using a stack. Whenever we encounter a new stock price, we check the prices in the stack. If the new price is greater than the prices in the stack, we pop those prices and accumulate their spans. Then, we append the new price to the stack along with its span.

## Approach

1. Define a `Price` struct to store the price value and its corresponding span.
2. Implement a `StockSpanner` struct with a stack to store prices.
3. Implement a constructor function `Constructor` that initializes a new `StockSpanner` instance.
4. Implement a `Next` method for the `StockSpanner` struct to calculate the span for the next stock price.
5. In the `Next` method:
   - Initialize the span to 1.
   - While the stack is not empty and the current price is greater than or equal to the price at the top of the stack:
     - Pop the top element from the stack.
     - Accumulate its span to the current span.
   - Append the current price along with its span to the stack.
   - Return the current span.

## Time Complexity

- The `Next` operation has an average time complexity of O(1) because each price is pushed and popped from the stack at most once.

## Space Complexity

- The space complexity is O(n), where n is the number of elements in the stack, which represents the number of stock prices encountered so far.