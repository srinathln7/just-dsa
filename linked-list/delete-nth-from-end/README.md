# Remove Nth Node From End of List

Given the head of a linked list, remove the `n`th node from the end of the list and return its head.

## Intuition

### Main
To remove the nth node from the end of the linked list, we can use the two-pointer technique. We'll use two pointers, fast and slow, with a gap of n nodes between them. Then, as we move both pointers forward, when the fast pointer reaches the end of the list, the slow pointer will be pointing to the node just before the node to be removed. Once we identify the node before the one to be removed, we can simply adjust the pointers to skip the nth node.

## Alternate
To remove the `n`th node from the end of the linked list, we need to find the `N - n + 1`th node from the beginning, where `N` is the total number of nodes in the list. We can determine the length of the linked list by iterating through it once. Once we have the length of the list, we can easily find the node to be removed.

## Approach

1. Initialize two pointers, fast and slow, pointing to the head of the linked list.
2. Move the fast pointer forward by n nodes.
3. Move both pointers forward until the fast pointer reaches the end of the list.
4. At this point, the slow pointer will be pointing to the node just before the node to be removed.
5. Adjust the pointers to skip the nth node.
6. Return the updated head of the linked list.


1. Initialize a pointer `curr` to traverse the linked list and a variable `N` to store the length of the list.
2. Traverse the linked list while updating `N` to count the number of nodes.
3. Calculate the value of `k = N - n + 1`, which represents the position of the node to be removed from the beginning.
4. Handle edge cases:
   - If `n` is greater than `N` or the list is empty, return `nil`.
   - If the list contains only one node and `n` is 1, return `nil` as there are no nodes left in the list after removal.
5. If `k` is 1, it means the first node needs to be removed. In this case, update the `head` pointer to the next node and return.
6. Traverse the list again to find the `k - 1`th node from the beginning. Update the `Next` pointer of this node to skip the `k`th node.
7. Return the updated head of the linked list.

## Time Complexity

The time complexity of this approach is O(N), where N is the number of nodes in the linked list. We traverse the list twice, once to move the fast pointer to the end and once to find the node before the one to be removed.


The time complexity of this solution is O(N), where N is the number of nodes in the linked list. We need to traverse the list twice: once to find its length and again to remove the `n`th node from the end.

## Space Complexity

The space complexity is O(1) since we only use a constant amount of extra space regardless of the input size.


The space complexity of this solution is O(1) because it only uses a constant amount of extra space to store pointers and variables regardless of the size of the input linked list.


