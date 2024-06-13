## Problem with using DFS for currency conversion problem 

Running DFS to convert currency rates between different currencies using a graph representation has the following issues

1. **Currency Conversion Logic**: The current DFS logic doesn't handle cycles or paths correctly, and the result might not be computed as intended.
2. **Rate Calculation**: The DFS implementation may accumulate incorrect rates if there are multiple paths.
3. **Error Handling**: There's a missing mechanism to handle cases where no path exists between two currencies.
4. **Bidirectional Rates**: Handling bidirectional rates should be careful to avoid incorrect conversions.
5. **Non-Existing Queries**: Queries for non-existing currencies should be managed.

### Solution:
We will modify the code to use Breadth-First Search (BFS) for more reliable path finding and better rate calculations. We'll also add a mechanism to handle non-existing queries and disconnected graphs.


### Explanation:

1. **Adjacency List**: The adjacency list is created to represent the graph.
2. **BFS Implementation**: A BFS is used to explore paths and calculate the conversion rate, which ensures we explore shortest paths and avoids cycles.
3. **Queue**: A queue stores the current node and the cumulative rate.
4. **Visited Map**: Keeps track of visited nodes to avoid re-processing and cycles.
5. **Handling Non-Existing Paths**: If no path exists, the function returns `-1` to indicate no conversion is possible.

This approach ensures a reliable calculation of conversion rates, handles disconnected graphs, and provides clear outputs for invalid queries.

## BFS vs Djkstras 

Breadth-First Search (BFS) and Dijkstra's algorithm are both used for traversing graphs, but they are suited for different types of problems due to their differing mechanisms and purposes. Here’s when to use each:

### Breadth-First Search (BFS)

**Definition:**
BFS is a graph traversal algorithm that explores all the nodes at the present depth level before moving on to the nodes at the next depth level. It uses a queue to keep track of the nodes to be explored next.

**When to Use BFS:**

1. **Unweighted Graphs:**
   - BFS is ideal for finding the shortest path in unweighted graphs because it explores all nodes at the current level before moving to the next level, ensuring that the first time it reaches a node, it is via the shortest path.

2. **Level-Order Traversal:**
   - When you need to explore nodes level by level, such as in tree traversal (e.g., finding all nodes at a certain depth).

3. **Connectivity:**
   - To check if a graph is connected, or to find all connected components in an undirected graph.

4. **Shortest Path in Unweighted Scenarios:**
   - When you need the shortest path in scenarios where all edges have the same weight (effectively unweighted).

**Example Use Cases:**
   - Finding the shortest path in a maze or grid where each move has an equal cost.
   - Social network analysis to find the degree of connection between people (e.g., "degrees of separation").

### Dijkstra's Algorithm

**Definition:**
Dijkstra's algorithm is a graph traversal algorithm used to find the shortest paths from a single source node to all other nodes in a weighted graph with non-negative weights. It uses a priority queue (often implemented as a min-heap) to select the next node with the smallest tentative distance.

**When to Use Dijkstra's:**

1. **Weighted Graphs:**
   - When you need to find the shortest path in graphs where edges have different weights.

2. **Non-Negative Weights:**
   - When the edge weights are non-negative. Dijkstra’s algorithm does not work correctly with negative weights as it may get stuck in an infinite loop of decreasing path costs.

3. **Single Source Shortest Path:**
   - When you need the shortest path from a single source to all other nodes in the graph, or to a specific node.

**Example Use Cases:**
   - Finding the shortest route on a road map where roads have different lengths or travel times.
   - Network routing protocols like OSPF (Open Shortest Path First) which find the shortest path for data packets in a network.

### Summary

- **BFS**: Use it in unweighted graphs for shortest path calculations, level-order traversal, connectivity checks, and finding connected components.
- **Dijkstra’s Algorithm**: Use it in weighted graphs with non-negative weights for single-source shortest path calculations.
