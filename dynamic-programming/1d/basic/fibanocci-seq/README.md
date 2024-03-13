# Fibonacci Number

## Problem Statement

The Fibonacci numbers, commonly denoted `F(n)`, form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1. Given `n`, calculate `F(n)`.

## Intuition

The Fibonacci sequence can be generated iteratively using a constant amount of space. The key idea is to maintain two variables representing the current Fibonacci number and the next Fibonacci number. By iterating through the sequence, each time updating these two variables according to the formula `next = current + next`, we can compute `F(n)` efficiently.

## Approach

1. Initialize two variables `current` and `next` to represent the current Fibonacci number and the next Fibonacci number, respectively.
2. If `n` is 0, return 0, and if `n` is 1, return 1, as they are the base cases of the Fibonacci sequence.
3. Start iterating from 2 to `n` (inclusive) using a loop.
4. In each iteration, update `current` and `next` to represent the current Fibonacci number and the next Fibonacci number, respectively, based on the formula `next = current + next`.
5. After the loop, return the value of `next`, which represents `F(n)`.

## Time Complexity

The time complexity of this solution is O(n) because it iterates through `n` steps once to calculate the Fibonacci number for each step.

## Space Complexity

The space complexity of this solution is O(1) because it only uses a constant amount of extra space to store the `current` and `next` variables.

