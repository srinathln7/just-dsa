# Sliding Window Maximum 
You are given an array of integers `nums` and an integer `k`. There is a sliding window of size `k` which is moving from the very left of the array to the very right. You can only see the `k` numbers in the window. Each time the sliding window moves right by one position.

## Intuition
To efficiently solve this problem, we need a data structure that allows constant time access to the maximum element within the sliding window and supports efficient deletion and addition of elements from both ends. A deque (double-ended queue) fits this requirement as it can store indices of elements and facilitates constant time access to the maximum element.

## Approach
1. Initialize an empty deque and an empty result slice.
2. Implement a function `cleanDeque` to remove indices from the front of the deque that are out of the current window and to remove indices from the back of the deque if the current element is greater than or equal to the corresponding elements in the deque.
3. Iterate through the array:
    - Clean the deque by calling `cleanDeque` for the current index.
    - Add the current index to the deque.
    - If the current index is greater than or equal to `k - 1`, append the maximum element from the deque to the result.
4. Return the result.

## Time Complexity
- The time complexity of this approach is O(n), where n is the number of elements in the input array `nums`. This is because we process each element once.

## Space Complexity
- The space complexity of this approach is O(n + k), where n is the number of elements in the input array `nums` and k is the size of the sliding window. This is because we use a deque to store indices, and its size can be at most k, and we also store the result array.