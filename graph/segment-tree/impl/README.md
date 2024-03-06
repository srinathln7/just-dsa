# Segment Tree

The problem involves implementing a Segment Tree data structure that supports range queries and point updates efficiently.

## Intuition

Segment trees are versatile data structures that facilitate efficient querying and updating operations on a range of elements in an array. The key intuition behind segment trees is to divide the array into smaller segments recursively until each segment represents a single element. Each node in the segment tree stores information about the range it represents, such as the sum, maximum, minimum, etc.

## Approach

1. **Construction**: To build the segment tree, we recursively divide the array into segments until we reach the base case (i.e., a single element). At each step, we compute the value stored in the current node based on the values of its children nodes.
   
2. **Point Update**: When updating a value at a specific index in the array, we traverse the segment tree to locate the corresponding leaf node representing the index. We update the value at this leaf node and propagate the changes upwards by updating the parent nodes accordingly.
   
3. **Range Query**: For range queries, we start from the root of the segment tree and traverse down to the leaf nodes that cover the given range. At each step, we determine whether the current segment fully lies within the range, partially overlaps with it, or lies outside it. We compute the query result accordingly and aggregate the values obtained from different segments.

## Time Complexity

- **Construction**: O(N), where N is the number of elements in the input array.
- **Point Update**: O(log N), where N is the number of elements in the input array.
- **Range Query**: O(log N), where N is the number of elements in the input array.

## Space Complexity

O(N), where N is the number of elements in the input array, due to the space required to store the segment tree.

## Remark
This implementation provides a basic understanding of segment trees and their operations. It can be extended to support various types of range queries (e.g., minimum, maximum, sum) and updates (e.g., addition, subtraction).

Segment trees are needed to efficiently perform range queries on an array or a list of elements. To understand why they are needed, let's consider a simple example:

Suppose you have an array representing the daily temperatures of a city for a month. You want to answer queries like finding the maximum temperature over a given range of days or finding the average temperature over a specific period. 

Without any pre-processing or data structure, you would need to iterate through the array and calculate the required result for each query. This approach would be inefficient, especially for large arrays or frequent queries.

Segment trees provide a solution to this problem by pre-computing and organizing the information in a tree-like structure. Each node of the segment tree represents a range of elements from the original array. 

Here's how segment trees help:

1. **Pre-computation**: During construction, the segment tree stores pre-computed values for each node based on the range it represents. For example, if the query is to find the maximum temperature over a range, each node in the segment tree will store the maximum temperature within its corresponding range.

2. **Efficient Querying**: When a query is made, the segment tree efficiently traverses through the tree, identifying the relevant nodes that overlap with the query range. By combining the information stored in these nodes, it can quickly compute the result for the query.

3. **Reduced Time Complexity**: With segment trees, the time complexity of range queries is significantly reduced from O(n) (linear search through the array) to O(log n), where n is the size of the input array. This is because segment trees exploit the inherent structure of the data to perform queries more efficiently.

In summary, segment trees are needed to efficiently handle range queries on arrays or lists of data, providing faster responses to common queries and reducing the overall computational complexity of operations like finding maximum, minimum, sum, average, etc., over specified ranges.