# Design Min-Heap

The goal is to implement a min-heap data structure that supports the following operations:
- Push: Inserts an element into the heap.
- Pop: Removes and returns the minimum element from the heap.

## Intuition:
A min-heap is a binary tree where the value of each node is less than or equal to the values of its children. When an element is pushed into the heap, it's placed at the bottom and then moved up (heapify up) until the heap property is restored. Similarly, when an element is popped, the last element is moved to the root, and then moved down (heapify down) until the heap property is restored.

## Approach:
1. **MinHeap Struct**: Defines a structure for the min-heap. It contains an array to store elements.
2. **Push Method**: Adds a new element to the heap and calls heapify up to maintain the heap property.
3. **Pop Method**: Removes the minimum element from the heap, replaces it with the last element, and calls heapify down to maintain the heap property.
4. **HeapifyUp Method**: Restores the heap property by moving a newly added element upwards until its parent's value is smaller.
5. **HeapifyDown Method**: Restores the heap property by moving a replaced root element downwards until its children's values are greater.

## Time Complexity:
- Push and Pop operations: O(log n) due to heapify operations.
- Construction of the heap: O(n) for building the heap from an array of size n.

## Space Complexity:
- O(n) for storing the elements in the array.


## Remark

The time complexity for constructing a heap from an array of size \( n \) is \( O(n) \). This process is also known as heapify or heap construction.

When constructing a heap, we start from the last non-leaf node and perform heapify down on each node until we reach the root. Since approximately half of the nodes in a binary tree are leaf nodes, we only need to heapify down the non-leaf nodes, which takes \( O(\log n) \) time for each node.

Considering there are \( n/2 \) non-leaf nodes in a heap, the total time complexity for heap construction is \( O(n \log n) \). However, due to the property of logarithms, we can simplify this to \( O(n) \) since \( \log n \) grows slowly compared to \( n \) as \( n \) becomes large. There is a very neat mathematical summation that is formed when we add all the work together, which simplifies to O(n).


Therefore, the time complexity for constructing a heap from an array of size \( n \) is \( O(n) \).
