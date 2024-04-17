# Level Order Traversal of N-ary Tree

Given the root of an N-ary tree, return its level order traversal.

## Intuition:
Performing a breadth-first search (BFS) is an efficient approach to traverse an N-ary tree level by level. We start from the root node and process each level by enqueueing all child nodes of the current level nodes.

## Approach:
1. Initialize an empty 2D array `result` to store the level order traversal.
2. If the root node is `nil`, return the empty result array.
3. Create a queue data structure to perform BFS. Enqueue the root node into the queue.
4. Define a BFS function to process each level of the tree.
5. Within the BFS function, iterate while the queue is not empty:
   - Initialize an empty array `path` to store the node values at the current level.
   - Get the size of the queue (number of nodes at the current level).
   - Iterate `levelSize` times to process all nodes at the current level:
     - Dequeue the front node from the queue.
     - Append the value of the dequeued node to the `path`.
     - Enqueue all child nodes of the dequeued node into the queue.
   - Append the `path` array to the `result` array to store the current level values.
6. Call the BFS function to perform level order traversal.
7. Return the `result` array containing the level order traversal of the N-ary tree.

## Time Complexity:
The time complexity is O(N), where N is the total number of nodes in the N-ary tree. This is because we visit each node exactly once during the BFS traversal.

## Space Complexity:
The space complexity is O(N), where N is the total number of nodes in the N-ary tree. This is because the space required for the queue can grow up to O(N) in the worst case, where all nodes are at the last level. Additionally, the `result` array also requires O(N) space to store the level order traversal.