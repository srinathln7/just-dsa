# Level-Order Traversal 

Given a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

## Intuition:

The problem can be efficiently solved using Breadth-First Search (BFS). We traverse the tree level by level, starting from the root node. To implement BFS, we use a queue data structure, typically implemented as a double-ended queue (deque) in Go.

## Approach:

1. Begin with an empty result slice to store the level order traversal.
2. Initialize a deque with the root node.
3. While the deque is not empty, do the following:
   - Determine the number of nodes at the current level (`levelSize`) by getting the length of the deque.
   - Create an empty slice `level` to store the elements at the current level.
   - Iterate `levelSize` times:
     - Remove the front node from the deque.
     - Append its value to the `level` slice.
     - Enqueue its left and right child nodes if they exist.
   - Append the `level` slice to the `result` slice.
4. Return the `result` slice containing the level order traversal.

## Time Complexity:

The time complexity of this approach is O(n), where n is the number of nodes in the binary tree. This is because we visit each node once during the traversal.

## Space Complexity:

The space complexity is also O(n), where n is the number of nodes in the binary tree. This is due to the space required for the deque and the result slice, both of which can grow linearly with the number of nodes in the tree.


## Additional Remark

### Mistakes made

```
var result [][]int
var elements []int

// This will cause a runtime PANIC because we have only declared the 2D-array and no memory has been allocated yet
// Hence directly accessing the index will cause runtime panic
result[i] = append(result[i], elements...)

// Correct way
result = append(result, elements)


This is same as 1D-array
var result []int

// RUN TIME PANIC
result[i] = 1

// CORRECT
result = append(result, 1)


Alternatively this will work too

result := make([]int, 10)
result[0] = 1

```