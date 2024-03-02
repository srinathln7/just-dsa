# Clone a Graph

You are given a graph represented as an adjacency list. Clone this graph.

## Intuition:
To clone the graph, we need to create a deep copy of each node and its neighbors while preserving the structure and connections of the original graph.

## Approach:
We can use Depth-First Search (DFS) or Breadth-First Search (BFS) to traverse the original graph and create its clone.

## Depth-First Search (DFS) Approach:
1. Create a map to store the mapping between original nodes and their clones.
2. Start DFS traversal from the given node.
3. During traversal, for each node encountered:
   - If the node has already been visited, return its corresponding clone from the map.
   - If not visited, create a clone of the current node, mark it as visited, and recursively clone its neighbors.
4. Return the clone of the given node, which represents the cloned graph.

## Breadth-First Search (BFS) Approach:
1. Start BFS traversal from the given node.
2. Maintain a visit map to keep track of the mapping between original nodes and their clones.
3. Use a queue to perform level-order traversal of the original graph.
4. During traversal, for each node and its neighbors:
   - If the neighbor has not been visited, create a clone of it, mark it as visited, link it to the clone of the current node, and enqueue it for further exploration.
   - If visited, link the clone of the neighbor to the clone of the current node.
5. Return the clone of the given node, which represents the cloned graph.

### Time Complexity:
Both DFS and BFS approaches visit each node and its neighbors exactly once, resulting in a time complexity of O(V + E), where V is the number of vertices (nodes) and E is the number of edges in the graph.

### Space Complexity:
- **DFS Approach:** Requires space for the recursion stack and the visit map. The space complexity is O(V), where V is the number of vertices.
- **BFS Approach:** Requires space for the queue and the visit map. The space complexity is also O(V), where V is the number of vertices.

Overall, both approaches have similar time and space complexities, and the choice between them depends on factors such as personal preference and specific requirements of the problem.