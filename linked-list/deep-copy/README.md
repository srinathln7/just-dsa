# Deep-Copy of LinkedList
You are given a linked list where each node has two pointers: `Next` and `Random`. The `Next` pointer points to the next node in the list, and the `Random` pointer can point to any node in the list or null. You need to create a deep copy of the original linked list.

## Intuition
To create a deep copy of the linked list with random pointers, we can traverse the original linked list while creating a copy of each node. As we traverse the list, we store a mapping from each node in the original list to its corresponding node in the copied list. This mapping allows us to properly set the `Next` and `Random` pointers for each node in the copied list.

## Approach
1. Initialize an empty map `visited` to store the mapping from original nodes to copied nodes.
2. Define a recursive DFS function `dfs` that takes a node `curr` as input and returns the corresponding copied node.
3. In the `dfs` function:
   - Check if `curr` exists in the `visited` map. If it does, return the corresponding copied node.
   - Create a copy node `copyNode` with the same value as `curr`.
   - Store the mapping from `curr` to `copyNode` in the `visited` map.
   - Recursively call `dfs` for the `Next` and `Random` pointers of `curr`, and assign the corresponding copied nodes to `copyNode.Next` and `copyNode.Random`.
   - Return `copyNode`.
4. Start the DFS traversal from the `head` of the original linked list.
5. Return the copied head node obtained from the DFS traversal.

## Time Complexity
- The time complexity of this approach is O(N), where N is the number of nodes in the linked list. This is because we visit each node exactly once during the DFS traversal.

## Space Complexity
- The space complexity of this approach is O(N), where N is the number of nodes in the linked list. This space is used for the `visited` map, which stores the mapping from original nodes to copied nodes. Additionally, the recursive stack space for DFS traversal can also be O(N) in the worst case if the linked list has a linear structure.
