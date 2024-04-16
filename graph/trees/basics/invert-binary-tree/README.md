# Invert binary tree
Given the root of a binary tree, invert the tree, i.e., swap the left and right subtrees of every node.

## Intuition
The intuition behind the problem is that the inverted tree is the mirror image of the original tree. We can achieve this by recursively inverting each subtree of the given tree.

## Approach
1. Start by defining a function `invertTree` that takes a pointer to the root of the binary tree as input and returns a pointer to the root of the inverted tree.
2. In the `invertTree` function:
   - Base cases:
     - If the root is nil, return nil.
     - If the root has no left or right child (i.e., it is a leaf node), return the root itself.
   - Otherwise, create a new node `newroot` with the same value as the root.
   - Recursively invert the left subtree of the original tree and assign it to the right child of `newroot`.
   - Recursively invert the right subtree of the original tree and assign it to the left child of `newroot`.
   - Finally, return `newroot`.
3. Call the `invertTree` function with the root of the binary tree to obtain the inverted tree.

## Time Complexity
The time complexity of this approach is O(n), where n is the number of nodes in the binary tree. This is because we visit each node exactly once.

## Space Complexity
The space complexity of this approach is O(n), where n is the number of nodes in the binary tree. This space is used for the recursive call stack.