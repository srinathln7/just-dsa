# Convert sorted array to BST:
Given a sorted array `nums` of integers, the task is to construct a balanced binary search tree (BST) from the array.

## Intuition:
- The solution utilizes a recursive approach to construct a balanced BST from the sorted array.
- The idea is to select the middle element of the array as the root of the tree, then recursively construct the left and right subtrees using the elements to the left and right of the middle element.

## Approach:
1. Define a function `sortedArrayToBST` that takes a sorted array `nums` as input and returns the root node of the constructed BST.
2. Implement the recursive construction logic:
   - Base cases:
     - If the length of the array `nums` is 0, return nil.
     - If the length of the array `nums` is 1, create a new tree node with the value of the only element in the array and return it.
   - Calculate the index of the middle element of the array as `mid`.
   - Create a new tree node `newRoot` with the value of `nums[mid]`.
   - Recursively construct the left subtree of `newRoot` by calling `sortedArrayToBST` with the subarray `nums[:mid]`.
   - Recursively construct the right subtree of `newRoot` by calling `sortedArrayToBST` with the subarray `nums[mid+1:]`.
   - Assign the left and right subtrees to `newRoot.Left` and `newRoot.Right` respectively.
3. Return the `newRoot` node as the root of the constructed BST.

## Time Complexity:
- The time complexity of this approach is O(n), where n is the number of elements in the sorted array. This is because each element of the array is visited once during the construction of the BST.

## Space Complexity:
- The space complexity is O(log n) on average and O(n) in the worst case due to the recursive function calls, where n is the number of elements in the sorted array. This space is used for the call stack during recursion.

