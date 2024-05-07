# Convert BST to GST
Given the root of a Binary Search Tree (BST), convert it to a Greater Tree such that every key of the original BST is changed to the original key plus the sum of all keys greater than the original key in BST.

## Intuition
The problem can be solved using a reverse inorder traversal of the BST. By traversing the BST in reverse inorder (Right -> Root -> Left), we visit nodes in descending order of their values. During this traversal, we maintain a running sum of all node values greater than the current node. We update each node's value by adding this running sum to the node's original value.

## Approach
1. Implement a reverse inorder traversal function `revInorderTraversal` that takes a node pointer and a sum pointer as arguments.
2. In the `revInorderTraversal` function:
   - If the current node is nil, return.
   - Recursively call the function on the right subtree.
   - Update the current node's value by adding the running sum to its original value. Update the running sum by adding the current node's value.
   - Recursively call the function on the left subtree.
3. Initialize a running sum variable `sum` to 0.
4. Call the `revInorderTraversal` function with the root node and the address of the `sum` variable.
5. Return the modified root node.

## Time Complexity
The time complexity of the reverse inorder traversal is O(N), where N is the number of nodes in the BST.

## Space Complexity
The space complexity is O(H), where H is the height of the BST, due to the recursive stack space. In the worst-case scenario, the BST could be skewed, leading to O(N) space complexity.