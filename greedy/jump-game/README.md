# Jump Game
Given an integer array `nums`, where each element represents your maximum jump length at that position, determine if you can reach the last index.

## Intuition
This algorithm iterates through the array from right to left. It maintains a variable `goal` which represents the last index (goal) that can reach the end of the array. If an index `i` plus the jump length at that index can reach `goal`, then `goal` is updated to `i`, indicating that `i` can reach the end of the array. If `goal` is ultimately updated to 0, it means that the first index is reachable from the last index.

## Approach
* Initialize goal as the last index of the array.
* Iterate through the array from right to left.
* For each index i, calculate maxJump as i + nums[i], which represents the maximum index that can be reached from index i.
* If maxJump is greater than or equal to goal, update goal to i, indicating that i can reach the end of the array.
* After the loop completes, check if goal is 0. If so, return true, indicating that the first index is reachable from the last index.
* If goal is not 0, return false.

## Time Complexity
The time complexity of this approach is O(n), where n is the length of the input array `nums`. This is because we iterate through the array once.

## Space Complexity
The space complexity is O(1) because we use only a constant amount of extra space for variables regardless of the size of the input array.