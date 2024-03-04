# Max turbulent array size
Given an integer array arr, return the length of a maximum size turbulent subarray of arr. A subarray is turbulent if the comparison sign flips between each adjacent pair of elements in the subarray.

## Intuition:
The function uses a sliding window approach with two pointers `l` and `r`. It iterates through the array, maintaining a window that satisfies the conditions for a turbulent subarray. When a turbulent subarray is found, it updates the maximum length.

## Approach:
1. Initialize pointers `l` and `r` at the beginning of the array.
2. Iterate through the array with the `r` pointer.
3. Check the relationship between the current and next elements to determine if they form a turbulent subarray using a `switch` statement.
   - If the elements are equal, move the left pointer `l` to the next position of the right pointer.
   - If the conditions for a turbulent subarray are met, update the maximum length.
   - If the current elements don't form a turbulent subarray, update the left pointer `l` to the current position of the right pointer.
4. Increment the `r` pointer.
5. Continue until the end of the array is reached.

## Time Complexity:
The time complexity of this solution is O(n), where n is the length of the input array `arr`.

## Space Complexity:
The space complexity is O(1) because the algorithm uses only a constant amount of extra space regardless of the size of the input array.

Overall, this code efficiently finds the length of the longest turbulent subarray in linear time and constant space complexity.