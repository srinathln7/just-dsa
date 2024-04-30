# Binary Tree Maximum Path Sum

Given a binary tree, find the maximum path sum.

## Intuition

To find the maximum path sum in a binary tree, we can use a depth-first search (DFS) approach. At each node, we calculate the maximum path sum starting from that node and update the global maximum path sum accordingly. 

## Approach

1. Implement a function `maxPathSum(root *TreeNode) int` that takes the root of the binary tree as input and returns the maximum path sum.
2. In the `maxPathSum` function, handle the base case where the root is nil by returning 0.
3. Initialize a variable `maxSum` to the minimum integer value to account for cases where all node values are negative.
4. Implement a helper function `dfs(curr *TreeNode, maxSum *int) int` that performs a depth-first traversal of the tree.
5. In the `dfs` function:
    - Handle the base case where the current node is nil by returning 0.
    - Calculate the maximum path sum in the left and right subtrees. If the sum is negative in either subtree, set it to 0 since we are interested in finding the maximum path sum.
    - Update the `maxSum` variable by adding the current node's value to the maximum path sum in the left and right subtrees.
    - Return the maximum path sum starting from the current node.
6. Return the `maxSum` as the result.

## Time Complexity

- \(O(N)\), where \(N\) is the number of nodes in the binary tree. We visit each node once during the depth-first traversal.

## Space Complexity

- \(O(H)\), where \(H\) is the height of the binary tree. The space complexity is determined by the recursion stack used for the depth-first traversal.