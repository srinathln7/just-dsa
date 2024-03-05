# Insert a node in a BST

Given the root of a binary search tree (BST) and an integer `val`, the task is to insert `val` into the BST while maintaining its properties.

## Intuition:
The intuition behind inserting a node into a binary search tree is to recursively navigate through the tree based on the comparison of the value to be inserted (`val`) with the values of the nodes. At each step, we determine whether to traverse to the left subtree or the right subtree based on the comparison. If we reach a null node, we create a new node with the given value and insert it at the appropriate position.

## Approach:
1. Define a function `insertIntoBST` that takes the root of the BST and the value to be inserted (`val`).
2. Use a switch statement to handle different cases based on the relationship between `val` and the current node's value:
   - If the root is nil, return a new node with value `val`.
   - If `val` is less than the current node's value, recursively call `insertIntoBST` on the left subtree and update the `root.Left` pointer.
   - If `val` is greater than the current node's value, recursively call `insertIntoBST` on the right subtree and update the `root.Right` pointer.
3. Return the root node of the BST after insertion.

## Time Complexity:
The time complexity of this approach depends on the height of the BST. For a balanced BST it is `O(logn)` In the worst case (skewed trees), the height of the BST is O(N), where N is the number of nodes in the tree. Thus, the time complexity is O(n).

## Space Complexity:
The space complexity is O(H), where `H=log(n)` is the height of the BST, due to the recursive calls. In the worst case, the height of the tree is O(n), leading to a space complexity of O(n).
