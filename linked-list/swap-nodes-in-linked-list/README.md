# [Swap Nodes in a Linked List](https://leetcode.com/problems/swapping-nodes-in-a-linked-list/)

You are given the head of a linked list and an integer `k`. Return the head of the linked list after swapping the values of the `k`-th node from the beginning and the `k`-th node from the end (the list is 1-indexed).

## Intuition
To swap the `k`-th node from the beginning with the `k`-th node from the end, we need to identify both nodes. We can achieve this by first determining the length of the linked list, then locating the respective nodes, and finally swapping their values.

## Approach
1. **Calculate the Length of the List**:
   - Traverse the list to count the number of nodes.
2. **Identify the Nodes to Swap**:
   - Traverse the list to find the `k`-th node from the beginning.
   - Traverse the list to find the `k`-th node from the end. This node is the `(n - k + 1)`-th node from the beginning, where `n` is the length of the list.
3. **Swap Values**:
   - Swap the values of the two identified nodes.
4. **Return the Modified List**:
   - Return the head of the modified list.

## Time Complexity
- **O(n)**: We traverse the list to calculate its length and then traverse it again to locate the nodes to swap. Thus, the total time complexity is linear with respect to the number of nodes in the list.

## Space Complexity
- **O(1)**: We only use a few extra variables, which does not depend on the input size.

