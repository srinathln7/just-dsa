# Add Two Numbers (Linked List)

## Intuition:
The problem involves adding two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. We need to return a new linked list representing the sum of the two numbers.

## Approach:
To solve this problem, we can iterate through both linked lists simultaneously, starting from the least significant digit (the head of each list). We keep track of the carry generated by summing up digits at each position. We calculate the sum of digits and the carry, insert the sum (modulo 10) into the result linked list, and update the carry. If one linked list is shorter than the other, we continue the addition with a zero value. Finally, if there's still a carry after processing both lists, we append a new node with the carry to the end of the result list.

## Time Complexity:
The time complexity of this approach is O(max(m, n)), where m and n are the lengths of the input linked lists `l1` and `l2`, respectively. We iterate through both lists once, and the length of the resulting list is at most max(m, n) + 1.

## Space Complexity:
The space complexity is also O(max(m, n)), where m and n are the lengths of the input linked lists `l1` and `l2`, respectively. The space is primarily used for the output linked list, which can have a maximum length of max(m, n) + 1.