# Remove Element

## Problem

This function removes all occurrences of a specified value `val` from an array `nums` in-place and returns the length of the resulting array.

## Approach:
The function initializes a variable `count` to keep track of the index where the next non-`val` element should be placed. It iterates through the array, and whenever it encounters an element that is not equal to `val`, it moves it to the position indicated by `count` and then increments `count`. By doing this, it effectively removes all occurrences of `val` from the array.

## Time Complexity:
O(n)

## Space Complexity:
O(1)
