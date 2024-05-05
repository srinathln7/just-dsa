## Problem Statement
Given a list of sequences representing a tree structure, find the root node and print the entire tree.

## Intuition
To find the root node, we first extract the edges from the sequences. Then, we construct an adjacency list to represent the parent-child relationships. Finally, we perform a depth-first search (DFS) from the root node to construct the entire tree.

## Approach
1. Define a `TreeNode` struct to represent a node in the tree.
2. Implement a function `getEdges` to extract the edges from the sequences.
   - Create a set to store unique edges.
   - Iterate through each sequence and extract edges between adjacent elements.
   - Add each edge to the set.
3. Implement a function `findRoot` to find the root node of the tree.
   - Construct an adjacency list and a set to store child nodes.
   - Iterate through the edges and populate the adjacency list. Simultaneously, add child nodes to the set.
   - Find the root node by checking for a node that appears as a parent but not as a child.
   - Perform a DFS from the root node to construct the tree:
     - Start DFS from the root node.
     - Mark visited nodes to avoid cycles.
     - For each child of the current node, recursively explore unvisited children.
     - Append child nodes to the current node's children list.
4. Implement a function `printTree` to print the tree in a human-readable format.
   - Recursively traverse the tree and print each node along with its level in the tree.

## Time Complexity
- Extracting edges: O(N), where N is the total number of elements in all sequences.
- Finding the root node: O(N), where N is the number of edges.
- Constructing the tree using DFS: O(N), where N is the number of edges.
- Printing the tree: O(N), where N is the number of nodes in the tree.
Overall time complexity: O(N).

## Space Complexity
- Space required for the edge set: O(N), where N is the number of edges.
- Space required for the adjacency list and child set: O(N), where N is the number of edges.
- Space required for the visited map during DFS: O(N), where N is the number of nodes.
Overall space complexity: O(N).

