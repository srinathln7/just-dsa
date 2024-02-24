**Problem Statement:**
Given an integer array `nums` and an integer `k`, the task is to determine if there are two distinct indices `i` and `j` in the array such that `nums[i] == nums[j]` and the absolute difference between `i` and `j` is at most `k`.

**Intuition:**
To efficiently solve this problem, we can utilize a sliding window approach along with a set data structure. We maintain a set to store the elements within the sliding window of size `k`. As we iterate through the array, we add each element to the set. If a duplicate element is found within the window, we return true. If no duplicates are found after iterating through the entire array, we return false.

**Approach:**
1. Initialize an empty set `window` to store elements within the sliding window.
2. Iterate through the array `nums` using a for loop with index `i`.
3. For each number `num` at index `i`:
   - Check if `num` exists in the `window` set. If found, return true as a duplicate within the window is found.
   - Add `num` to the `window` set.
   - If the size of the `window` set exceeds `k`, remove the leftmost element from the set to maintain the window size.
4. If no duplicates are found after iterating through the entire array, return false.

**Time Complexity:**
- The time complexity of this approach is O(n), where n is the size of the input array `nums`. This is because we iterate through the array once.
  
**Space Complexity:**
- The space complexity is O(k), where k is the size of the sliding window. We use a set to store at most `k` elements within the window.

This approach efficiently solves the problem while maintaining a bounded space complexity.