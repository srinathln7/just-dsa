# Convert Binary to String 
Given the root of a binary tree, construct a string consisting of parenthesis and integers from a binary tree with the following rules:

1. The root value of the tree is always placed in the middle of the string.
2. The string representation of the left subtree follows the same rules.
3. The string representation of the right subtree follows the same rules.
4. If a right subtree is empty but a left subtree is not, you still need to place a parenthesis () to indicate the absence of a right subtree.
5. Omit the unnecessary parenthesis () characters.

## Intuition
We can perform a depth-first search (DFS) traversal of the binary tree to construct the string representation according to the given rules.

## Approach
1. Implement a function `tree2str` that takes the root node of the binary tree as input.
2. Check if the root node is nil. If so, return an empty string.
3. Check if the root node has no left and right children. If so, return the string representation of the root node's value.
4. Define a recursive DFS function `dfs` that traverses the binary tree in preorder (root, left, right) manner and constructs the string representation.
5. In the `dfs` function:
   - Append the value of the current node to the result string.
   - If the current node has a left child or both left and right children, append "(" to the result string and recursively call `dfs` for the left child.
   - If the current node has both left and right children, append ")" after the left subtree is processed and recursively call `dfs` for the right child.
   - If the current node has no left child but has a right child, append "()" to indicate the absence of the left subtree and recursively call `dfs` for the right child.
6. Initialize an empty string `result` and invoke the `dfs` function with the root node and a pointer to the result string.
7. Return the constructed string representation.

## Time Complexity
The time complexity of the DFS traversal is O(N), where N is the number of nodes in the binary tree.

## Space Complexity
The space complexity is O(H), where H is the height of the binary tree, due to the recursion stack. In the worst case, the space complexity can be O(N) for a completely unbalanced binary tree.