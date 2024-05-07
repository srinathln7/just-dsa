# Binary Search

## Intuition

This algorithm efficiently searches for a target value within a **sorted** array. The key intuition behind binary search is that by comparing the target value to the middle element of the array, one can determine which half of the array the target is likely to be in (if it is present). The search then continues recursively or iteratively in the identified half, effectively halving the search space with each step.

## Approach

1. **Initialize Pointers:** Two pointers, `l` and `r`, are initialized to point to the start and end of the array, respectively.
2. **Iterative Search:** The search proceeds iteratively until `l` exceeds `r`. Within each iteration:
   - The middle index `mid` is calculated.
   - If the middle element `nums[mid]` matches the target, the search is successful, and `mid` is returned.
   - If the target is greater than `nums[mid]`, the search space is narrowed to the right half by setting `l` to `mid + 1`.
   - If the target is less than `nums[mid]`, the search space is narrowed to the left half by setting `r` to `mid - 1`.
3. **Completion:** If the loop exits without finding the target, `-1` is returned to indicate that the target is not present in the array.

## Time Complexity

- **Best Case:** \(O(1)\) when the target is at the middle of the array.
- **Average and Worst Case:** \(O(\log n)\), where \(n\) is the number of elements in the array. This is because with each comparison, the search space is halved.

## Space Complexity

- **O(1):** The space complexity is constant as the algorithm uses a fixed amount of space (a few variables) and does not depend on the input array size.

### Example

Consider searching for the target `3` in the array `nums = [1, 2, 3, 4, 5]`.

- Initial setup: `l = 0`, `r = 4`.
- First iteration: `mid = 2`, `nums[mid] = 3`, which matches the target. The function returns `2`.

This example demonstrates a best-case scenario where the target is found in the middle of the array on the first try.