# Longest common subsequence

Given two strings `text1` and `text2`, the task is to find the length of their longest common subsequence. A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters. A common subsequence of two strings is a subsequence that is common to both strings.

## Intuition:
The problem can be solved using dynamic programming, specifically by filling a 2D array `dp` where `dp[i][j]` represents the length of the longest common subsequence between substrings `text1[0...i-1]` and `text2[0...j-1]`. By iteratively filling this array based on the characters of the input strings, we can efficiently compute the length of the longest common subsequence.

## Approach:
1. Initialize a 2D array `dp` of size `(m+1) x (n+1)` to store the lengths of common subsequences for different substrings of `text1` and `text2`.
2. Iterate over each character of both strings and fill the `dp` array based on the characters' matches.
3. If characters at corresponding positions match, increment the length of the common subsequence by 1. Otherwise, take the maximum of the lengths of common subsequences obtained by excluding either character.
4. Finally, return `dp[m][n]`, where `m` and `n` are the lengths of `text1` and `text2` respectively.

## Time Complexity:
The time complexity of this solution is O(m * n), where `m` and `n` are the lengths of the input strings `text1` and `text2`. This complexity arises from the nested loop that iterates over each character of both strings.

## Space Complexity:
The space complexity is also O(m * n) due to the usage of a 2D array `dp` with dimensions `(m+1) x (n+1)`, where `m` and `n` are the lengths of the input strings `text1` and `text2`. This array stores the lengths of common subsequences for different substrings of `text1` and `text2`, leading to quadratic space complexity.
The space complexity of this approach can be further optimzed to O(n), where n is the length of text2. This is achieved by using only two arrays, dp1 and currentRow, instead of the entire 2D dynamic programming array.