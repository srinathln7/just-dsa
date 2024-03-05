# Shortest Path in Binary Matrix

Given an `n x n` binary matrix `grid`, where `grid[i][j]` is `0` or `1`, return the length of the shortest clear path in the matrix. If there is no clear path, return `-1`.

A clear path in a binary matrix is a path from the top-left cell `(0, 0)` to the bottom-right cell `(n - 1, n - 1)` such that:

- All the visited cells of the path are `0`.
- All the adjacent cells of the path are 8-directionally connected (i.e., they are different, and they share an edge or a corner).
- The length of a clear path is the number of visited cells of this path.

## Intuition

To find the shortest clear path in the binary matrix, we can use the breadth-first search (BFS) algorithm. Starting from the top-left cell, we explore the neighboring cells in all eight directions. We maintain a queue to perform BFS traversal and keep track of the path length. As soon as we reach the bottom-right cell, we return the length of the clear path. If there is no clear path, we return `-1`.

## Approach

1. Check if the grid is empty or if the top-left or bottom-right cell is blocked (contains a `1`). If so, return `-1`.
2. Initialize variables for the path length and a queue for BFS traversal.
3. Add the top-left cell `(0, 0)` to the queue and set the path length to `0`.
4. During BFS traversal, dequeue cells from the front of the queue and explore their neighbors in eight directions.
5. For each neighbor cell:
   - If it's out of bounds or already visited or blocked, continue to the next neighbor.
   - Otherwise, enqueue the neighbor cell, mark it as visited, and increment the path length.
6. Continue BFS traversal until reaching the bottom-right cell or until the queue becomes empty.
7. Return the length of the clear path if it exists; otherwise, return `-1`.

## Time Complexity

The time complexity of BFS traversal is O(V + E), where V is the number of vertices (cells) and E is the number of edges (connections between cells). In the worst case, all cells are traversed, so the time complexity is O(n^2).

## Space Complexity

The space complexity is O(n^2) since we use a queue for BFS traversal and additional space to store the grid and other variables.

## Why BFS will always yield the shortest path?
BFS (Breadth-First Search) ensures the length of the shortest path in a graph or grid by exploring nodes in layers, starting from the source node. Here's how BFS ensures that the length of the path it finds is the shortest:

1. **Layer-by-Layer Exploration**: BFS explores nodes in layers, starting from the source node and moving outward. It explores all nodes at a particular distance (or depth) from the source before moving to nodes at the next distance. This ensures that if there are multiple paths from the source to any other node, BFS will always discover the shortest path first because it explores nodes closer to the source before exploring nodes farther away.

2. **Queue-based Approach**: BFS uses a queue data structure to maintain the order in which nodes are visited. Nodes are added to the queue as they are discovered, and they are processed in the order they were added. Since BFS explores nodes in the order they are added to the queue, it naturally explores nodes at each depth level before moving to deeper levels. This queue-based approach ensures that BFS explores the shortest path first before exploring longer paths.

3. **Marking Visited Nodes**: BFS marks visited nodes to avoid revisiting them. Once a node is visited, it is marked as visited, and BFS does not revisit it. This prevents BFS from exploring cycles in the graph and ensures that each node is visited only once. By marking visited nodes, BFS avoids exploring longer paths that may result from revisiting nodes.

4. **Termination Condition**: BFS terminates when it reaches the target node (e.g., the destination node in a shortest path problem). When BFS reaches the target node, it has explored all nodes at the current depth level, and the length of the shortest path to the target node is determined. Since BFS explores nodes in layers and terminates when it finds the target node, it guarantees that the path it finds is the shortest.

Overall, BFS explores nodes layer by layer, maintains the order of exploration using a queue, marks visited nodes to avoid revisiting them, and terminates when it finds the target node. These properties ensure that BFS always finds the shortest path in a graph or grid.