# K closest points
Given an array `arr` of integers sorted in non-decreasing order, an integer `k`, and an integer `x`, find the `k` closest elements to `x` in the array. The result should be sorted in ascending order if there is a tie.

## Intuition
To find the `k` closest elements to `x`, we can maintain a max heap of size `k` containing elements with the highest absolute difference from `x`. As we iterate through the array, we add each element along with its absolute difference from `x` to the max heap. If the size of the max heap exceeds `k`, we pop the element with the maximum absolute difference. Finally, we sort the elements in ascending order and return the result.

## Approach
1. Define a struct `Element` with fields `val` and `diff`, where `val` represents the value from the array and `diff` represents its absolute difference from `x`.
2. Implement a max heap `MaxHeap` to store elements based on their absolute differences.
3. Implement methods for the max heap:
   - `Len` to return the length of the heap.
   - `Less` to compare elements based on their absolute differences. If two elements have the same difference, compare their values.
   - `Swap` to swap elements in the heap.
   - `Top` to return the top element of the heap.
   - `Push` to push an element into the heap.
   - `Pop` to pop the top element from the heap.
4. Implement the `findClosestElements` function:
   - Initialize `result` as an empty slice to store the `k` closest elements.
   - Create a max heap `maxHeap`.
   - Iterate through the array `arr`:
     - Calculate the absolute difference between the current element and `x`.
     - Push the element along with its absolute difference into `maxHeap`.
     - If the size of `maxHeap` exceeds `k`, pop the top element.
   - Pop elements from `maxHeap` and store them in `result` in reverse order (to maintain ascending order).
   - Sort `result` in ascending order.
   - Return `result`.
5. Define a helper function `abs` to return the absolute value of an integer.

## Time Complexity
The time complexity of this approach is O(n log k), where n is the length of the array `arr`. The logarithmic factor comes from maintaining the max heap of size `k`.

## Space Complexity
The space complexity is O(k) to store the max heap.
