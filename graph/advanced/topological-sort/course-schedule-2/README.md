# Course Schedule 2
Given the number of courses `numCourses` and an array `prerequisites` where `prerequisites[i] = [ai, bi]` indicates that course `ai` must be taken before course `bi`, return a valid order in which the courses can be taken. 

## Intuition
The key idea is to form a directed graph representing the prerequisite relationships between courses and then perform a depth-first search (DFS) to detect cycles in the graph. If the graph is acyclic, we can perform a topological sort to determine the order of courses. 

## Approach
1. **Forming the Adjacency List:** Create an adjacency list representing the graph, where each course is a node and each prerequisite relationship is an edge.
   
2. **Detecting Cycles with DFS:** Perform a DFS to detect cycles in the graph. Use two maps, `visited` and `visiting`, to keep track of visited and visiting vertices during each iteration. If a cycle is detected, return `nil`.
   
3. **Topological Sort:** Once the graph is confirmed to be acyclic, perform a topological sort to determine the order of courses. Append each course to the `topo` slice as it is visited in the DFS traversal.

4. **Returning the Result:** If the graph is acyclic, return the topological order of courses. No need to reverse the order since the nature of the problem already gives the order in reverse.

## Time Complexity
The time complexity of this approach is O(V + E), where V is the number of courses and E is the number of prerequisite relationships.

## Space Complexity
The space complexity is O(V + E), where V is the number of courses and E is the number of prerequisite relationships. This is due to the space used by the adjacency list and auxiliary data structures for DFS.


