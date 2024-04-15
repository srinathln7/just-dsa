# Blockchain 

A blockchain is a chain of connected block such that each block keeps the hash of its previous block.
You are given a list of blocks represented by pair of numbers. You have to find the maximum length that can be made by connecting these blocks.
The rule for connecting the blocks is:
Suppose a pair (x,y) is followed by (m,n). Then y must be less than m.

Input format:
First line contains n (number of pairs).
Each of the next n lines contains a and b separated by space of pair (a,b).
Every pair will be in the form (a,b) where a<b.
 
Output Format:
Print the maximum length of blockchain that can be made.

## Intuition
To solve this problem, we can use dynamic programming to find the longest increasing subsequence (LIS) among the pairs.

To find the length of the longest increasing subsequence, we can use dynamic programming. We iterate through the pairs from the end to the beginning, and for each pair, we check all pairs to the right of it. If we find a pair whose first element is greater than the second element of the current pair, we update the length of the longest increasing subsequence ending at the current pair. Finally, we find the maximum length among all pairs.

## Approach
1. Initialize a dynamic programming array `dp` with all elements set to 1.
2. Iterate through the pairs from the end to the beginning.
3. For each pair, iterate through all pairs to the right of it.
4. If we find a pair whose first element is greater than the second element of the current pair, update the length of the longest increasing subsequence ending at the current pair.
5. Find the maximum length among all pairs.
6. Return the maximum length.

## Time Complexity
The time complexity of this approach is O(n^2), where n is the number of pairs. This is because we iterate through all pairs twice.

## Space Complexity
The space complexity is O(n), where n is the number of pairs, as we use a dynamic programming array of size n to store the lengths of the longest increasing subsequences.
