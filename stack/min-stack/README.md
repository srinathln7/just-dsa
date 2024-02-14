# Min Stack

## Problem

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

- `push(x)` -- Push element `x` onto stack.
- `pop()` -- Removes the element on top of the stack.
- `top()` -- Get the top element.
- `getMin()` -- Retrieve the minimum element in the stack.

## Approach:
The `MinStack` struct contains a slice of `Element`, where each `Element` consists of two fields: `val` representing the value of the element and `min` representing the minimum value in the stack up to that element.

- During `Push`, if the stack is empty, the minimum value is set to the current value being pushed. Otherwise, it compares the current value with the minimum value of the top element in the stack and sets the minimum value accordingly.
- `Pop` simply removes the top element from the stack.
- `Top` returns the value of the top element in the stack.
- `GetMin` returns the minimum value stored in the stack.

## Time Complexity:
- Push: O(1)
- Pop: O(1)
- Top: O(1)
- GetMin: O(1)

## Space Complexity:
O(n)

