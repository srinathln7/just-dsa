# Problem: Intersection of Two Linked Lists

## Intuition:
Given two singly-linked lists, we need to find and return the node where the two lists intersect. If the lists do not intersect, we return nil. We can solve this problem efficiently by utilizing the concept of two pointers and cyclic patterns.

## Approach:
We initialize two pointers `l1` and `l2` to the heads of the input lists `headA` and `headB`, respectively. We traverse both lists simultaneously, moving `l1` and `l2` to their next nodes until one of them reaches the end of its list. If `l1` reaches the end of list A, we reset it to the head of list B (creating a cyclic pattern), and similarly, if `l2` reaches the end of list B, we reset it to the head of list A. Note after every iteration , the two pointers are |`m-n`| steps apart. Refer returning `kth` element from the end of the linked list for a similar pattern. Eventually, `l1` and `l2` will meet at the intersection point, or they will both reach the end of their lists (in case of no intersection).

## Time Complexity:
The time complexity of this approach is O(m + n), where m and n are the lengths of the input linked lists `headA` and `headB`, respectively. We traverse both lists once to find the intersection point or to reach the end of both lists.

## Space Complexity:
The space complexity is O(1) since we only use a constant amount of extra space (two pointers) regardless of the input size.
