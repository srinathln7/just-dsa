# Delete a node in the BST

Given a binary search tree (BST) and a key, the task is to delete the node with the given key while maintaining the BST property.

# Intuition

The intuition behind deleting a node in a binary search tree involves three cases:
1. If the node to be deleted has no children, it can be removed directly.
2. If it has only one child, that child can replace the node.
3. If it has two children, we find the minimum value in the right subtree (or maximum value in the left subtree), replace the node's value with this minimum value, and then recursively delete the minimum node from the right subtree.

# Approach

1. Initialize a function `deleteNode` that takes the root of the BST and the key to be deleted.
2. Use a switch statement to handle different cases based on the relationship between the key and the current node's value:
   - If the root is nil, return nil (base case).
   - If the key is less than the current node's value, recursively call `deleteNode` on the left subtree.
   - If the key is greater than the current node's value, recursively call `deleteNode` on the right subtree.
   - If the key matches the current node's value:
     - If the node has no left child, return the right child.
     - If the node has no right child, return the left child.
     - If the node has both left and right children, find the minimum node in the right subtree, replace the current node's value with the minimum node's value, and recursively delete the minimum node from the right subtree.
3. Define a helper function `findMin` to find the node with the minimum value in a given subtree.
4. Return the root node of the updated BST.

# Time Complexity

The time complexity of this approach is O(H), where `H=log(n)` is the height of the BST. In the worst case (skewed trees), the height of the tree is O(N), where N is the number of nodes in the tree. Thus, the time complexity is O(N).

# Space Complexity

The space complexity is O(H) due to the recursive calls, where `H=log(n)` is the height of the BST. In the worst case, the height of the tree is O(N), leading to a space complexity of O(N).

