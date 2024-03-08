# Median from a data stream

Design a data structure to find the median of a stream of integers.

## Intuition
To efficiently find the median of a stream of integers, we can use two heaps: a max-heap to store the lower half of the elements and a min-heap to store the upper half of the elements. By maintaining these two heaps, we can efficiently find the median.

## Approach
1. Implement two heaps: a max-heap and a min-heap.
2. When adding a new number to the stream:
   - If the max-heap is empty or the number is less than the maximum element in the max-heap, push the number to the max-heap.
   - Otherwise, push the number to the min-heap.
3. Rebalance the heaps if necessary to ensure the size difference between them is at most 1.
4. To find the median:
   - If the sizes of the heaps are equal, return the average of the top elements of both heaps.
   - If the max-heap has more elements, return the top element of the max-heap.
   - If the min-heap has more elements, return the top element of the min-heap.

## Time Complexity
- Adding a number: O(log n), where n is the total number of elements inserted.
- Finding the median: O(1).

## Space Complexity
- O(n), where n is the total number of elements inserted.