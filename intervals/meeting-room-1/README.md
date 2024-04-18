# Meeting Room 

Given an array of meeting time intervals, determine if a person could attend all meetings.

## Intuition:

To check if a person can attend all meetings, we need to ensure that there are no overlaps between any two meetings.

## Approach:

1. Sort the intervals based on their starting time.
2. Iterate through the sorted intervals.
3. Check if the starting time of the current interval is less than the ending time of the previous interval.
4. If such a situation is encountered, it indicates an overlap, and we return `false`.
5. If no overlaps are found, return `true`.

## Time Complexity:

The time complexity of this approach is O(n log n), where n is the number of intervals. This complexity arises from the sorting step.

## Space Complexity:

The space complexity is O(1) because we are not using any extra space that scales with the input size. We are modifying the input array in place.