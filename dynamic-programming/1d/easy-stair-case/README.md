# Climbing Stairs

## Problem Statement

Given `n` steps, you are climbing a staircase. Each time you can either climb 1 or 2 steps. In how many distinct ways can you reach the top?

## Intuition

The problem can be solved using dynamic programming. The intuition behind the solution is based on recognizing the pattern in climbing stairs. Since there are only two options for each step (taking one step or taking two steps), the number of unique ways to climb `n` steps is the sum of the number of unique ways to climb `n-1` steps and `n-2` steps. This is similar to the Fibonacci sequence, where each term is the sum of the two preceding terms.


```
                5
               / \
              4   3
             /\   /\           
```



## Approach

1. Initialize two variables `current` and `next` to represent the number of unique ways to climb the current step and the next step, respectively.
2. Start iterating from step 2 to `n` (inclusive) using a loop.
3. In each iteration, update `current` and `next` to represent the number of unique ways to climb the current step and the next step, respectively, based on the formula `next = current + next`.
4. After the loop, return the value of `current`, which represents the number of unique ways to climb `n` steps.

## Time Complexity

The time complexity of this solution is O(n) because it iterates through `n` steps once to calculate the number of unique ways to climb each step.

## Space Complexity

The space complexity of this solution is O(1) because it only uses a constant amount of extra space to store the `current` and `next` variables.


