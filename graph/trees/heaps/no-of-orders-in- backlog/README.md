# Number of backlogs in Order Book

You are given a list of orders where each order contains the price, amount, and type. Type 0 represents a buy order, and type 1 represents a sell order. Your task is to process these orders and determine the backlog of orders that have not been executed. 

## Intuition
To process the orders efficiently, we can maintain two heaps: a max heap for buy orders and a min heap for sell orders. We iterate through each order and perform the following actions:
- For buy orders:
  - Check if there are any sell orders in the sell heap with prices lower than or equal to the current buy order's price. If so, execute the transactions accordingly.
  - If the buy order cannot be fully executed, add the remaining amount to the buy heap.
- For sell orders:
  - Check if there are any buy orders in the buy heap with prices higher than or equal to the current sell order's price. If so, execute the transactions accordingly.
  - If the sell order cannot be fully executed, add the remaining amount to the sell heap.

## Approach
1. Implement two heaps: a max heap for buy orders (`buyHeap`) and a min heap for sell orders (`sellHeap`).
2. Iterate through each order in the input list.
3. For each order:
   - If it is a buy order (type 0), process it by checking the sell heap for matching orders and executing transactions.
   - If it is a sell order (type 1), process it by checking the buy heap for matching orders and executing transactions.
4. After processing all orders, calculate the backlog by summing the remaining amounts in both heaps.
5. Return the backlog modulo 10^9 + 7.

## Time Complexity
- Pushing and popping elements from the heap takes O(log n) time.
- Processing each order takes O(log n) time in the worst case.
- Overall, the time complexity is O(n log n), where n is the number of orders.

## Space Complexity
- We use two heaps to store the orders, each with a maximum size of n.
- Additionally, we use variables to store the backlog.
- Thus, the space complexity is O(n).
