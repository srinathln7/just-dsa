# Construct Binary Tree from Preorder and Inorder Traversal
Given the preorder and inorder traversals of a binary tree, the task is to construct and return the binary tree.

## Intuition:
To construct the binary tree, we utilize the preorder traversal which provides the root of the tree and divide the inorder traversal into left and right subtrees. Recursively, we construct the left and right subtrees by determining their root nodes and corresponding inorder and preorder traversals.

## Approach:
1. If either the preorder or inorder traversal array is empty, return nil since there are no nodes to construct.
2. Extract the root node value from the first index of the preorder traversal array.
3. Find the index of the root node's value in the inorder traversal array, dividing it into left and right subtrees.
4. Recursively build the left subtree using the left part of both preorder and inorder traversals.
5. Recursively build the right subtree using the right part of both preorder and inorder traversals.
6. Return the root node.

## Time Complexity:
The time complexity of this approach is O(n), where n is the number of nodes in the binary tree. This is because each node is visited exactly once during the construction process.

## Space Complexity:
The space complexity of this approach is also O(n), where n is the number of nodes in the binary tree. This is due to the recursive calls on the stack and the space used by the created TreeNode structures.