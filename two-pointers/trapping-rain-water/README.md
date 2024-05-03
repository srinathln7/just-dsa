# Container trapping with Rain Water
Given an elevation map represented by an array where each element represents the height of a bar, calculate how much water it can trap after raining.

## Intuition
The key idea is to visualize each unit block as a container. The amount of water each container can hold is determined by the minimum height of the left and right walls of the container. By calculating the left and right maximum heights for each index, we can determine the amount of water trapped at each index.

## Approach
1. Iterate through the elevation map to calculate the left maximum height for each index and store it in an array `leftMax`.
2. Iterate through the elevation map in reverse order to calculate the right maximum height for each index and store it in an array `rightMax`.
3. Iterate through the elevation map and for each index, calculate the amount of water trapped by taking the minimum of `leftMax` and `rightMax` at that index, and subtracting the height of the current bar.
4. Accumulate the total amount of trapped water across all indices.
5. Return the total amount of trapped water.

## Time Complexity
 O(N), where N is the number of elements in the elevation map.
- We iterate through the elevation map three times: once to calculate the left maximum heights, once to calculate the right maximum heights, and once to calculate the trapped water.


## Space Complexity:
 O(N), where N is the number of elements in the elevation map.
- We use additional space to store the left maximum heights, right maximum heights, and variables for the total amount of trapped water.
