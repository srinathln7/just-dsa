# Reverse LinkedList

LinkedList `1->2->3`


## Problem

Given a singly linked list, reverse it and return the reversed list.

## Iterative Approach

### Idea
We can reverse the linked list iteratively by maintaining three pointers: `head`, `newHead`, and `next`. We iterate through the list, reversing the pointers as we traverse it.

### Implementation
1. Initialize `newHead` as `nil`.
2. Iterate through the list:
   - Store the next node of the current node (`next = head.Next`).
   - Set the `Next` pointer of the current node (`head`) to `newHead`.
   - Update `newHead` to be the current node.
   - Move `head` to the next node (`head = next`).
3. Return `newHead` as the new head of the reversed list.

### Time Complexity
O(n), where n is the number of nodes in the linked list.

### Space Complexity
O(1)

## Recursive Approach

### Idea
We can reverse the linked list recursively by traversing to the end of the list, then adjusting the pointers as we backtrack.

### Base Cases
- If the current node is `nil` or the next node is `nil`, return the current node (head) itself.

### Implementation
1. Recursively call `RreverseList` on the next node to reverse the rest of the list.
2. Adjust the pointers to reverse the current node and its next node.
3. Return the new head of the reversed list.

### Time Complexity
O(n), where n is the number of nodes in the linked list.

### Space Complexity
O(n), due to the recursive call stack.



### Simulation

1. **Initial Call:**
   - We start with the initial call to `reverseList(1)`.
   - This call triggers a recursive call to `reverseList(2)`.

2. **Recursive Call with Node `2`:**
   - We make a recursive call to `reverseList(2)` with `head` as the node containing `2`.
   - This call triggers another recursive call to `reverseList(3)`.

3. **Recursive Call with Node `3`:**
   - We make a recursive call to `reverseList(3)` with `head` as the node containing `3`.
   - As `head.Next` is `nil`, indicating the end of the list, this call returns `3` as the new head of the reversed list.

4. **Reversing Pointers for Node `2`:**
   - After the recursive call with node `3` returns, we are back at the call with node `2`.
   - We set `2.Next.Next` to `2`, effectively reversing the connection from `2->3` to `3->2`.
   - Then, we set `2.Next` to `nil` to make `2` the new tail of the list.
   - We return `3`, which is now the new head of the reversed list.

In summary, the reversal of the connection from `2->3` to `3->2` happens implicitly during the recursive calls and adjustments of the `Next` pointers within the `reverseList` function.
