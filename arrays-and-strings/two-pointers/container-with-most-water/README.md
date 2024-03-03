# Container with max area

Given an integer array `height` representing the heights of vertical lines, find two lines that together with the x-axis form a container that contains the most water. Return the maximum amount of water the container can store.

## Intuition:
The key idea is to use the two-pointer approach to find the maximum area. We initialize two pointers, `l` and `r`, at the beginning and end of the array, respectively. We calculate the area formed by the lines at these pointers, update the maximum area, and then move the pointer pointing to the shorter line inward. This approach guarantees that we explore all possible container configurations efficiently.

## Approach:
1. Initialize variables `maxArea` to store the maximum area and `l` and `r` as pointers pointing to the start and end of the array, respectively.
2. While `l` is less than `r`, calculate the length (difference between `r` and `l`), width (minimum of heights at `height[l]` and `height[r]`), and area (product of length and width).
3. Update `maxArea` with the maximum of the current area and the previous maximum area.
4. Move the pointer pointing to the shorter line inward (`l++` if `height[l] < height[r]`, otherwise `r--`).
5. Return `maxArea` as the result.

## Complexity Analysis:
- Time complexity: O(n), where n is the length of the input array. We traverse the array once using the two-pointer approach.
- Space complexity: O(1), as we use only a constant amount of extra space.

