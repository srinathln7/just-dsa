# Search an element in a BST

Given the root of a binary search tree (BST) and an integer `val`, the task is to search for the node with the value `val` in the BST and return the node if found, otherwise return nil.

## Intuition
The intuition behind searching for a node in a binary search tree involves recursively navigating through the tree based on the comparison of the target value (`val`) with the values of the nodes. At each step, we compare the target value with the value of the current node. If the target value matches the current node's value, we return the current node. Otherwise, we traverse either to the left subtree or the right subtree based on the comparison result.

## Approach
1. Define a function `searchBST` that takes the root of the BST and the target value `val`.
2. Use a switch statement to handle different cases based on the relationship between `val` and the current node's value:
   - If the root is nil, return nil indicating that the target value is not found in the BST.
   - If `val` is equal to the current node's value, return the current node.
   - If `val` is less than the current node's value, recursively call `searchBST` on the left subtree.
   - If `val` is greater than the current node's value, recursively call `searchBST` on the right subtree.
3. Return nil if the target value is not found in the BST.

## Time Complexity:
The time complexity of this approach depends on the height of the BST `O(logn)`. 

## Space Complexity:
`O(1)`
