# Interleaving of strings
You are given strings `s1`, `s2`, and `s3`. Determine whether `s3` can be formed by interleaving the characters of `s1` and `s2`, while preserving the order of characters from each string.

An interleaving of two strings `s` and `t` is a configuration where `s` and `t` are divided into `n` and `m` substrings respectively, such that:
- `s = s1 + s2 + ... + sn`
- `t = t1 + t2 + ... + tm`
- The relative order of the characters from each string is preserved.

## Intuition
To check if `s3` can be formed by interleaving `s1` and `s2`, we can use dynamic programming. We can create a 2D array `dp` where `dp[i][j]` represents whether the first `i` characters of `s1` and the first `j` characters of `s2` can form the first `i + j` characters of `s3`.

## Approach
1. Check if the lengths of `s1` and `s2` add up to the length of `s3`. If not, return `false`.
2. Create a 2D boolean array `dp` of size `(len(s1) + 1) x (len(s2) + 1)`.
3. Initialize `dp[0][0]` as `true` since two empty strings can form an empty string.
4. Fill the first column of `dp`:
   - For each `i` from `1` to `len(s1)`, set `dp[i][0]` as `true` if `s1[:i]` equals `s3[:i]`.
5. Fill the first row of `dp`:
   - For each `j` from `1` to `len(s2)`, set `dp[0][j]` as `true` if `s2[:j]` equals `s3[:j]`.
6. Fill the remaining entries of `dp`:
   - For each `i` from `1` to `len(s1)` and each `j` from `1` to `len(s2)`, set `dp[i][j]` as `true` if either:
     - `s1[i-1]` matches `s3[i+j-1]` and `dp[i-1][j]` is `true`, or
     - `s2[j-1]` matches `s3[i+j-1]` and `dp[i][j-1]` is `true`.
7. Finally, return `dp[len(s1)][len(s2)]`.

## Time Complexity
The time complexity of this approach is O(m * n), where m and n are the lengths of strings s1 and s2, respectively.

## Space Complexity
The space complexity of this approach is O(m * n) since we are using a 2D array `dp` of size `(len(s1) + 1) x (len(s2) + 1)`.