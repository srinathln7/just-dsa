# Palindrome Linked List

## Intuition:
To check if a linked list is a palindrome, we can follow these steps:
1. Find the middle of the linked list using the two-pointer technique.
2. Reverse the second half of the linked list.
3. Compare the values of nodes from the first half with the reversed second half.
4. If all values match, the linked list is a palindrome.

## Approach:
1. We use the two-pointer technique to find the middle of the linked list.
2. After finding the middle, we reverse the second half of the linked list.
3. We then compare the values of nodes from the first half with the reversed second half.
4. If any values don't match, we return false; otherwise, we return true.

## Time Complexity:
 O(N), where N is the number of nodes in the linked list. We traverse the linked list twice, once to find the middle and once to reverse the second half.

## Space Complexity: 
O(1). We use only a constant amount of extra space for pointers and variables, regardless of the size of the input linked list.