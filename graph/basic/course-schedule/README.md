# Course Schedule 
You are given a total of `numCourses` courses labeled from 0 to `numCourses - 1`. Each course has prerequisites represented as pairs in the `prerequisites` array, where `prerequisites[i] = [ai, bi]` indicates that you must take course `bi` first if you want to take course `ai`. Determine if it is possible to finish all courses.

## Intuition:
The problem involves forming a graph using adjacency list where each course is represented as a node and prerequisites are represented as directed edges between nodes. We need to check if there's any cycle in the graph. If there's a cycle, it's impossible to finish all courses. Otherwise, it's possible.

## Approach

1. Build a graph using an adjacency list representation from the given prerequisites.
2. Implement a depth-first search (DFS) function to detect cycles in the graph.
3. In DFS, maintain two maps: `visited` to keep track of already visited vertices and `visiting` to keep track of vertices currently being visited to detect cycles.
4. Traverse each course and check if there's a cycle using DFS. If a cycle is detected, return false. Otherwise, return true.

## Time Complexity

- Building the graph: O(E), where E is the number of edges (prerequisites).
- DFS traversal: O(V + E), where V is the number of vertices (courses) and E is the number of edges (prerequisites).
- Overall time complexity: O(V + E), where V is the number of vertices and E is the number of edges.

## Space Complexity

- Space for the graph: O(V + E), where V is the number of vertices and E is the number of edges.
- Space for the `visited` and `visiting` maps: O(V), where V is the number of vertices.
- Overall space complexity: O(V + E), where V is the number of vertices and E is the number of edges.
