# Shortest Path in a Weighted Path

You are given a network of `n` nodes labeled from 1 to `n`. The network is represented by a list of directed edges `times`, where `times[i] = (ui, vi, wi)` denotes a directed edge from node `ui` to node `vi` with a signal travel time of `wi`. You need to determine the minimum time it takes for a signal to reach all nodes in the network when the signal starts from a given source node `k`.

## Intuition
This problem can be solved using Dijkstra's algorithm, which is a shortest path algorithm. We'll treat each node as a vertex in a graph, and each directed edge `(u, v, w)` as an edge from vertex `u` to vertex `v` with weight `w`, representing the time it takes for the signal to travel from node `u` to node `v`. We'll start from the source node `k` and find the shortest paths to all other nodes in the network.

## Approach
1. Implement Dijkstra's algorithm to find the shortest path from the source node `k` to all other nodes in the network.
2. Use a priority queue (min-heap) to efficiently select the next node with the minimum distance.
3. Initialize a map to store the minimum time it takes to reach each node from the source node `k`.
4. Iterate through all nodes using Dijkstra's algorithm, updating the minimum time for each node whenever a shorter path is found.
5. After processing all nodes, return the maximum time from the map, which represents the minimum time it takes for the signal to reach all nodes.
6. If any node is unreachable, return `-1`.

## Time Complexity
The time complexity of Dijkstra's algorithm with a priority queue implementation is O(E * log(V)), where V is the number of vertices (nodes) and E is the number of edges in the graph.

## Space Complexity
The space complexity is O(V + E), where V is the number of vertices (nodes) and E is the number of edges in the graph. This space is primarily used for storing the adjacency list and the priority queue.
