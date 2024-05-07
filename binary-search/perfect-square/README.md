# Valid Perfect Square 
Given a non-negative integer `num`, decide whether it is a perfect square.

## Intuition
A perfect square is a number that can be expressed as the product of an integer with itself. We can use binary search to efficiently find whether a given number is a perfect square.

## Approach
1. Initialize two pointers `l` and `r` to represent the search space, with `l` starting from 1 and `r` starting from `num`.
2. Perform binary search until `l` is less than or equal to `r`.
3. At each step of the binary search:
   - Calculate the middle number `mid`.
   - If the square of `mid` is equal to `num`, return true as `num` is a perfect square.
   - If the square of `mid` is less than `num`, update `l` to `mid + 1`.
   - If the square of `mid` is greater than `num`, update `r` to `mid - 1`.
4. If the binary search concludes without finding a perfect square, return false.

## Time Complexity
The time complexity of this approach is O(log(num)) as it involves binary search.

## Space Complexity
The space complexity is O(1) as it uses only a constant amount of extra space.
