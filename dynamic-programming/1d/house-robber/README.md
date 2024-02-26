# House Robber Problem

You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, but adjacent houses have security systems connected. You cannot rob two adjacent houses on the same night without alerting the police. Given an integer array `nums` representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.

## Intuition
  - The problem can be efficiently solved using dynamic programming.
  - At each house `i`, we need to decide whether to rob it or skip it.
  - We keep track of the maximum amount of money we can have if we had robbed until the previous house (`prevMax`) and until the current house (`currMax`). For this we initialize two variables: prevMax and currMax. prevMax will store the maximum amount of money that can be robbed up to the previous house without alerting the police, and currMax will store the maximum amount of money that can be robbed up to the current house without alerting the police.

## Approach:
  1. **Handle edge cases**:
     - If there are no houses (`len(nums) == 0`), return 0.
     - If there is only one house (`len(nums) == 1`), return the money amount of that house.
  2. **Initialize** `prevMax` to `nums[0]` and `currMax` to the maximum of the first two houses (`max(nums[0], nums[1])`).
  3. **Iterate** through the array of house money amounts from the third house (`i = 2`) to the last house (`i = n-1`).
  4. **For each house `i`**, update `currMax` as the maximum of:
     - The money amount of the current house plus the money amount of the house two houses before (`nums[i] + prevMax`).
     - The maximum amount of money until the previous house (`currMax`).
  5. **Update** `prevMax` to be the old value of `currMax`.
  6. After iterating through all the houses, **return** `currMax`, which represents the maximum amount of money that can be robbed without alerting the police.

## Time Complexity 
The time complexity of the solution is `O(n)`, where `n` is the number of houses.

## Space Complexity
The space complexity of the solution is `O(1)` since we only use a constant amount of extra space for variables.

