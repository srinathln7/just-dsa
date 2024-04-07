# Add Intervals
You are given an array of non-overlapping intervals `intervals`, where `intervals[i] = [starti, endi]` represent the start and the end of the ith interval. The array `intervals` is sorted in ascending order by `starti`. You are also given an interval `newInterval = [start, end]` that represents the start and end of another interval.
Insert `newInterval` into `intervals` such that `intervals` is still sorted in ascending order by `starti` and `intervals` still does not have any overlapping intervals (merge overlapping intervals if necessary). Return `intervals` after the insertion.

## Intuition
Since the intervals are already sorted by their start points, we can efficiently find the correct position to insert the new interval by iterating through the intervals and comparing their start points with the start point of the new interval.

## Approach
1. Initialize an empty array `result` to store the merged intervals.
2. Iterate through each interval `curr` in `intervals`.
3. If the end point of `curr` is less than the start point of `newInterval`, add `curr` to `result`.
4. If the start point of `curr` is greater than the end point of `newInterval`, add `newInterval` to `result` and set `newInterval` to `curr`.
5. If there is an overlap between `curr` and `newInterval`, update the start and end points of `newInterval` to cover both intervals.
6. After iterating through all intervals, add the remaining `newInterval` to `result`.
7. Return `result`.

## Time Complexity
The time complexity of this approach is O(n), where n is the number of intervals in the array `intervals`. We iterate through each interval once.

## Space Complexity
The space complexity is O(n) since we may need to store all intervals in the merged result array, and in the worst case, it can contain all intervals from the input array `intervals`.
