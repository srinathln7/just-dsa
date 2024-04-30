# Fenwick Tree (aka) Binary Index Trees
Implement a Binary Index Tree (BIT), also known as a Fenwick tree, in Go. The goal is to support efficient updates and prefix sum queries on an array.

## Intuition
A Binary Index Tree (BIT), or Fenwick tree, is a data structure that efficiently maintains the prefix sum of an array while supporting updates to individual elements. It achieves this by storing cumulative sums at certain indices in a binary tree-like structure.

## Approach
1. **Initialization**: Initialize a BIT with an array of zeros with the same length as the input array.
2. **Update Operation**: To update an element at index `idx`, traverse the BIT tree and update all nodes whose ranges cover `idx`. Update each node by adding `delta` to its current value.
3. **Prefix Sum Query**: To compute the prefix sum up to index `idx`, traverse the BIT tree and accumulate the sum of all nodes from `idx` up to the root. Return the accumulated sum as the prefix sum.

## Time Complexity:
   - **Initialization**: O(n)
   - **Update Operation**: O(log n)
   - **Prefix Sum Query**: O(log n)

## Space Complexity: O(n)
