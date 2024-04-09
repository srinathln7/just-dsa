# Exponentiation
Implement the `myPow` function, which calculates `x` raised to the power `n`. 

## Approach
The code implements binary exponentiation to efficiently compute `x` raised to the power `n` in logarithmic time complexity. Here's the breakdown:
1. If `x` is 0.0, return 0.0.
2. If `n` is 0 or `x` is 1.0, return 1.0.
3. If `n` is negative, update `x` to its reciprocal (`1 / x`) and make `n` positive.
4. Initialize a variable `result` to 1.0.
5. Iterate through `n` until it becomes 0:
   - If the least significant bit of `n` is 1, multiply `result` by `x`.
   - Square `x` in each iteration.
   - Right-shift `n` by 1 to divide it by 2.
6. Return the final value of `result`, which represents `x` raised to the power `n`.

## Time Complexity
The time complexity of this algorithm is O(log n), as it effectively reduces `n` by half in each iteration.

## Space Complexity
The space complexity is O(1) as it uses only a constant amount of additional space.
