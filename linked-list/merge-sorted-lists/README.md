# Merge Two Sorted Lists

## Problem

Given two sorted linked lists, merge them into a single sorted linked list and return it.

## Approach

The provided approach merges two sorted linked lists efficiently using the iterative method. It maintains a dummy node `sortedListNode` as the head of the merged list and a `tail` pointer to append nodes during merging.

1. **Initialization:** Initialize a dummy node `sortedListNode`.
2. **Merging Process:** 
   - Iterate through both lists while either list is not empty:
     - Compare the values of the current nodes of `list1` and `list2`.
     - Append the smaller value node to the merged list by setting `tail.Next` to the node.
     - Move the pointer of the selected list (`list1` or `list2`) to its next node.
     - Progress the `tail` pointer further.
3. **Append Remaining Nodes:** If any list still has remaining nodes after merging, append them directly to the merged list.
4. **Return Result:** Return the next node of the dummy node (`sortedListNode`), which represents the head of the merged list.

## Complexity Analysis

The time complexity of this approach is O(n + m), where n and m are the lengths of the input lists. It iterates through both lists once.
The space complexity is O(1), as it only uses a constant amount of extra space for pointers.

