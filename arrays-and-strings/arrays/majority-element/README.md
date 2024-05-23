# [Majority Element](https://leetcode.com/problems/majority-element/)
Given an array `nums` of size `n`, return the majority element. The majority element is the element that appears more than `⌊n / 2⌋` times. You may assume that the majority element always exists in the array.

## Intuition
Since the majority element appears more than half the time in the array, it is guaranteed to be the element with the highest frequency. The Boyer-Moore Voting Algorithm is an efficient method to find this element with a linear time complexity and constant space complexity.

## Approach
The Boyer-Moore Voting Algorithm works as follows:
1. Initialize a `candidate` for the majority element and a `count` set to 0.
2. Iterate through the array:
   - If `count` is 0, set the current element as the `candidate`.
   - If the current element is the same as the `candidate`, increment the `count`.
   - Otherwise, decrement the `count`.
3. By the end of the iteration, the `candidate` will be the majority element.

## Time Complexity
- **O(n)**: We traverse the array once to find the majority element.

## Space Complexity
- **O(1)**: We use a constant amount of extra space.