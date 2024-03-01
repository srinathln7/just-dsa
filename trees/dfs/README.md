# In-order Traversal

Given the root of a binary tree, the task is to perform an inorder traversal of the tree and return the list of node values in sorted order.

## Intuition:
Inorder traversal of a binary tree involves visiting the nodes in the following order: left subtree, root, right subtree. This traversal method results in the nodes being visited in sorted order for a binary search tree (BST).

## Approach:
1. Define a function `inorderTraversal` that takes the root of the binary tree.
2. Use a switch statement to handle different cases based on the existence of the root and its children:
   - If the root is nil, return an empty list.
   - If the root has no children (i.e., leaf node), add its value to the result list.
   - If the root has children, recursively perform inorder traversal on the left subtree and right subtree.
   - Combine the results of inorder traversal of the left subtree, the value of the current root node, and inorder traversal of the right subtree.
3. Return the resulting list containing the inorder traversal of the binary tree.

## Time Complexity:
The time complexity of this approach is O(n), where N is the number of nodes in the binary tree. This is because each node is visited once during the traversal.

## Space Complexity:
The space complexity is also O(n) due to the recursive calls and the space required to store the result list. In the worst case, when the binary tree is skewed (i.e., every node has only one child), the space complexity may reach O(n).