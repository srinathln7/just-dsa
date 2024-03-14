# 1s and 0s
You are given an array of strings `strs`, where `strs[i]` consists of `0`s and `1`s only, and two integers `m` and `n`. Your task is to find the maximum number of strings that you can form with `m` `0`s and `n` `1`s by picking strings from the array `strs`. Each string can be used at most once.

## Intuition
This problem can be solved using dynamic programming with memoization. We can construct a decision tree where each node represents a choice of either including or excluding a string from the array. At each node, we check if the current string can fit into the remaining `0`s and `1`s available (`m` and `n` respectively). If it can, we recurse to the next level, subtracting the number of `0`s and `1`s in the current string from `m` and `n` respectively. We memoize the results to avoid redundant calculations.

## Approach
1. Initialize a memoization map `dp` to store results for each state `(m, n, idx)`, where `m` represents the remaining number of `0`s, `n` represents the remaining number of `1`s, and `idx` represents the current index in the array `strs`.
2. Define a recursive function `dfs` that takes `m`, `n`, and `idx` as arguments and returns the maximum number of strings that can be formed.
3. In the `dfs` function:
   - Base case: If `idx` reaches the end of the array `strs`, return 0.
   - Check if the result for the current state `(m, n, idx)` exists in the memoization map. If it does, return the stored value.
   - Make two decisions:
     - Decision 1: Skip the current string and move to the next index.
     - Decision 2: Include the current string if there are enough `0`s and `1`s available. Recurse to the next level, subtracting the number of `0`s and `1`s in the current string from `m` and `n` respectively.
   - Update the memoization map with the maximum value between the two decisions.
   - Return the result for the current state.
4. Return the result of the `dfs` function with initial parameters `m`, `n`, and `idx` set to 0.

## Time Complexity
The time complexity of this solution is O(len(strs) * m * n), where len(strs) is the length of the input array `strs`, and `m` and `n` are the given integers.

## Space Complexity
The space complexity is O(len(strs) * m * n) due to the memoization map `dp`.
