# Palindrome Partition

Given a string `s`, partition `s` such that every substring of the partition is a palindrome. Return all possible palindrome partitioning of `s`.

## Intuition:

To generate all possible palindrome partitioning of the given string, we can use a Depth-First Search (DFS) approach with backtracking. We'll explore all possible substrings of the input string and check if each substring is a palindrome. If it is, we'll append it to our current partition and recursively explore further. After exploring all possible substrings, we'll backtrack and remove the last appended substring from the partition.

## Approach:

1. Initialize an empty slice `result` to store all possible partitions.
2. Implement a DFS function that takes the starting index of the substring to consider.
3. In the DFS function:
   - If the starting index exceeds the length of the string, append the current partition to the result.
   - Iterate over all possible ending indices for the substring, starting from the current starting index.
   - Check if the substring from the starting index to the current ending index is a palindrome using the `isPalindrome` function.
   - If it is a palindrome, append it to the current partition, recursively call the DFS function with the next starting index, and backtrack by removing the last appended substring.
4. Start the DFS function with the starting index 0.
5. Return the `result` containing all possible palindrome partitions.

## Time Complexity:

O(2^n)

## Space Complexity:

The space complexity is dominated by the storage of the result, which can contain up to \( O(2^n) \) partitions, each partition storing a substring of the input string. Therefore, the space complexity is  `O(2^n . n)` Additionally, the recursion stack space for the DFS function requires \( O(n) \) space.