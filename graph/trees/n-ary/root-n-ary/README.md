# Find root of n-ary tree 
You are given all the nodes of an N-ary tree as an array of `Node` objects, where each node has a unique value. Return the root of the N-ary tree.

## Intuition:
In an N-ary tree, each node appears once as the parent and once as a child. So, if we XOR all the values of the nodes, then the root node's value will be left out as it appears only once.

## Approach:
1. Iterate through each node in the given `tree` array.
2. XOR the value of each node with `x`.
3. Iterate through the `tree` array again to find the node with value `x`. This will be the root node.

## Time Complexity: 
`O(N)` where `N` is the number of nodes in the tree.

## Space Complexity: 
`O(1)`