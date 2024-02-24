# Permuation in a string

Given two strings `s1` and `s2`, determine if one string is a permutation of the other.

## Intuition

A permutation of a string is another string with the same characters but possibly in a different order. To determine if one string is a permutation of another, we can use the sliding window technique. We'll slide a window of the same length as `s1` over `s2`, checking if any of the substrings of `s2` are permutations of `s1`.

## Approach

1. **Check Lengths**: If the length of `s2` is smaller than the length of `s1`, `s1` cannot be a permutation of `s2`. In this case, return `false`.

2. **Check Initial Substring**: Take the first substring of `s2` with the same length as `s1`. Check if it's a permutation of `s1`. If true, return `true`.

3. **Sliding Window**: Iterate through `s2` starting from index `n1`, where `n1` is the length of `s1`. At each iteration, update the substring `subStr` to the substring of `s2` starting from the current index and ending at the next `n1` characters. Check if `subStr` is a permutation of `s1`. If true, return `true`.

4. **No Permutation Found**: If no permutation is found after iterating through `s2`, return `false`.

## Time Complexity

- **Check Lengths**: O(1)
- **Check Initial Substring**: O(n1)
- **Sliding Window**: O((n2 - n1) * n1)

The overall time complexity is O((n2 - n1) * n1), where n1 is the length of `s1` and n2 is the length of `s2`.

## Space Complexity

- **Count Array**: O(26) = O(1)

The overall space complexity is O(1) since the space used is constant regardless of the input size.

