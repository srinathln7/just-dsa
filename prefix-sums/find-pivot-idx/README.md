# Pivot Index

Given an integer array `nums`, the task is to find the index `i` such that the sum of all elements to the left of `i` is equal to the sum of all elements to the right of `i`. If no such index exists, return -1.

**Intuition:**

The solution follows a simple approach. We first calculate the total sum of all elements in the array. Then, we iterate through the array and maintain the sum of elements to the left of the current index (`leftSum`). If at any index `i`, the sum of elements to the left (`leftSum`) is equal to the total sum minus `leftSum` and the element at index `i`, we return `i` as the pivot index.

**Approach:**

1. Calculate the total sum of all elements in the array.
2. Iterate through the array:
   - At each index `i`, check if the sum of elements to the left of `i` (`leftSum`) is equal to the total sum minus `leftSum` and the element at index `i`.
   - If the condition is satisfied, return `i` as the pivot index.
   - Otherwise, update `leftSum` by adding the current element.
3. If no pivot index is found, return -1.

**Time Complexity:**

The time complexity of this approach is O(n), where n is the length of the input array `nums`. We traverse the array twice: once to calculate the total sum and once to find the pivot index.

**Space Complexity:**

The space complexity of this approach is O(1) as we only use a constant amount of extra space for variables `sum` and `leftSum`.



