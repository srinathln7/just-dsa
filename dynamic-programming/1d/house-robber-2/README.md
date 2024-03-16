# House Robber 2
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected, and it will automatically contact the police if two adjacent houses were broken into on the same night. You can assume that all houses are arranged in a circle, which means the first house is adjacent to the last house.

Given an integer array `nums` representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.

## Intuition
This problem is a variation of the House Robber problem, where you need to maximize the amount of money robbed without triggering the security system. We can adopt a bottom-up dynamic programming approach similar to House Robber problem 1. However, since the houses are arranged in a circle, we need to consider two options:
1. Start robbing from the first house (index 0) in a clockwise direction without alerting the police.
2. Start robbing from the last house (index `n-1`) in an anti-clockwise direction without alerting the police.

## Approach
1. Check the base cases:
   - If there are no houses, return 0.
   - If there is only one house, return the amount of money in that house.
2. Initialize variables `prevMax` and `currMax` to track the maximum amount of money robbed from the first and second houses, respectively.
3. Iterate over the houses from the second house (index 1) to the second-to-last house (index `n-2`):
   - Calculate the maximum amount of money that can be robbed from the current house:
     - `currMax = max(currMax, prevMax + nums[i])`
     - Here, `prevMax` represents the maximum amount of money robbed from the previous house, and `nums[i]` represents the amount of money in the current house.
     - Update `prevMax` to store the previous value of `currMax`.
4. Store the maximum amount of money robbed in a clockwise direction (`maxMoneyForward`) obtained from the above steps.
5. Repeat steps 2-4, but this time start robbing from the last house (index `n-1`) in an anti-clockwise direction.
6. Return the maximum of `maxMoneyForward` and the maximum amount of money robbed in the anti-clockwise direction.

## Time Complexity
The time complexity of this solution is O(n), where n is the number of houses.

## Space Complexity
The space complexity of this solution is O(1), as we are using only a constant amount of extra space regardless of the input size.
