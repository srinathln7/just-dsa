# Sum from Root to Leaf 

You are given the root of a binary tree containing digits from 0 to 9 only. Each root-to-leaf path in the tree represents a number. Return the total sum of all root-to-leaf numbers. Test cases are generated so that the answer will fit in a 32-bit integer. A leaf node is a node with no children.

## Intuition

To compute the sum of all root-to-leaf numbers, we can perform a depth-first traversal of the binary tree, keeping track of the current number formed from the root to the current node. When we reach a leaf node, we add the current number to the total sum.

## Approach

1. Define a function `sumNumbers` that takes the root of a binary tree as input.
2. Implement a depth-first traversal function (`dfs`) that takes a node and the current number formed so far as input.
3. In the `dfs` function:
   - If the node is `nil`, return.
   - Update the current number by multiplying it by 10 and adding the value of the current node.
   - If the node is a leaf node (i.e., it has no children), add the current number to the total sum and return.
   - Recursively call `dfs` for the left and right children of the node.
4. Call the `dfs` function with the root of the binary tree and an initial current number of 0.
5. Return the total sum.

## Time Complexity

The time complexity of this approach is O(n), where n is the number of nodes in the binary tree. This is because we visit each node exactly once during the depth-first traversal.