# Design Linked List

## Problem

Design your implementation of a singly linked list.

## Approach

The provided code implements a singly linked list with the following functionalities:

- **Constructor:** Initializes an empty linked list.
- **Get(index int) int:** Gets the value of the index-th node in the linked list. If the index is invalid, returns -1.
- **AddAtHead(val int):** Adds a node of value val at the beginning of the linked list.
- **AddAtTail(val int):** Adds a node of value val at the end of the linked list.
- **AddAtIndex(index int, val int):** Adds a node of value val at the index-th position in the linked list. If the index is greater than the length of the linked list, the node will not be inserted.
- **DeleteAtIndex(index int):** Deletes the index-th node in the linked list, if the index is valid.

## Complexity Analysis

The time complexity for each operation is as follows:

- **Get:** O(n), where n is the number of nodes in the linked list.
- **AddAtHead:** O(1).
- **AddAtTail:** O(n), as it iterates through the entire list to find the tail node.
- **AddAtIndex:** O(n), as it may need to iterate through the list to find the correct position to insert the node.
- **DeleteAtIndex:** O(n), as it may need to iterate through the list to find the node to delete.

The space complexity for each operation is O(1), as it only uses a constant amount of extra space for pointers.




