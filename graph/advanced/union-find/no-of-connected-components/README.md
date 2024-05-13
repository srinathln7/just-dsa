# [No of connected components](https://neetcode.io/problems/count-connected-components)
Given a number of nodes and a list of edges representing connections between nodes, count the number of connected components in the graph.

## Intuition
The problem requires identifying connected components in an undirected graph. We can achieve this efficiently using the Union-Find (Disjoint Set Union) data structure.

## Approach
1. Implement the Union-Find data structure to efficiently manage the connected components.
2. Initialize each node with its parent pointing to itself (representing its own component).
3. Iterate through the edges and union the components of the nodes connected by each edge.
4. After processing all edges, the number of connected components is equal to the number of distinct roots in the Union-Find data structure.
5. Return the total number of nodes minus the count of connected components.

## Time Complexity
- Building the Union-Find data structure: \(O(n)\), where \(n\) is the number of nodes.
- Processing each edge and union operation: \(O(m\alpha(n))\), where \(m\) is the number of edges and \(\alpha(n)\) is the inverse Ackermann function (almost constant).
- Overall time complexity: \(O(n + m\alpha(n))\).

## Space Complexity
- Space required for the Union-Find data structure: \(O(n)\), where \(n\) is the number of nodes.
- Overall space complexity: \(O(n)\).
