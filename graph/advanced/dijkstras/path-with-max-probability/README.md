# Path with Max. Probability
You are given a network with `n` nodes labeled from 0 to `n-1`. Each edge `(u, v)` in the network has an associated probability `prob[i]`, which represents the probability of successfully traveling from node `u` to node `v`. You need to find the maximum probability of successfully reaching the destination node `end_node` starting from the source node `start_node`.

## Intuition
This problem can be solved using a modified version of Dijkstra's algorithm. Instead of minimizing the cost, we need to maximize the probability of successfully reaching the destination node. We can achieve this by using a max-heap to prioritize paths with higher probabilities.

## Approach
1. Create an adjacency list representing the undirected graph with probabilities associated with each edge.
2. Initialize a max-heap with the start node and maximum probability of 1.0.
3. Perform Dijkstra's algorithm using the max-heap:
   - Pop the node with the maximum probability from the heap.
   - Mark the node as visited and store the maximum probability associated with reaching that node.
   - Update the probabilities of its neighboring nodes by multiplying them with the probability of the current node and push them into the max-heap.
4. After processing all nodes, return the maximum probability of successfully reaching the destination node.
5. If the destination node is unreachable from the start node, return 0.

## Time Complexity
The time complexity of Dijkstra's algorithm with a max-heap implementation is O((V + E) * log(V)), where V is the number of vertices (nodes) and E is the number of edges in the graph.

## Space Complexity
The space complexity is O(V + E), where V is the number of vertices (nodes) and E is the number of edges in the graph. This space is primarily used for storing the adjacency list, the max-heap, and the visited nodes.