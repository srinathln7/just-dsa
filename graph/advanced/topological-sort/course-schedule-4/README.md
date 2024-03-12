# Course Schedule 4

Given a list of prerequisites and a list of queries, determine whether each query course is a prerequisite of the corresponding course.

## Intution
The intuition behind the solution is to represent the prerequisites as a directed graph, where each course is a node and each prerequisite relationship is an edge. Then, for each query, we perform a depth-first search (DFS) from the source course to check if it can reach the destination course.

## Approach:
1. Form the adjacency list based on prerequisites, where each course is a key and its prerequisites are the values.
2. For each query, use DFS to check if the query course can be reached from the source course.
3. If the destination course is reachable from the source course, return true; otherwise, return false.

## Time Complexity:
- Building the adjacency list: O(N), where N is the number of prerequisites.
- DFS for each query: O(N + E), where N is the number of courses and E is the total number of edges in the graph.
- Overall time complexity: O(Q * (N + E)), where Q is the number of queries.

## Space Complexity:
- Adjacency list: O(N + E), where N is the number of courses and E is the total number of edges in the graph.
- Visited map: O(N), where N is the number of courses.
- Overall space complexity: O(N + E).