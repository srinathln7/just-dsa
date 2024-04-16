# Balanced Binary Tree 
Given a binary tree, determine if it is height-balanced. A height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differs by more than one.

## Intuition
To check if a binary tree is height-balanced, we need to ensure that the depth of the left and right subtrees of every node differs by at most one. This can be achieved by recursively calculating the depths of each subtree while simultaneously checking for balance. If at any point we encounter a subtree that is unbalanced (i.e., the difference in depths exceeds one), we can immediately return `false` as the tree is not height-balanced.

## Approach
1. Define a helper function `dfs` that takes a node as input and returns the depth of the subtree rooted at that node along with a boolean indicating whether the subtree is balanced.
2. In the `dfs` function:
   - If the node is `nil`, return depth 0 and `true` for balance.
   - Recursively calculate the depths of the left and right subtrees using `dfs`.
   - Check if either subtree is unbalanced or if the difference in depths exceeds 1. If so, return `-1` for depth and `false` for balance.
   - Otherwise, calculate the depth of the current subtree as 1 plus the maximum depth of the left and right subtrees.
   - Return the depth of the current subtree and `true` for balance.
3. Invoke the `dfs` function on the root node of the binary tree.
4. If the root node is `nil`, return `true` as an empty tree is considered height-balanced.
5. Otherwise, return the balance status obtained from the `dfs` function.

## Time Complexity
The time complexity of this approach is O(n), where n is the number of nodes in the binary tree. We traverse each node once to calculate the depths and check for balance.

## Space Complexity
The space complexity of this approach is O(h), where h is the height of the binary tree. This space is used for the recursive function calls on the call stack. In the worst case, the height of the binary tree could be equal to the number of nodes in the tree, resulting in O(n) space complexity. However, in a balanced binary tree, the height is logarithmic with respect to the number of nodes, leading to O(log n) space complexity.