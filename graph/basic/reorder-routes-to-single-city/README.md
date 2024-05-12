# [Reorder Routes](https://leetcode.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/description/)
You are given a network of cities represented by a directed graph. The roads in this network can only be traversed in one direction due to their narrowness. Your task is to find the minimum number of edges that need to be reversed so that each city can reach city 0.

## Intuition
Since the roads can only be one way, for all cities to reach city 0, there should be no outgoing paths (edges) from city 0. We can achieve this by running a depth-first search (DFS) starting from city 0 and counting the total number of outgoing paths. This count represents the number of directions that need to be reversed.

## Approach
1. **Adjacency List Construction**:
   - Construct an adjacency list to represent the directed graph. For each edge `(src, dst)`:
     - Add `dst` to the list of neighbors of `src` to signify an outgoing edge from `src` to `dst`.
     - Add `-src` to the list of neighbors of `dst` to signify an incoming edge from `dst` to `src`.

2. **DFS Traversal**:
   - Initialize a count variable to keep track of the number of outgoing edges that need to be reversed.
   - Initialize a visited map to keep track of visited cities during DFS traversal.
   - Define a DFS function to traverse the graph starting from city 0:
     - If the current city is already visited, return.
     - Mark the current city as visited.
     - For each neighbor of the current city:
       - If the neighbor is not visited:
         - If the neighbor represents a forward direction (positive value), increment the count and recursively call DFS on the neighbor.
         - If the neighbor represents a backward direction (negative value), recursively call DFS on the absolute value of the neighbor.
3. **Return Result**:
   - After DFS traversal, return the count, which represents the minimum number of edges that need to be reversed.

## Time Complexity
The time complexity of this approach is O(V + E), where V is the number of vertices (cities) and E is the number of edges (connections). This is because we perform a DFS traversal of the graph, visiting each vertex and each edge once.

## Space Complexity
The space complexity of this approach is O(V + E), where V is the space used for the adjacency list and visited map, and E is the space used for the recursive call stack during DFS traversal.