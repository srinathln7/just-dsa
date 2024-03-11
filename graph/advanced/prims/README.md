# Minimum Cost to Connect Points (aka) Minimum Spanning Tree

## Intuition
Given a set of points on a 2D plane, the problem is to find the minimum cost to connect all the points such that there is exactly one simple path between any two points. The cost of connecting two points is defined as the Manhattan distance between them.

## Approach
We can solve this problem using Prim's algorithm, which is a greedy algorithm used to find the minimum spanning tree (MST) of a connected, undirected graph. The steps involved are as follows:

1. Form the adjacency list representing the graph where each vertex is connected to all other vertices with edges weighted by the Manhattan distance between the corresponding points.
2. Initialize a priority queue (min-heap) to store the vertices with their associated costs.
3. Choose an arbitrary starting point and add it to the priority queue with a cost of 0.
4. While there are unvisited vertices:
   - Extract the vertex with the minimum cost from the priority queue.
   - Mark the extracted vertex as visited.
   - Add the cost of the extracted vertex to the total minimum cost.
   - Visit all unvisited neighbors of the extracted vertex and add them to the priority queue with their associated costs.
5. Return the total minimum cost calculated.

## Time Complexity
The time complexity of Prim's algorithm is O(E logV), where V is the number of vertices and E is the number of edges. 

## Space Complexity
The space complexity is O(E) for storing the adjacency list and the priority queue.
