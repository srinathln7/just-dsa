# Concatenation of Array

## Problem

Given an integer array `nums` of length `n`, you need to create a new array `ans` of length `2*n`, where `ans[i] == nums[i % n]` and `ans[i + n] == nums[i % n]` for `0 <= i < n`. Refer [here](https://leetcode.com/problems/concatenation-of-array/)

## Approach:
The function initializes a new array `ans` with double the length of the input array `nums`. It then iterates through the original array, copying each element to its corresponding position in `ans` and also copying the same element to the position `i + n` in `ans`. This effectively creates the concatenated array as required.

## Time Complexity:
O(n)

## Space Complexity:
O(n)
