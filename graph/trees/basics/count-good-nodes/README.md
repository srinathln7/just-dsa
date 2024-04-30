# Good nodes

Given a binary tree, return the number of good nodes.

## Intuition

To find the number of good nodes in a binary tree, we can perform a depth-first search (DFS) traversal of the tree. While traversing each node, we keep track of the maximum value encountered from the root to the current node. If the value of the current node is greater than or equal to this maximum value, it is considered a good node. By recursively traversing the tree and updating the maximum value, we can count the number of good nodes in the binary tree.

## Approach

1. Implement a function `goodNodes(root *TreeNode) int` that takes the root of the binary tree as input and returns the number of good nodes.
2. Initialize a variable `count` to keep track of the number of good nodes.
3. Implement a helper function `dfs(node *TreeNode, currMax int, count *int)` that performs a depth-first traversal of the tree.
4. In the `dfs` function:
    - Handle the base case where the current node is nil by returning.
    - Check if the value of the current node is greater than or equal to `currMax`.
        - If true, update `currMax` to the value of the current node and increment `count`.
    - Recursively call `dfs` for the left and right subtrees, passing the updated `currMax` and `count`.
5. Return the value of `count` as the result.

## Time Complexity

- \(O(N)\), where \(N\) is the number of nodes in the binary tree. We visit each node once during the depth-first traversal.

## Space Complexity

- \(O(H)\), where \(H\) is the height of the binary tree. The space complexity is determined by the recursion stack used for the depth-first traversal.
