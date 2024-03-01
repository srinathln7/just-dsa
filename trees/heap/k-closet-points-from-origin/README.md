# K points closet to Origin

You are given a set of points in the form of coordinates on a 2D plane. You need to find the `k` closest points to the origin.

## Intuition

The problem involves calculating the Euclidean distance from each point to the origin and selecting the `k` points with the smallest distances.

## Approach

1. Define a `Point` struct to represent a point with coordinates (X, Y) and its distance from the origin.
2. Implement a min-heap to store the points based on their distances.
3. Iterate through each point in the input array.
4. Calculate the Euclidean distance for each point and add it to the min-heap.
5. Pop the `k` smallest elements from the min-heap.
6. Return these `k` closest points.

## Time Complexity

Let `n` be the number of points.

1. Calculating the Euclidean distance for each point: O(n).
2. Constructing the min-heap: O(n log n).
3. Popping `k` elements from the min-heap: O(k log n).

Overall time complexity: O(n log n).

## Space Complexity

- Space required for storing points and distances: O(n).
- Space required for the min-heap: O(n).

Total space complexity: O(n).

Alternatively, if we had used a max heap, our space complexity would be just O(k) instead of O(n).