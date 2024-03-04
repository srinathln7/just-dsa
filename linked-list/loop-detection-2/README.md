# Detect Cycle in Linked List

## Intuition
The problem aims to detect the presence of a cycle in a given linked list. This is a classic problem in computer science, often solved using Floyd's Tortoise and Hare algorithm, also known as the two-pointer technique. The intuition behind this algorithm is that if there is a cycle in the linked list, then two pointers moving at different speeds will eventually meet inside the cycle. By using two pointers - one slow and one fast - traversing the linked list, we can detect the presence of a cycle.

## Approach
1. Initialize two pointers, slow and fast, both starting from the head of the linked list.
2. Move the slow pointer one step at a time and the fast pointer two steps at a time.
3. If the fast pointer ever reaches the end of the list or becomes nil, there is no cycle.
4. If the fast pointer catches up to the slow pointer at any point, there is a cycle.
5. Move the slow pointer back to the head of the linked list.
6. Move both slow and fast pointers at the same pace until they meet again. The meeting point is the start of the cycle.

## Time Complexity
The time complexity of this approach is O(n), where n is the number of nodes in the linked list. This is because both pointers traverse the linked list at most once.

## Space Complexity
The space complexity of this approach is O(1) since it uses only a constant amount of extra space for the two pointers.

