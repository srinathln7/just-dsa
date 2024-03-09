# Minimum Network Delay Time

You are given a network with `n` nodes labeled from 1 to `n`. Each directed edge `(ui, vi, wi)` represents a signal traveling from node `ui` to node `vi` with a time delay of `wi`. You need to find the minimum time it takes for a signal to reach all nodes in the network when starting from a given source node `k`. If it is impossible for all nodes to receive the signal, return -1.

## Intuition
To find the minimum time for the signal to reach all nodes, we can use Dijkstra's algorithm. We'll treat each node as a vertex in a graph and each directed edge as an edge with a weight representing the time delay. Starting from the source node `k`, we'll calculate the shortest paths to all other nodes in the network.

## Approach
1. Implement Dijkstra's algorithm to find the shortest paths from the source node `k` to all other nodes.
2. Use a priority queue (min-heap) to efficiently select the next node with the minimum distance.
3. Initialize a map to store the minimum time it takes to reach each node from the source node `k`.
4. Iterate through all nodes using Dijkstra's algorithm, updating the minimum time for each node whenever a shorter path is found.
5. After processing all nodes, return the maximum time from the map, which represents the minimum time it takes for the signal to reach all nodes.
6. If any node is unreachable, return -1.

## Time Complexity
The time complexity of Dijkstra's algorithm with a priority queue implementation is O((V + E) * log(V)), where V is the number of vertices (nodes) and E is the number of edges in the graph.

## Space Complexity
The space complexity is O(V + E), where V is the number of vertices (nodes) and E is the number of edges in the graph. This space is primarily used for storing the adjacency list, the priority queue, and the shortest path map.