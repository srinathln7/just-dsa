# Kth smallest element in BST

Given the root of a binary search tree (BST) and an integer k, the task is to find the kth smallest value (1-indexed) among all the values of the nodes in the tree.

**Intuition:**
To find the kth smallest value in a BST, we can perform an inorder traversal of the tree and keep track of the count of visited nodes. When the count reaches k, we return the value of the current node.

**Approach:**
1. Define the function `kthSmallest` which takes the root of the BST and the integer k as input.
2. Initialize a variable `result` to store the kth smallest value.
3. Call a helper function `inorderTraversalTillK` to perform an inorder traversal of the BST until the kth smallest element is found.
4. In the `inorderTraversalTillK` function, recursively traverse the left subtree, visit the current node, decrement the count `k`, and recursively traverse the right subtree.
5. When `k` becomes zero, update the `result` with the value of the current node.
6. Finally, return the `result`.

**Time Complexity:**
The time complexity of this approach depends on the height of the BST. In the worst-case scenario, where the BST is highly unbalanced, the time complexity is O(N), where N is the number of nodes in the tree. However, on average, the time complexity is O(log N) due to the balanced nature of BSTs.

**Space Complexity:**
The space complexity of this approach is O(H), where H is the height of the BST, due to the recursive calls on the stack. In the worst-case scenario, the height of the tree could be O(N) if the tree is highly unbalanced, leading to a space complexity of O(N).
