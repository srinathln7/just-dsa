# Merge Intervals 
Given an array of intervals `intervals` where `intervals[i] = [starti, endi]`, merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

## Intuition
To merge overlapping intervals efficiently, we can first sort the intervals based on their start points. Then, we can iterate through the sorted intervals and merge overlapping intervals.

## Approach
1. Sort the intervals based on their start points.
2. Initialize an empty array `merged` to store the merged intervals.
3. Iterate through each interval `curr` in the sorted intervals.
4. If the `merged` array is empty or the end point of the last interval in `merged` is less than the start point of `curr`, add `curr` to `merged`.
5. If there is an overlap between `curr` and the last interval in `merged`, update the end point of the last interval in `merged` to cover both intervals.
6. After iterating through all intervals, return `merged`.

## Time Complexity
The time complexity of this approach is O(n log n), where n is the number of intervals in the array `intervals`. Sorting the intervals takes O(n log n) time, and iterating through the sorted intervals takes linear time.

## Space Complexity
The space complexity is O(n) since we need to store the merged intervals in the `merged` array, and in the worst case, it can contain all intervals from the input array `intervals`.
