# Merge Two Binary Trees 
Given two binary trees `root1` and `root2`, the task is to merge them into a new binary tree. The merging rule is that if two nodes overlap, the sum of their values is used as the value of the new node. Otherwise, the non-null node is used as the node of the new tree.

## Intuition:
- The solution utilizes a recursive approach to merge the two trees.
- It handles various cases, including when one of the trees is empty, when both trees are leaf nodes, and when there are non-null nodes in both trees.

## Approach:
1. Define a function `mergeTrees` that takes two binary tree nodes `root1` and `root2` as input and returns the merged binary tree.
2. Implement the merging logic recursively:
   - Base cases:
     - If `root2` is nil, return `root1`.
     - If `root1` is nil, return `root2`.
     - If both `root1` and `root2` are leaf nodes, create a new leaf node with the sum of their values and return it.
   - Create a new node `newRoot` with the sum of values from `root1` and `root2`.
   - Recursively merge the left subtrees of `root1` and `root2` and assign the result to `newRoot.Left`.
   - Recursively merge the right subtrees of `root1` and `root2` and assign the result to `newRoot.Right`.
3. Return the `newRoot` node as the merged binary tree.

## Time Complexity:
- The time complexity of this approach is O(min(n, m)), where n and m are the number of nodes in the two input binary trees. This is because the algorithm traverses each node once.

## Space Complexity:
- The space complexity is O(min(n, m)), where n and m are the number of nodes in the two input binary trees. This space is used for the recursive function calls and the space required for the output merged binary tree.

