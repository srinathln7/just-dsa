# Populating Next Right Pointers in Each Node
Given a binary tree, connect each node to its next right node.

## Intuition
The problem can be solved using a level order traversal (BFS) of the binary tree. At each level of the tree, we connect each node to its immediate next right node.

## Approach
1. Implement a function `connect` that takes the root node of the binary tree as input.
2. Initialize a queue slice with the root node and perform the following steps:
   - While the queue is not empty:
     - Get the size of the current level (levelSize) by calculating the length of the queue.
     - Connect each node in the current level to its immediate next right node by iterating through the first `levelSize - 1` nodes in the queue and setting their `Next` pointers to the next node in the queue.
     - Traverse the current level by popping nodes from the front of the queue and enqueuing their left and right children.
3. Repeat the above steps until all levels of the binary tree are processed.
4. Return the modified root node with updated `Next` pointers.

## Time Complexity
The time complexity of the level order traversal (BFS) is O(N), where N is the number of nodes in the binary tree.

## Space Complexity
The space complexity is O(W), where W is the maximum width of the binary tree (number of nodes in the level with the most nodes). In the worst case, the binary tree could be completely unbalanced, leading to O(N) space complexity.