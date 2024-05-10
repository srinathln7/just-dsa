# Swap Pairs
Given a singly-linked list, swap every pair of adjacent nodes and return its head.

## Intuition:
The problem can be solved efficiently using a recursive approach. By swapping each pair of adjacent nodes recursively, we can obtain the desired result.

## Approach:
1. **Base Case:** If the input head is nil or there is only one node in the list, return the head as it is.
2. Initialize a variable `next` to the next node after the head.
3. Recursively call `swapPairs` on the node after the next node (`next.Next`). This will return the head of the swapped pairs.
4. Swap the current pair of nodes by pointing `head.Next` to the swapped pairs and `next.Next` to the `head`.
5. Return `next` as the new head of the swapped list.

## Time Complexity: 
O(N), where N is the number of nodes in the linked list. The algorithm traverses each node once.

## Space Complexity: 
O(N), where N is the number of nodes in the linked list. The space complexity is due to the recursive stack space used during the recursive calls.