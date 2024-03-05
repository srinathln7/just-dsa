# Validate a BST

Given the root of a binary tree, determine if it is a valid binary search tree (BST).

## Intuition:

A binary search tree (BST) follows a specific ordering property, where for each node, all elements in its left subtree are less than its value, and all elements in its right subtree are greater than its value. We can exploit this property to validate whether a given binary tree is a valid BST.

## Approach:

- Implement a recursive function `isValidBSTUtil` to validate whether a subtree is a valid BST.
- For each node in the binary tree:
  - Check if the node's value violates the BST property (less than the minimum allowed value or greater than the maximum allowed value).
  - Recursively check the left subtree with the updated maximum allowed value (current node's value) and the right subtree with the updated minimum allowed value (current node's value).
- If the above conditions are satisfied for all nodes, the binary tree is a valid BST.

## Time Complexity:

The time complexity of this approach is O(n), where n is the number of nodes in the binary tree. We visit each node once during the traversal.

## Space Complexity:

The space complexity is O(n), where n is the number of nodes in the binary tree. This space is used for the recursive call stack during the depth-first search traversal.