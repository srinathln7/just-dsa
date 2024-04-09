# Reverse a 32-bit integer
Given a signed 32-bit integer `x`, return `x` with its digits reversed. If reversing `x` causes the value to go outside the signed 32-bit integer range \([-2^{31}, 2^{31} - 1]\), then return `0`.

### Intuition
To reverse the digits of the integer `x`, we can repeatedly extract the last digit of `x`, add it to the result after shifting the result one decimal place to the left, and update `x` by removing the last digit. We need to handle overflow conditions where the reversed number exceeds the range of a 32-bit signed integer.

### Approach
1. Initialize `result` as 0.
2. Iterate until `x` becomes 0:
    - Extract the last digit of `x` using the modulus operator `%`.
    - Check for overflow conditions:
        - If `result` is greater than `INT_MAX/10` or if it equals `INT_MAX/10` and the current digit is greater than `INT_MAX%10`, return 0 to indicate overflow (max-out condition).
        - If `result` is less than `INT_MIN/10` or if it equals `INT_MIN/10` and the current digit is less than `INT_MIN%10`, return 0 to indicate overflow (min-out condition).
    - Update `result` by multiplying it by 10 and adding the current digit.
    - Update `x` by removing the last digit using integer division `/`.
3. Return the `result`.

### Time Complexity
The time complexity of this approach is \(O(\log|x|)\), where \(|x|\) represents the number of digits in the input integer `x`.

### Space Complexity
The space complexity is \(O(1)\) since we are using only a constant amount of extra space.