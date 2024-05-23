# [Flower Bed](https://leetcode.com/problems/can-place-flowers/)
You have a long flowerbed in which some plots are planted, and some are not. However, flowers cannot be planted in adjacent plots. Given an array `flowerbed` (containing `0` and `1`), and a number `n`, return if `n` new flowers can be planted in the `flowerbed` without violating the no-adjacent-flowers rule.

## Intuition
To determine if we can plant `n` flowers, we need to check for groups of empty plots (`0`s) in the `flowerbed` that have enough space to plant a flower without violating the rule. Specifically, three contiguous empty plots are needed to plant one flower.

## Approach
1. **Padding with Zeros**:
   - To handle edge cases where planting might be possible at the ends of the flowerbed, prepend and append a zero to the `flowerbed`.
2. **Iterate and Plant**:
   - Iterate through the padded `flowerbed` array starting from the second element and stopping before the last element.
   - Check for three contiguous zeros. If found, plant a flower in the middle plot and decrement `n`.
3. **Check the Result**:
   - After processing, if `n` is less than or equal to zero, return `true` (indicating all `n` flowers can be planted); otherwise, return `false`.

## Time Complexity
- **O(m)**: We traverse the `flowerbed` array once, where `m` is the length of the `flowerbed`.

## Space Complexity
- **O(m)**: We use an extra array to store the padded `flowerbed`.
