# Find Duplicate Number in Array

## Intuition
The problem can be framed as a linked list cycle detection problem. By treating the array as a linked list where each element points to the value at its index, and considering the integer range as the index, we can detect the duplicate number. The algorithm utilizes Floyd's Tortoise and Hare algorithm to efficiently find the duplicate number with constant extra space.

## Approach
1. Initialize two pointers, slow and fast, pointing to the first element of the array.
2. Advance the slow pointer by one step and the fast pointer by two steps until they meet at an intersection point inside the cycle formed by the repeated number.
3. Reset one pointer to the beginning of the array and move both pointers at the same pace (one step at a time) until they meet again. This meeting point represents the start of the cycle, which corresponds to the duplicate number in the array.
4. Return the duplicate number.

## Time Complexity
The algorithm traverses the array twice with pointers, resulting in a linear time complexity of O(n), where n is the length of the array.

## Space Complexity
The algorithm utilizes only constant extra space for storing pointers, resulting in a space complexity of O(1).

