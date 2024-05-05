# Unique Binary Search Trees
Given an integer `n`, return the number of structurally unique BSTs (binary search trees) which have exactly `n` nodes of unique values from 1 to `n`.

## Intuition
The number of structurally unique BSTs can be calculated using dynamic programming. At each step, we consider each number from 1 to `n` as the root of the BST and calculate the number of unique BSTs recursively for the left and right subtrees.

## Approach
1. Initialize an array `dp` of size `n+1` to store the number of unique BSTs with `i` nodes, where `dp[i]` represents the number of unique BSTs with `i` nodes.
2. Initialize `dp[0]` and `dp[1]` to 1, as there is only one unique BST with 0 or 1 node.
3. Iterate from `i=2` to `n` and calculate `dp[i]` using the formula:
   - For each number `j` from 1 to `i`, calculate the number of unique BSTs by considering `j` as the root of the BST.
   - For each root `j`, calculate the number of unique BSTs for the left subtree with `j-1` nodes and for the right subtree with `i-j` nodes.
   - Add the product of `dp[j-1]` and `dp[i-j]` to `dp[i]`.
4. Return `dp[n]`, which represents the number of unique BSTs with `n` nodes.

## Time Complexity
- The time complexity of this approach is O(n^2) because of the nested loops used to calculate `dp[i]`.

## Space Complexity
- The space complexity is O(n) to store the `dp` array.