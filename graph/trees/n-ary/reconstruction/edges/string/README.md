```markdown
## Problem Statement
Given a list of parent-child relationships representing an N-ary tree, find the root node and print the entire tree.

## Intuition
To find the root node, we identify the node that appears as a parent but not as a child. Then, we perform a depth-first search (DFS) from the root node to construct the entire tree.

## Approach
1. Build an adjacency list to represent the parent-child relationships. Also, create a set to store child nodes.
2. Iterate through the edges and populate the adjacency list. Simultaneously, add child nodes to the set.
3. Find the root node by checking for a node that appears as a parent but not as a child.
4. Perform a DFS from the root node to construct the tree:
   - Start DFS from the root node.
   - Mark visited nodes to avoid cycles.
   - For each child of the current node, recursively explore unvisited children.
   - Append child nodes to the current node's children list.
5. Print the tree by recursively traversing and printing each node along with its level in the tree.

## Time Complexity
- Constructing the adjacency list and finding the root node: O(E), where E is the number of edges in the input.
- Constructing the tree using DFS: O(E), where E is the number of edges.
- Printing the tree: O(N), where N is the number of nodes in the tree.
Overall time complexity: O(E + N).

## Space Complexity
- Space required for the adjacency list and child set: O(E), where E is the number of edges.
- Space required for the visited map during DFS: O(N), where N is the number of nodes.
Overall space complexity: O(E + N).

