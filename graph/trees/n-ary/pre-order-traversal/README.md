# Pre-order traversal of N-ary tree 

Given the definition of a node in an N-ary tree, represented by the struct `Node`, consisting of a value and a slice of children nodes, implement a function `preorder` to return the pre-order traversal of the tree.


## Intuition:

Pre-order traversal visits the root node first, then recursively visits all the children nodes in left-to-right order.

## Approach:

1. Initialize an empty slice `result` to store the pre-order traversal.
2. If the root node is nil, return an empty `result`.
3. Traverse the root node and its children in pre-order recursively.
4. Append the value of the root node to the `result`.
5. Recursively traverse each child node and append its value to the `result`.
6. Return the `result` slice.

## Time Complexity: \(O(N)\), where \(N\) is the total number of nodes in the N-ary tree.

## Space Complexity: \(O(N)\), where \(N\) is the total number of nodes in the N-ary tree. This space is used to store the result list.