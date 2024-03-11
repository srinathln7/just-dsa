# Min. cost to connect all points (aka) Minimum Spanning Tree
You are given an array `points` representing integer coordinates of some points on a 2D-plane, where `points[i] = [xi, yi]`. The cost of connecting two points `[xi, yi]` and `[xj, yj]` is the Manhattan distance between them: `|xi - xj| + |yi - yj|`, where `|val|` denotes the absolute value of `val`. Return the minimum cost to make all points connected. All points are connected if there is exactly one simple path between any two points.

## Intuition
To find the minimum cost to connect all points, we can use Kruskal's algorithm, which finds the minimum spanning tree (MST) of the graph. In Kruskal's algorithm, we sort the edges by weight and add them to the MST one by one, ensuring that no cycle is formed.

## Approach
1. Create a struct `UnionFind` to implement the union-find data structure.
2. Implement methods to initialize the data structure, find the parent of a node, and perform union operation.
3. Define a helper function `abs` to calculate the absolute value of an integer.
4. Implement the `minCostConnectPoints` function:
   - Initialize an empty slice `edges` to store the edges of the graph.
   - Iterate over each pair of points to calculate the Manhattan distance between them and add the edge to the `edges` slice.
   - Sort the `edges` slice in ascending order based on the distance.
   - Initialize a variable `mincost` to store the total minimum cost.
   - Create a `UnionFind` data structure.
   - Iterate through the sorted edges and perform the union operation if it does not create a cycle.
   - Accumulate the cost of each added edge.
5. Return the `mincost`.

## Time Complexity
The time complexity of this approach is dominated by the sorting of edges, which takes O(E log V).

## Space Complexity
The space complexity is O(V), where V is the number of vertices, for storing the parent and rank arrays in the union-find data structure.
To form the edges O(E) where `E` is the number of points. Since `E <= v^2`, the space complexity will be O(E). 