# All Paths from Source to Target
Given a directed acyclic graph (DAG) of n nodes labeled from 0 to n - 1, find all possible paths from node 0 to node n - 1 and return them in any order. The graph is given as follows: graph[i] is a list of all nodes you can visit from node i (i.e., there is a directed edge from node i to node graph[i][j]).

## Intuition
We can solve this problem using depth-first search (DFS). We start from node 0 and explore all possible paths until we reach node n - 1. During the DFS traversal, we keep track of the current path using a list. When we reach the destination node n - 1, we add the current path to the result. We use backtracking to explore all possible paths.

## Approach
1. Initialize an empty list `result` to store all possible paths.
2. Initialize an empty list `path` to store the current path during DFS traversal.
3. Initialize a boolean map `visited` to keep track of visited nodes.
4. Define a DFS function `dfs` that takes the current source node `src` and the destination node `dst`.
5. In the DFS function:
   - Mark the current node `src` as visited.
   - Append the current node `src` to the `path`.
   - If `src` is equal to `dst`, create a copy of the `path` and append it to `result`.
   - Iterate through all neighbors of `src` in the graph:
     - If the neighbor node is not visited, recursively call the DFS function with the neighbor node as the new source.
   - Backtrack by marking the current node `src` as unvisited and removing it from the `path`.
6. Call the DFS function with the initial source node 0 and destination node n - 1.
7. Return the `result`.

## Time Complexity
The time complexity of this approach is O(V + E), where V is the number of vertices (nodes) and E is the number of edges in the graph. This is because we visit each vertex and each edge exactly once during the DFS traversal.

## Space Complexity
The space complexity of this approach is O(V), where V is the number of vertices (nodes) in the graph. This space is used to store the current path during DFS traversal.
