# Insertion Sort
The task is to perform insertion sort on a singly-linked list.

## Intuition:
Insertion sort works by iteratively inserting each element from an unsorted portion of the list into its correct position in the sorted portion of the list.

## Approach:
1. We start by initializing a sentinel node, which serves as the starting point of the sorted list. This simplifies the insertion process, especially for edge cases.
2. We iterate through each node in the input list, starting from the second node (`curr`).
3. For each node `curr`, we compare its value with the values of nodes in the sorted portion of the list (from the sentinel node).
4. If `curr`'s value is greater than or equal to the value of the previous node (`prev`), we continue to the next iteration.
5. If `curr`'s value is less than the value of the previous node, we need to insert `curr` into the sorted portion of the list.
6. We traverse the sorted portion of the list from the sentinel node (`tmp`) until we find the correct position to insert `curr`.
7. Once we find the correct position (`tmp`), we update the pointers to insert `curr` into the sorted list:
   - We disconnect `curr` from its current position by updating `prev.Next`.
   - We connect `curr` to the sorted list by updating `tmp.Next`.
8. After inserting `curr`, we move to the next node in the input list and repeat the process.
9. Finally, we return the `sentinel.Next`, which points to the head of the sorted list.

## Time Complexity:
- The time complexity of this algorithm is O(n^2), where n is the number of nodes in the input list. This is because, in the worst case, we may have to traverse the sorted portion of the list for each node.

## Space Complexity:
- The space complexity is O(1) since we only use a constant amount of extra space for storing temporary variables and pointers.