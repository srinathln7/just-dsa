# Valid Parentheses

## Problem

Given a string `s` containing just the characters `'(', ')', '{', '}', '[' and ']'`, determine if the input string is valid.

An input string is valid if:

1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.

## Approach:
The function utilizes a stack to check the validity of the input string `s`. It iterates through each character in `s` and performs the following actions:

- If the character is an opening bracket (`'('`, `'['`, or `'{'`), it pushes the corresponding closing bracket onto the stack.
- If the character is a closing bracket, it checks if the stack is empty or if the top of the stack contains the corresponding opening bracket. If not, it returns `false`. Otherwise, it pops the top element from the stack.
- After iterating through all characters, if the stack is empty, it means all opening brackets were properly closed, and the function returns `true`. Otherwise, it returns `false`.

## Time Complexity:
O(n)

## Space Complexity:
O(n)

