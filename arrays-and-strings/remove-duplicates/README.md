# Remove Duplicates from Sorted Array

## Problem

[Remove duplicates from sorted array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/)

## Intuition:
This solution utilizes two pointers to remove duplicates in a sorted array in-place. The variable `j` represents the index where there is no duplicate element.

## Approach:
We initialize `j` to 1 since the first element is always unique. Then, we iterate through the array using pointer `i`. If the current element `nums[i]` is not equal to the previous element `nums[i-1]`, it means it's a unique element, so we move it to the position represented by `j`, and then increment `j`.

## Time Complexity:
O(n)

## Space Complexity:
O(1)
