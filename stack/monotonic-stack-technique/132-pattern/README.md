# Find 132 Pattern

Given an array of n integers nums, a 132 pattern is a subsequence of three integers nums[i], nums[j] and nums[k] such that i < j < k and nums[i] < nums[k] < nums[j].

Return true if there is a 132 pattern in nums, otherwise, return false.

 
## Intuition:
To find a 132 pattern, we aim to identify a number `k` that is greater than the minimum element to its left (`min`), yet smaller than the numbers to its right. We maintain a monotonic decreasing stack to efficiently identify such candidates.

## Approach:
1. We iterate through the array and maintain a stack to store elements along with their corresponding minimum element to the left.
2. At each step, we compare the current number with the top of the stack.
3. If the current number is greater than or equal to the top of the stack, we pop elements from the stack until the top element is less than the current number. This ensures that the stack remains monotonically decreasing.
4. Next, we check if the current number is greater than the minimum element to its left and less than the top element of the stack. If this condition is satisfied, we have found a 132 pattern.
5. If no such pattern is found after iterating through the array, we return false.

## Time Complexity: 
O(N), where N is the length of the input array `nums`. We iterate through the array once.

## Space Complexity: 
O(N), where N is the length of the input array `nums`. The space complexity is dominated by the size of the stack.