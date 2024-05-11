# Delete Duplicates in Linked List 
Given a sorted singly-linked list, remove duplicates such that each element appears only once.

## Intuition
Since the given list is already sorted, duplicates will appear consecutively. Therefore, we can iterate through the list, comparing each element with its next element. If the values are equal, we skip the next node to remove duplicates. If they are not equal, we move to the next node.

## Approach
1. Handle the base cases: If the list is empty or has only one node, return the list unchanged.
2. Initialize a pointer `curr` to the head of the list.
3. Iterate through the list:
   - If `curr` and its next node (`curr.Next`) have equal values, skip the next node by updating `curr.Next` to `curr.Next.Next`.
   - Otherwise, move `curr` to its next node (`curr = curr.Next`).
4. Continue this process until `curr` reaches the end of the list.
5. Return the head of the modified list.

## Time Complexity
- Traversing the entire linked list once takes O(n) time, where n is the number of nodes in the list.
- In each iteration, we perform constant time operations (comparisons and pointer updates).
- Thus, the overall time complexity is O(n).

