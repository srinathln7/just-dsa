# Search Insert Position
Given a sorted array of distinct integers `nums` and a target value `target`, return the index if `target` is found in the array. If not, return the index where `target` would be if it were inserted in order.

## Intuition
To find the index where `target` should be inserted into the sorted array `nums`, we can utilize binary search. If the target is found in the array, we return its index. If not, we adjust the search space based on the comparison of `target` with the elements at the middle index during each iteration.

## Approach
1. Initialize two pointers `l` and `r` to the start and end of the array, respectively.
2. Perform binary search while `l` is less than or equal to `r`.
3. Calculate the middle index `mid` as `(l + r) / 2`.
4. If the target is found at index `mid`, return `mid`.
5. If the target is less than the element at index `mid`, set `r` to `mid - 1`.
6. If the target is greater than the element at index `mid`, set `l` to `mid + 1`.
7. After the loop, if the target is not found, `l` will represent the correct index to insert `target`.
8. Return `l`.

## Time Complexity
The time complexity of this approach is O(log n), where n is the length of the input array `nums`. This is because binary search reduces the search space by half in each iteration.

## Space Complexity
The space complexity is O(1) as we use only a constant amount of extra space for storing variables.
