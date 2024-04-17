# Find Maximum Depth of N-ary Tree

You are given all the nodes of an N-ary tree as an array of `Node` objects, where each node has a unique value. Return the maximum depth of the N-ary tree.

## Intuition:
To find the maximum depth of an N-ary tree, we can utilize recursion. We start from the root node and traverse each child node recursively. At each level, we calculate the depth by adding 1 to the depth of its children.

## Approach:
1. Define a recursive function `maxDepth` to calculate the maximum depth of the tree.
2. If the root node is `nil`, return 0 as the depth.
3. Initialize the `depth` variable to 1, representing the depth at the current root level.
4. Iterate through each child of the root node and recursively call the `maxDepth` function to find the depth of each child node.
5. Update the `depth` variable by taking the maximum depth among all child nodes plus 1 (for the current root node).
6. Return the calculated `depth` as the maximum depth of the tree.

## Time Complexity: 
The time complexity is O(N), where N is the total number of nodes in the N-ary tree. This is because we visit each node exactly once during the traversal.

## Space Complexity: 
The space complexity is O(H), where H is the height of the tree. In the worst case, where the tree is highly unbalanced, the space complexity approaches O(N), but in the average case (balanced tree), it is O(log N).