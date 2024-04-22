# Reorder LinkedList 
Given the head of a singly linked-list, reorder the list to be on the following form: L0 → Ln → L1 → Ln-1 → L2 → Ln-2 → ...

## Input
- The head of a singly linked-list.

## Output
- No return value. The function should modify the given linked list in-place.

## Intuition
To reorder the linked list as described, we can follow these steps:
1. Find the middle of the linked list using the two-pointer technique (slow and fast pointers).
2. Reverse the second half of the linked list.
3. Merge the first half and the reversed second half by interleaving them.

## Approach
1. If the given linked list is empty or has only one node, return as no reordering is required.
2. Find the middle of the linked list using the slow and fast pointer technique.
3. Disconnect the first half from the second half by updating the next pointer of the node before the middle node to nil.
4. Reverse the second half of the linked list using the `revList` function.
5. Merge the first half and the reversed second half by interleaving them.
6. Update the head of the linked list to point to the modified reordered list.

## Time Complexity
The time complexity is O(n), where n is the number of nodes in the linked list. We traverse the entire list once to find the middle and reverse the second half.

## Space Complexity
The space complexity is O(1) as we use only a constant amount of extra space for pointers and variables.
