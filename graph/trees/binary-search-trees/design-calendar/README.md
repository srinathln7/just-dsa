# Design Calendar

You are implementing a program to manage a calendar. Events can be added to the calendar if they do not cause a double booking. A double booking occurs when two events have a non-empty intersection of time intervals.

## Intuition

We can use a binary search tree (BST) to manage the calendar efficiently. The BST ensures that the events are stored in sorted order based on their start times. When adding a new event, we traverse the BST and check for overlapping time intervals.

## Approach

1. Initialize an empty calendar represented by a BST.
2. Implement the `Book` method to add new events to the calendar.
3. When adding a new event, traverse the BST recursively.
4. If the BST is empty, the new event becomes the root of the BST.
5. Otherwise, compare the start and end times of the new event with the current node in the BST.
6. If there is no overlap, insert the new event into the appropriate subtree based on the comparison.
7. If there is an overlap, return `false` to indicate a double booking.
8. Repeat steps 3-7 until the new event is successfully added or determined to cause a double booking.

## Time Complexity

- Adding an event (booking) to the calendar: O(log N), where N is the number of events in the calendar. This is the time complexity of searching and inserting into a balanced BST.

## Space Complexity

- Space complexity of the BST: O(N), where N is the number of events in the calendar. This is the space required to store the events in the BST.

