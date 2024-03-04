# Product of array except self

Given an integer array `nums`, return an array `result` such that `result[i]` is equal to the product of all the elements of `nums` except `nums[i]`.

**Intuition:**

The solution leverages the concept of two pointers, `prefix` and `suffix`, to keep track of the product of elements to the left and right of each index, respectively.

**Approach:**

1. Initialize an array `result` of the same length as `nums` to store the final products.
2. Initialize two variables `prefix` and `suffix` to keep track of the product of elements to the left and right of the current index, respectively.
3. Calculate Prefix Products:
   - Iterate over `nums` from left to right.
   - Store the product of all elements to the left of the current index in the `result` array.
   - Update `prefix` by multiplying it with the current element `nums[i]`.
4. Calculate Suffix Products:
   - Iterate over `nums` from right to left.
   - Multiply the previously calculated product stored in the `result` array by the product of all elements to the right of the current index.
   - Update `suffix` by multiplying it with the current element `nums[i]`.
5. Return the `result` array containing the product of all elements in `nums` except themselves.

**Time Complexity:**

The time complexity of this approach is O(n), where n is the length of the input array `nums`, as we perform two passes over the array.

**Space Complexity:**

The space complexity of this approach is O(1) excluding the output array `result`, as we only use a constant amount of extra space for variables `prefix` and `suffix`.

