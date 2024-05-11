# Reverse Nodes in a Linked List II

## Problem Statement

Given the head of a singly linked list and two integers `left` and `right` where `left` <= `right`, reverse the nodes of the list from position `left` to position `right`, and return the reversed list.

## Intuition

To reverse a portion of a linked list, we need to traverse the list until the node at position `left - 1`, then reverse the nodes from position `left` to `right`, and finally connect the reversed portion back to the original list.

## Approach

1. Initialize a sentinel node to handle cases where the reversal starts from the beginning of the list.
2. Traverse the list until the node at position `left - 1` and maintain a pointer to it.
3. Reverse the segment of the list between positions `left` and `right` by iteratively changing the pointers of the nodes within that segment.
4. Update the pointers of the adjacent nodes to maintain the connectivity of the list.
5. Return the head of the modified list.

## Time Complexity

The time complexity is O(n), where n is the number of nodes in the linked list. This is because we only traverse the list once to perform the reversal.

## Space Complexity

The space complexity is O(1) since we use a constant amount of extra space regardless of the input size.
