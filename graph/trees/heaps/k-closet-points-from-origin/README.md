# [K Closest Points To Origin](https://leetcode.com/problems/k-closest-points-to-origin/description/)
Given a list of points in the Cartesian plane, find the k closest points to the origin (0, 0).

## Intuition:
To find the k closest points, we can calculate the Euclidean distance of each point from the origin and maintain a max heap of size k. This way, the top of the heap will always contain the k closest points seen so far.

## Approach:
1. Define a struct `Point` to store the index of the point and its distance from the origin.
2. Implement a max heap data structure (`MaxHeap`) and necessary methods (`Len`, `Less`, `Swap`, `Top`, `Push`, `Pop`) to maintain the heap property.
3. Define a function `getEuclideanDist` to calculate the Euclidean distance between a point and the origin.
4. Iterate through the given points:
   - Calculate the Euclidean distance of each point from the origin.
   - Push the point onto the max heap.
   - If the size of the heap exceeds k, pop the top element.
5. After processing all points, the heap will contain the k closest points.
6. Pop elements from the heap and store them in the result array.
7. Return the result array containing the k closest points.

## Time Complexity:
- Calculating the Euclidean distance for each point: O(N), where N is the number of points.
- Pushing and popping elements from the max heap: O(log K), where K is the size of the heap.
- Overall time complexity: O(N log K).

## Space Complexity:
- Storing points in the max heap: O(K).
- Storing the result array: O(K).
- Overall space complexity: O(K), where K is the number of closest points to be returned.
