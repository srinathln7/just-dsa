# Problem Statement
Given the root of a binary tree and an integer targetSum, return true if the tree has a root-to-leaf path such that adding up all the values along the path equals targetSum. A leaf is a node with no children.

# Intuition
This problem can be solved using backtracking, where we traverse the binary tree recursively from the root to the leaf nodes. At each node, we subtract the node's value from the target sum and check if it reaches zero when we reach a leaf node. If it does, we have found a valid path.

# Approach
We traverse the binary tree recursively using a depth-first search (DFS) approach. At each node, we subtract the node's value from the target sum and continue the traversal in both the left and right subtrees. We also maintain a stack to keep track of the current path. If we encounter a leaf node and the target sum becomes zero, we return true, indicating that a valid path exists. Otherwise, if the traversal reaches a leaf node or a null node, we backtrack by removing the last element from the stack and continue the traversal.

# Time Complexity
The time complexity of this approach is O(n), where n is the number of nodes in the binary tree. This is because we visit each node once during the traversal.

# Space Complexity
The space complexity is also O(logn), which is the height of the binary tree. This is because the maximum space used by the call stack during recursion is proportional to the height of the tree.
