# Add Binary Strings

Given two binary strings `a` and `b`, return their sum as a binary string.

## Intuition

To add two binary numbers, we perform digit-by-digit addition, starting from the least significant digit (LSB) to the most significant digit (MSB). We also need to consider the carry generated during addition.

## Approach

1. Initialize an empty string `result` to store the sum of the binary numbers.
2. Start iterating from the LSB of both strings `a` and `b` simultaneously.
3. At each iteration:
   - Calculate the sum of the current digits from `a` and `b`, along with the carry from the previous iteration.
   - Append the LSB of the sum to the `result` string.
   - Update the carry for the next iteration.
4. After iterating through both strings:
   - If there is any remaining carry, append it to the `result` string.
5. Reverse the `result` string to obtain the correct binary sum.
6. Return the `result` string.

## Time and Space Complexity

Let `m` be the length of string `a` and `n` be the length of string `b`.

- The time complexity of this approach is O(max(m, n)) since we iterate through both strings once.
- The space complexity is O(max(m, n)) to store the resulting string `result`.
