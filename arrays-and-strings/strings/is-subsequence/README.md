# [Is Subsequence](https://leetcode.com/problems/is-subsequence/)
Given two strings `s` and `t`, determine if `s` is a subsequence of `t`. A subsequence of a string is a new string formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

## Intuition:
We can use a two-pointer approach to iterate through both strings. We move one pointer (`idx`) through string `s`, and another pointer through string `t`. We check if each character in `s` matches the corresponding character in `t`. If all characters in `s` are found in `t` in the same relative order, `s` is considered a subsequence of `t`.

## Approach:
1. Initialize two pointers, `idx` and `i`, to 0.
2. Iterate through string `t` using pointer `i`.
3. For each character in `t`, compare it with the character at position `idx` in string `s`.
4. If the characters match (`s[idx] == t[i]`), increment `idx`.
5. Continue this process until either `idx` reaches the end of string `s` or the end of string `t` is reached.
6. If `idx` reaches the end of string `s`, return `true`, indicating that `s` is a subsequence of `t`.
7. If the end of string `t` is reached and `idx` is not at the end of string `s`, return `false`, indicating that `s` is not a subsequence of `t`.

## Time Complexity:
- The time complexity of this algorithm is O(n), where n is the length of string `t`. We only need to iterate through `t` once to check if `s` is a subsequence.

## Space Complexity:
- The space complexity is O(1) since we only use a constant amount of extra space for storing temporary variables and pointers.