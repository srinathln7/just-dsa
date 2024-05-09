# Next Greater Element
You are given two distinct 0-indexed integer arrays `nums1` and `nums2`, where `nums1` is a subset of `nums2`. For each element in `nums1`, find the next greater element in `nums2` that is to the right of the element. If there is no such greater element, set the result as -1.

## Intuition
To find the next greater element for each element in `nums1`, we can use a monotonic stack technique. By iterating through `nums2`, we can maintain a stack of elements such that the top of the stack always contains elements smaller than the current element. If we encounter an element larger than the top of the stack, we pop elements from the stack and update the result for those elements.

## Approach
1. Initialize an array `result` to store the next greater elements for elements in `nums1`. Set all elements of `result` to -1 initially.
2. Create a map `n1Set` to store the indices of elements in `nums1`.
3. Iterate through `nums2`:
   - For each element `num` in `nums2`, check if it is present in `n1Set`.
   - If `num` is greater than the top element of the stack, pop elements from the stack and update the result for those elements.
   - Append `num` to the stack.
4. Return the `result` array.

## Time Complexity
- The time complexity of this approach is O(N + M), where N is the length of `nums1` and M is the length of `nums2`. This is because we iterate through both arrays once.

## Space Complexity
- The space complexity is O(N), where N is the length of `nums1`, as we use additional space to store the result array and the map `n1Set`.