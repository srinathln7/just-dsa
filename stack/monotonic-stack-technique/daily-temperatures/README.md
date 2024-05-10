# Daily Temperatures
Given an array of integers `temperatures` representing daily temperatures, return an array `result` such that `result[i]` is the number of days you have to wait after the ith day to get a warmer temperature. If there is no future day for which this is possible, keep `result[i]` equal to 0.

## Intuition
We can solve this problem using a stack data structure. The idea is to traverse the temperatures array from right to left and maintain a stack that keeps track of indices of elements whose temperatures we have not yet found a warmer day for. Whenever we encounter a temperature higher than the temperature of the current element, we pop indices from the stack and calculate the number of days to wait for a warmer temperature.

## Approach
1. Initialize an empty stack to store indices of temperatures.
2. Initialize an array `result` of the same length as the input `temperatures` to store the number of days to wait for a warmer temperature.
3. Traverse the `temperatures` array from right to left.
4. For each temperature `temperatures[i]`, do the following:
   - While the stack is not empty and the temperature at the top of the stack (i.e., `temperatures[stack[top]]`) is less than or equal to `temperatures[i]`, pop elements from the stack and calculate the number of days to wait for a warmer temperature.
   - If the stack is not empty after popping elements, update the `result` array with the difference between the current index `i` and the index at the top of the stack.
   - Push the current index `i` onto the stack.
5. Return the `result` array.

## Time Complexity
- The time complexity of this approach is O(n), where n is the number of elements in the `temperatures` array. This is because we traverse the array only once.

## Space Complexity
- The space complexity of this approach is O(n), where n is the number of elements in the `temperatures` array. This is because we use an additional stack of size O(n) to store indices of temperatures.