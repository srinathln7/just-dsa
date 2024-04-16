# Lowest Common Ancestor

Given a binary search tree (BST), find the lowest common ancestor (LCA) of two given nodes in the BST.

## Intuition
In a BST, all values in the left subtree are less than the root's value, and all values in the right subtree are greater than the root's value. We can exploit this property to find the LCA efficiently.

## Approach
1. Start from the root.
2. Check if either `p`'s value is less than or equal to the root's value and `q`'s value is greater than or equal to the root's value, or vice versa. If so, the LCA is the root.
3. If both `p` and `q` are strictly less than the root's value, recurse on the left subtree.
4. If both `p` and `q` are strictly greater than the root's value, recurse on the right subtree.
5. Repeat until the LCA is found.

## Time Complexity
The time complexity of this approach is O(h), where h is the height of the BST. In the worst-case scenario, the height of the BST can be O(n) for skewed trees.

## Space Complexity
The space complexity is O(h), where h is the height of the BST, due to the recursion stack. In the worst-case scenario, the space complexity can be O(n) for skewed trees.