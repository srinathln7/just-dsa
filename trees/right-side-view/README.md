# Binary Tree Right Side View

Given the root of a binary tree, return the right side view of the tree.

## Intuition:

To obtain the right side view of a binary tree, we perform a level-order traversal of the tree. However, instead of storing all nodes at each level, we only store the rightmost node of each level. This way, as we traverse the tree level by level, we ensure that only the rightmost node at each level is captured.

## Approach:

1. Initialize an empty result slice to store the right side view of the binary tree.
2. Initialize a deque with the root node.
3. While the deque is not empty, do the following:
   - Determine the number of nodes at the current level (`levelSize`) by getting the length of the deque.
   - Initialize a variable `rightSide` to store the value of the rightmost node at the current level.
   - Iterate `levelSize` times:
     - Remove the front node from the deque.
     - Update the `rightSide` variable with the value of the removed node.
     - Enqueue its left and right child nodes if they exist.
   - Append the `rightSide` value to the result slice.
4. Return the result slice containing the right side view of the binary tree.

## Time Complexity:

The time complexity of this approach is O(n), where n is the number of nodes in the binary tree. This is because we visit each node once during the traversal.

## Space Complexity:

The space complexity is also O(n), where n is the number of nodes in the binary tree. This is due to the space required for the deque and the result slice, both of which can grow linearly with the number of nodes in the tree.