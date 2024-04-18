# Minimum Number of Conference Rooms
Given an array of meeting time intervals where `intervals[i] = [start_i, end_i]`, determine the minimum number of conference rooms required.

## Intuition
To find the minimum number of conference rooms required, we need to determine the maximum number of overlapping intervals at any point in time. This can be achieved by sorting the intervals based on their end times and then checking for overlaps.

## Approach
1. Sort the intervals based on their end times.
2. Remove duplicates if any.
3. Iterate through the sorted intervals and count the number of overlaps.
4. The minimum number of conference rooms required is the count of overlaps plus one.

## Time Complexity
The time complexity of this approach is O(n log n), where n is the number of intervals. This is due to the sorting step.

## Space Complexity
The space complexity is O(1) because we are not using any additional space proportional to the input size.
