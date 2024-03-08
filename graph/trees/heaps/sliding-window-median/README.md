# Sliding Window Median 

You are given an array of integers `nums` and an integer `k`. For each window of size `k` in the array, find the median of the elements in that window. Return an array of the median values.

## Intuition

To efficiently find the median of each window, we can use two heaps: a max-heap to store the smaller half of the elements and a min-heap to store the larger half. By maintaining these heaps, we can quickly access the median element(s) within each window.

## Approach

1. Define two heap structures: `MinHeap` and `MaxHeap`, representing the min-heap and max-heap respectively.
2. Implement heap operations for both structures: `Push`, `Pop`, `Top`, `Len`, and `Empty`.
3. Implement a function `getMedian` to calculate the median based on the current state of the heaps.
4. Initialize two empty heaps (`minHeap` and `maxHeap`) and a map to track elements moving out of the window (`outOfWindow`).
5. Iterate through the array and populate the heaps for the first window.
6. Slide the window, updating the heaps and calculating the median for each subsequent window.
7. Append the calculated medians to the result array.
8. Return the result array.

## Time Complexity

The time complexity of this approach is O(N log K), where N is the number of elements in the input array and K is the size of the sliding window. This complexity arises from the heap operations performed for each element.

## Space Complexity

The space complexity of this approach is O(K), where K is the size of the sliding window. This space is used to store the elements within the window and the heaps used for calculating the medians.

