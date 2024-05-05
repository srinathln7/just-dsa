# Flatten Tree to LinkedList 

Given the root of a binary tree, flatten the tree into a "linked list" where the right child pointer points to the next node in the list, and the left child pointer is always null. The "linked list" should be in the same order as a pre-order traversal of the binary tree.

## Approach

The key idea behind the solution is to perform a pre-order traversal of the binary tree and recursively flatten the left and right subtrees. After flattening the left subtree, the right pointer of the root node is modified to point to the flattened left subtree, and the left pointer is set to nil. Then, the function finds the rightmost node of the flattened left subtree and attaches the original right subtree to it. This process is performed recursively until the entire binary tree is flattened.

The steps involved in the approach are as follows:

1. Base cases:
   - If the root is nil, return.
   - If the root has no children (both left and right pointers are nil), return.

2. Temporarily store the right subtree of the current node.

3. Recursively flatten the left subtree.

4. Recursively flatten the right subtree.

5. Modify the root's right pointer to point to the flattened left subtree.

6. Set the root's left pointer to nil.

7. Traverse to the rightmost node of the flattened left subtree.

8. Attach the original right subtree to the rightmost node of the flattened left subtree.

## Time and Space Complexity Analysis

The time complexity of this approach is O(n), where n is the number of nodes in the binary tree. This is because each node is visited once during the pre-order traversal.

The space complexity is also O(n), considering the recursive call stack, where n is the height of the binary tree in the worst case scenario.