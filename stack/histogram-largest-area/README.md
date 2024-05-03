# Largest area in Histogram
Given an array `heights` representing the histogram's bar height where the width of each bar is 1, the task is to find the area of the largest rectangle in the histogram.

## Intuition:
- This solution utilizes a stack to efficiently compute the maximum area of the rectangle.
- It iterates through each bar of the histogram, maintaining a stack to keep track of bars' heights and their corresponding indices.
- The algorithm handles three scenarios: histograms in increasing order of size (extendable), decreasing order of size (not extendable), and histograms with the same size (extendable).

## Approach:
1. Initialize variables: `maxArea` to store the maximum area, and `stack` to keep track of bars' heights and indices.
2. Iterate through each bar `height` and its corresponding index `i` in the `heights` array:
   - Determine the `start` index of the current bar.
   - Handle the non-extendable case by popping out bars from the stack whose heights are greater than the current height.
   - Calculate the area of the rectangle formed by the popped bar and update `maxArea` if necessary.
   - Append the current bar (height and index) to the stack.
3. After the iteration, compare `maxArea` with the areas formed by extending the bars in the stack till the last index.
4. Return the `maxArea`.

## Time and Space Complexity:
- **Time Complexity:** The time complexity of this approach is O(n), where n is the number of bars in the histogram.
- **Space Complexity:** The space complexity is also O(n) due to the stack used to store bars' heights and indices.

This solution efficiently handles the largest rectangle area problem by utilizing a stack-based approach to process the histogram's bars.