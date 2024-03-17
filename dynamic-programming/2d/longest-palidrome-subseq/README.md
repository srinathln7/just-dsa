## Longest Palindromic Subsequence

Given a string `s`, find the length of the longest palindromic subsequence in `s`.

## Intuition

To find the longest palindromic subsequence in a string, we can use dynamic programming. We observe that the longest palindromic subsequence of a string is related to the longest common subsequence of the string and its reverse.

## Approach

1. Reverse the string `s` to obtain `s_rev`.
2. Use dynamic programming to find the length of the longest common subsequence between `s` and `s_rev`.
3. The length of the longest palindromic subsequence is the same as the length of the longest common subsequence between `s` and `s_rev`.

## Time Complexity

The time complexity of this approach is O(n^2), where n is the length of the input string `s`. This is because we use dynamic programming to find the longest common subsequence.

## Space Complexity

The space complexity is O(n^2) because we use a 2D array to store the results of subproblems in dynamic programming.


## How does the Longest common subsequence of a string and its reverse give the longest palidromic substring

1. **Palindrome**: A palindrome is a string that reads the same backward as forward. For example, "level" and "racecar" are palindromes.

2. **Subsequence**: A subsequence of a string is a new string that is formed from the original string by deleting some (or none) of the characters without changing the relative order of the remaining characters. For example, "ace" is a subsequence of "abcde".

3. **Longest Palindromic Subsequence (LPS)**: The longest palindromic subsequence of a string is the longest subsequence that is also a palindrome. For example, for the string "bbbab", the longest palindromic subsequence is "bbbb" or "bbab".

4. **Longest Common Subsequence (LCS)**: The longest common subsequence of two strings is the longest subsequence that is common to both strings. For example, for the strings "abcde" and "ace", the longest common subsequence is "ace".

Now, let's consider the relationship between LPS and LCS:

- Suppose we have a string `s`.
- Let `s_rev` be the reverse of the string `s`.

We observe that finding the longest palindromic subsequence in `s` is equivalent to finding the longest common subsequence between `s` and `s_rev`.

Why is this the case?

- Let `lps` be the longest palindromic subsequence of `s`.
- Let `lcs` be the longest common subsequence between `s` and `s_rev`.

Since `s_rev` is the reverse of `s`, finding the longest common subsequence between `s` and `s_rev` is like finding the common characters between `s` and `s_rev` that are arranged in the same order (though not necessarily contiguous) in both strings.

Now, consider the palindromic property:

- If a sequence `x` is a palindrome, then its reverse is also `x`.

Therefore, if `lps` is a palindrome, then its reverse is also `lps`.

So, the longest common subsequence `lcs` between `s` and `s_rev` will be the same as `lps`.

Hence, the length of the longest palindromic subsequence is the same as the length of the longest common subsequence between `s` and `s_rev`. This is why we use dynamic programming to find the length of the LCS between `s` and `s_rev` to find the length of the LPS of `s`.