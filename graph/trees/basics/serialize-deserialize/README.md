# Serialize and Deserialize binary trees

Design an algorithm to serialize and deserialize a binary tree.

## Intuition

To serialize and deserialize a binary tree, we can perform a preorder traversal of the tree. During serialization, we append the value of each node to a string. We also append a special character (e.g., "N") to represent null nodes. During deserialization, we split the string into tokens and recursively construct the binary tree.

## Approach

1. **Serialization**
   - Implement a function `serialize(root *TreeNode) string` that takes the root of the binary tree as input and returns the serialized string.
   - Perform a preorder traversal of the binary tree.
   - Append the value of each node to a slice, and append "N" for null nodes.
   - Join the slice into a string using the comma (",") as a delimiter and return the result.

2. **Deserialization**
   - Implement a function `deserialize(data string) *TreeNode` that takes the serialized string as input and returns the root of the binary tree.
   - Split the serialized string into tokens using the comma (",") as a delimiter.
   - Define a helper function `dfs()` to recursively construct the binary tree.
   - In the `dfs` function:
     - If the current token is "N", increment the index (`idx`) and return nil to represent a null node.
     - Otherwise, parse the token as an integer and create a new TreeNode with that value.
     - Recursively call `dfs` to construct the left and right subtrees.
     - Return the constructed node.
   - Return the result of the `dfs` function.

## Time Complexity

- Serialization: \(O(N)\) - Preorder traversal of the binary tree.
- Deserialization: \(O(N)\) - Traversing the serialized string and constructing the binary tree.

## Space Complexity

- Serialization: \(O(N)\) - Using a slice to store the serialized string.
- Deserialization: \(O(N)\) - Using recursion and stack space for constructing the binary tree.