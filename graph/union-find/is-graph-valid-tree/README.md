# Make Graph a Valid Tree

You are given an integer `n` representing the number of nodes in a graph and a list of undirected edges `edges`, where `edges[i] = [a, b]` indicates that there is an edge between nodes `a` and `b`. You need to determine whether the given graph is a valid tree.

## Intuition

A valid tree must satisfy two conditions:
1. It must contain `n - 1` edges, where `n` is the number of nodes.
2. It must not contain any cycles.

To check for cycles and ensure connectivity, we can use the Union-Find data structure. We will iterate through all the edges and perform a union operation on their endpoints. If we encounter a redundant connection (i.e., the two endpoints are already in the same connected component), then there exists a cycle, and the graph cannot be a valid tree.

## Approach

1. Define a `UnionFind` data structure with two slices: `parent` to store the parent of each node and `rank` to store the rank of each node.
2. Implement methods `newUnionFind`, `find`, and `union` for the `UnionFind` struct.
3. Initialize a `UnionFind` instance with `n` nodes.
4. Iterate through all edges in `edges`:
   - Extract the endpoints `n1` and `n2` from the current edge.
   - Perform the union operation on `n1` and `n2`.
   - If the union operation returns `false`, it means `n1` and `n2` are already in the same connected component, indicating the presence of a cycle. Return `false`.
5. If the iteration completes without encountering any cycles and the number of edges equals `n - 1`, return `true`; otherwise, return `false`.

## Time and Space Complexity

Let `n` be the number of nodes and `m` be the number of edges.

The time complexity of this approach is O(n + m) for performing the union operations and iterating through the edges.

The space complexity is O(n) for storing the parent and rank arrays in the Union-Find data structure.