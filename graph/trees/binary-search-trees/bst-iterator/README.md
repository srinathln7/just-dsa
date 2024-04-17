# BST Iterator

Implement the BSTIterator class, which represents an iterator over the in-order traversal of a binary search tree (BST). The iterator should initialize with the root of the BST and support two operations: `Next()` and `HasNext()`.

## Intuition
The key idea is to use a stack to traverse the BST in an iterative manner. By pushing all values to the left of the root onto the stack initially, we ensure that the iterator returns elements in ascending order.

## Approach
1. Initialize a stack to hold tree nodes and push all left child nodes of the root onto the stack.
2. Implement the `next()` method:
   - Pop the top element from the stack, which represents the next smallest node.
   - If the popped node has a right child, push all the **left descendants of its right child** onto the stack.
3. Implement `HasNext()` to check if the stack is empty.

## Time Complexity
- `Next()`: O(1) amortized time complexity on average. It pops and pushes elements onto the stack in constant time.
- `HasNext()`: O(1) time complexity.

## Space Complexity
- O(h), where `h=log(n)` is the height of the tree. The stack may hold up to h nodes in the worst case, corresponding to the maximum depth of the tree.


