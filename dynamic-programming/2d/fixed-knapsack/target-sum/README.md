# Target Sum
Given an array `nums` containing integers and a target integer, find the number of ways to form the target sum by adding or subtracting the elements of `nums`.

## Intuition
We can use a recursive approach combined with memoization to solve this problem efficiently. At each step, we have two choices: either add the current element or subtract it from the running total. We'll recursively explore both options until we reach the end of the array. Memoization will be used to avoid redundant calculations.

## Approach
1. Initialize a memoization map to store previously computed results. The map key is a tuple `(index, total)`, where `index` represents the current index in the `nums` array and `total` represents the running total.
2. Define a recursive function `dfsBT` that takes the current index `idx` and the current total `total`.
3. In the recursive function:
   - Base Case: If `idx` equals the length of `nums`, return 1 if `total` equals the target; otherwise, return 0.
   - If the current `(idx, total)` pair exists in the memoization map, return the cached result.
   - Recursively call `dfsBT` with `idx+1` and `total+nums[idx]` (adding the current element) and with `total-nums[idx]` (subtracting the current element).
   - Update the memoization map with the result and return it.
4. Call the recursive function `dfsBT` with initial parameters `idx = 0` and `total = 0`.
5. Return the result obtained from the recursive function.

## Time Complexity
- The time complexity of this solution is O(N * M), where N is the length of the `nums` array and M is the absolute value of the target sum. This is because there are N * M possible combinations of index and total values.
- However, due to memoization, the effective time complexity is reduced significantly, as many subproblems are computed only once.

## Space Complexity
- The space complexity is O(N * M) due to the memoization map, where N is the length of the `nums` array and M is the absolute value of the target sum.