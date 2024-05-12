# [Find Eventual Safe Nodes](https://leetcode.com/problems/find-eventual-safe-states/description/)
- We are given a directed graph represented by an adjacency list (`graph`), where `graph[i]` contains the list of nodes adjacent to node `i`.
- A node is considered safe if all paths starting from that node lead to a terminal node (a node with no outgoing edges).

## Approach:
1. **DFS to Determine Safety:**
   - We traverse the graph using DFS from each node.
   - During DFS, if we encounter a cycle or a node that leads to an unsafe node, we mark the current node as unsafe.
   - Otherwise, if all child nodes are safe, we mark the current node as safe.

2. **Collect Safe Nodes:**
   - After DFS, we collect the indices of all nodes marked as safe and return them as the eventual safe nodes.

## Code Explanation:
- The `eventualSafeNodes` function initializes two arrays: `safe` to track the safety status of each node, and `visited` to mark visited nodes during DFS.
- It defines a nested DFS function `isSafe` that recursively checks if a node is safe by exploring its child nodes.
- If a node is already visited, the function returns its safety status from the `safe` array. Otherwise, it marks the node as visited and recursively checks its child nodes.
- If any child node is not safe, the current node is marked as unsafe. Otherwise, the current node is marked as safe.
- After DFS traversal, the function collects the indices of all safe nodes and returns them.

## Time Complexity:
- The time complexity of this solution is O(V + E), where V is the number of nodes and E is the number of edges in the graph. This is because we perform DFS traversal on each node once.

## Space Complexity:
- The space complexity is O(V), where V is the number of nodes. This is due to the space used by the `safe` and `visited` arrays.

