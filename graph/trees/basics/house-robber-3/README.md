# House Robber 3 
The problem is to find the maximum amount of money the thief can rob without alerting the police in a binary tree of houses. Each house in the tree has a certain amount of money, and the thief cannot rob two directly-linked houses on the same night.

## Intuition
At each node in the binary tree, the thief has two options: to include the current house in the robbery or to exclude it. The decision to include a house means the thief cannot rob the immediate children of that house. The decision to exclude a house means the thief can consider robbing its immediate children. By recursively exploring all possible combinations of including and excluding houses, we can determine the maximum amount of money the thief can rob.

## Approach
1. Define a helper function `dfs(node *TreeNode) (int, int)` that takes a node in the binary tree and returns two values: the maximum amount of money that can be robbed when including the current node (`incl`) and the maximum amount of money that can be robbed when excluding the current node (`excl`).
2. In the `dfs` function:
   - Base case: If the node is nil, return `(0, 0)`.
   - Recursively calculate `incl` and `excl` for the left and right subtrees.
   - Calculate `incl` by adding the current node's value to the sum of `leftExcl` and `rightExcl`.
   - Calculate `excl` by taking the maximum of `leftIncl` and `leftExcl`, and adding it to the maximum of `rightIncl` and `rightExcl`.
3. Call the `dfs` function with the root of the binary tree and return the maximum of `incl` and `excl`.

## Time Complexity
The time complexity of this approach is O(n), where n is the number of nodes in the binary tree. This is because we traverse each node once in the tree during the depth-first search.

## Space Complexity
The space complexity of this approach is O(h), where h is the height of the binary tree. This is because the recursion stack can go as deep as the height of the tree during the depth-first search.