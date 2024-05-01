# Walls and Gates
Given a 2D grid containing rooms with three possible values:
- `-1`: A wall or an obstacle.
- `0`: A gate.
- `INF`: Infinity, representing an empty room. Fill each empty room with the distance to its nearest gate. If it is impossible to reach a gate, it should be filled with `INF`.

## Intuition
To fill each empty room with the distance to its nearest gate, we can start from the gates and mark every adjacent empty room with the nearest distance values. Since the algorithm involves finding the shortest distance to its nearest gate, we can deploy a Breadth-First Search (BFS) algorithm.

## Approach
1. Initialize a queue to store the coordinates of gates.
2. Enqueue the coordinates of all gates into the queue.
3. Run a BFS on the grid:
   - For each cell `(r, c)` in the queue:
     - Increment the distance by 1.
     - Explore adjacent cells (up, down, left, right) and update their distances if they are empty rooms (`INF`).
     - Enqueue the coordinates of these adjacent empty rooms into the queue.
4. Repeat the BFS process until the queue is empty.
5. At the end of the BFS, the distances of all empty rooms will be updated.

## Time Complexity
- Let `m` be the number of rows and `n` be the number of columns in the grid.
- Each cell in the grid is visited once during the BFS process.
- Therefore, the time complexity is O(m * n).

## Space Complexity
- The space complexity is determined by the queue used in the BFS process and any additional data structures.
- In the worst case, the queue can store all the cells in the grid (if all cells are gates).
- Therefore, the space complexity is O(m * n) in the worst case.