# Max-Depth of a binary tree
Given the root of a binary tree, return its maximum depth. A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

## Intuition
To find the maximum depth of a binary tree, we can use a recursive approach. At each node, we calculate the depth of its left and right subtrees recursively. The maximum depth of the tree is the maximum of the depths of its left and right subtrees, plus one for the current node.

## Approach
1. Define a function `maxDepth` that takes a pointer to the root of a binary tree as input and returns an integer representing the maximum depth of the tree.
2. In the `maxDepth` function:
   - Base cases:
     - If the root is nil, return 0 (the depth of an empty tree is 0).
     - If the root has no children (both left and right pointers are nil), return 1 (the depth of a tree with only one node is 1).
   - Recursively calculate the depths of the left and right subtrees:
     - Set `leftDepth` to the result of calling `maxDepth` with the left child of the root.
     - Set `rightDepth` to the result of calling `maxDepth` with the right child of the root.
   - Return the maximum of `leftDepth` and `rightDepth`, plus 1 (for the current node).
3. Call the `maxDepth` function with the root of the binary tree to find its maximum depth.

## Time Complexity
The time complexity of this approach is O(n), where n is the number of nodes in the binary tree. This is because we visit each node of the tree exactly once during the recursive traversal.

## Space Complexity
The space complexity of this approach is O(h), where h is the height of the binary tree. This is because the depth of the recursion is proportional to the height of the tree. In the worst case, the recursion stack may contain all the nodes of the longest path from the root to a leaf node.