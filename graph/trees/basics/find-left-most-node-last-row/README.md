# Find left mostnode in the last row
Given the root of a binary tree, find the leftmost value in the last row of the tree.

## Approach
The key idea is to traverse the tree and find the bottom-most left value. We can achieve this by recursively exploring the tree and comparing the depths of the left and right subtrees at each level. 

1. **Base Case**: If the current node is a leaf node (both left and right children are nil), return its value as it is the bottom-most left value.
2. **Recursion**: Recursively find the maximum depth of the left and right subtrees using the `maxDepth` function.
3. **Comparison**: Compare the maximum depths of the left and right subtrees. If the right subtree has a greater depth, recursively explore the right subtree. Otherwise, recursively explore the left subtree.
4. **Return**: Finally, return the value obtained from the recursive calls.

## Time Complexity
The time complexity of this approach is O(N), where N is the number of nodes in the binary tree. In the worst case, we may need to traverse the entire tree to find the bottom-most left value.

## Space Complexity
The space complexity of this approach is O(H), where H is the height of the binary tree. This space is used for the recursive call stack. In the worst case, the recursion stack may go as deep as the height of the tree.
