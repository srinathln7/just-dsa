# Rotate List
Given a singly-linked list `head` and an integer `k`, rotate the linked list to the right by `k` positions.

## Intuition
To rotate the linked list to the right by `k` positions, we need to find the `(n - k)`-th node from the beginning, where `n` is the length of the linked list. Then, we set this node's next pointer as `nil` to indicate the new end of the rotated list. Finally, we set the next pointer of the current last node to the original head of the list and return the new head of the rotated list.

## Approach
1. Handle the base cases: If the input list is empty, has only one node, or `k` is `0`, return the head unchanged.
2. Determine the length `n` of the linked list by traversing it.
3. Calculate the effective rotation count `k = k % n`. If `k` is `0`, no rotation is needed, so return the head unchanged.
4. Find the `(n - k)`-th node from the beginning by advancing a fast pointer `k` steps ahead.
5. Traverse until the last node using two pointers (slow and fast).
6. Set the next pointer of the slow pointer as `nil` to disconnect the rotated part from the original list.
7. Set the next pointer of the fast pointer to the head of the original list to complete the rotation.
8. Return the new head of the rotated list.

## Time Complexity
- Calculating the length of the linked list requires traversing it once, which takes O(n) time.
- Finding the `(n - k)`-th node from the beginning and traversing to the last node both take O(n - k) time.
- Thus, the overall time complexity is O(n), where n is the number of nodes in the linked list.

## Space Complexity
`O(1)`