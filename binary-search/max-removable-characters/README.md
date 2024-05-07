# Maximum number of removable characters
You are given two strings `s` and `p`, where `p` is a subsequence of `s`. Additionally, you are given a distinct 0-indexed integer array `removable` containing a subset of indices of `s`. You want to choose an integer `k` (0 <= k <= len(removable)) such that, after removing `k` characters from `s` using the first `k` indices in `removable`, `p` is still a subsequence of `s`. Return the maximum `k` you can choose such that `p` is still a subsequence of `s` after the removals.

## Intuition
To find the maximum `k` such that `p` remains a subsequence of `s` after removing characters indexed by `removable`, we can use binary search. We can define a function `isSubsequence` to check if `p` is still a subsequence of `s` after removing `k` characters. Then, we perform binary search on the indices of `removable` to find the maximum `k` where `p` remains a subsequence.

## Approach
1. Define a function `isSubsequence` to check if `p` is still a subsequence of `s` after removing `k` characters.
2. Inside `isSubsequence`, create a set `removableSet` to store the indices of characters to be removed.
3. Iterate over the characters of `s` and check if `p` is still a subsequence after removing characters marked for removal.
4. Implement binary search to find the maximum `k`.
5. Initialize `l` as 0 and `r` as the length of `removable`.
6. While `l` is less than or equal to `r`, calculate the middle index `mid` and check if `p` remains a subsequence after removing `mid` characters. Adjust `l` or `r` accordingly.
7. Return `r`, which represents the maximum `k` such that `p` remains a subsequence of `s` after the removals.

## Time Complexity
The time complexity of this approach is O(n * log k), where n is the length of string `s` and m is the length of the removable array. The binary search takes O(log k) iterations, and for each iteration, we check if `p` is a subsequence of `s`, which takes O(n) time.

## Space Complexity
The space complexity is O(n) for storing the removable set.
