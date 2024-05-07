# Trim BST
Given the root of a binary search tree and the lowest and highest boundaries as `low` and `high`, trim the tree so that all its elements lie in the range `[low, high]`. Trimming the tree should not change the relative structure of the elements that will remain in the tree (i.e., any node's descendant should remain a descendant). It can be proven that there is a unique answer. Return the root of the trimmed binary search tree. Note that the root may change depending on the given bounds.

## Approach
The key idea is to narrow down the tree according to the range `[low, high]` while utilizing the property of a binary search tree (BST), where all values in the left subtree are less than the root value, and all values in the right subtree are greater than the root value.

1. **Base Case**: If the root is nil, return nil (empty tree).
2. **Root Value Exceeds High**: If the root's value itself exceeds the upper bound (`high`), then all values to the right will naturally exceed `high`. Hence, there is no need to trim the right subtree. Recursively trim the left subtree by calling `trimBST(root.Left, low, high)`.
3. **Root Value Subsides Low**: If the root's value itself is less than the lower bound (`low`), then all values to the left will be naturally lower than `low`. Hence, there is no need to trim the left subtree. Recursively trim the right subtree by calling `trimBST(root.Right, low, high)`.
4. **Within Range**: If the root's value is within the range `[low, high]`, then recursively trim both the left and right subtrees by calling `trimBST(root.Left, low, high)` and `trimBST(root.Right, low, high)`.
5. **Return**: Finally, return the root of the trimmed binary search tree.

## Time Complexity
The time complexity of this approach is O(N), where N is the number of nodes in the binary search tree.

## Space Complexity
The space complexity of this approach is O(H), where H is the height of the binary search tree. In the worst case, the recursion stack may go as deep as the height of the tree.