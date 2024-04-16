# Union-Find Data Structure

## Intuition:
The Union-Find data structure, also known as Disjoint Set Union (DSU), is a data structure that efficiently manages a collection of disjoint sets (aka) Forest of trees. It provides two main operations: finding the representative of the set to which an element belongs (`find`), and merging two sets (`union`). This implementation of Union-Find is particularly useful for detecting cycles in undirected graphs.

## Approach:
1. **Initialization**: When initializing a new Union-Find data structure, each node is its own parent, and the rank (height) of each node is set to zero. This means that initially, each node is in its own disjoint set.
2. **Find Operation**: The `find` operation returns the ancestor (parent or grandparent) of a given node. It utilizes path compression, a technique to optimize the `find` operation, by making each node directly point to its ancestor, reducing the time complexity of subsequent `find` operations.
3. **Union Operation**: The `union` operation combines two disjoint sets represented by two given nodes. It first finds the ancestors of the nodes using the `find` operation. If the ancestors are the same, it means the nodes are already in the same set, indicating a cycle in the graph. Otherwise, it merges the two sets by setting one ancestor as the parent of the other, taking into account the rank of each set to maintain the balance of the resulting tree structure.

## Time Complexity:
- The time complexity of the `find` operation is O(log n) on average due to path compression, where n is the number of nodes.
- The time complexity of the `union` operation is O(log n) on average due to rank-based union and path compression.

## Space Complexity:
- The space complexity is O(n), where n is the number of nodes, to store the parent and rank for each node in the Union-Find data structure.

