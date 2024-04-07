# Non-Overlapping Intervals
Given an array of intervals `intervals` where `intervals[i] = [starti, endi]`, return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.

## Intuition
To minimize the number of intervals to remove, we can sort the intervals based on their end points and greedily remove intervals that overlap with the current interval. By doing so, we ensure that the intervals with smaller end points are kept whenever possible.

## Approach
1. Sort the intervals based on their end points in non-decreasing order.
2. Initialize variables `count` to 0 and `prevEnd` to the end point of the first interval.
3. Iterate through each interval `curr` starting from the second interval.
4. If the start point of `curr` is less than `prevEnd`, increment `count`.
5. Otherwise, update `prevEnd` to the end point of `curr`.
6. After iterating through all intervals, return `count`.

## Time Complexity
The time complexity of this approach is O(n log n), where n is the number of intervals in the array `intervals`. Sorting the intervals takes O(n log n) time.

## Space Complexity
The space complexity is O(1) since we use only a constant amount of extra space.
