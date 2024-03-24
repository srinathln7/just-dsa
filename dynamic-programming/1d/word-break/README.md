# Word Break
Given a string `s` and a dictionary of strings `wordDict`, determine if `s` can be segmented into a space-separated sequence of one or more dictionary words.

## Intuition
This problem can be solved using dynamic programming. We iterate through each character in the string and check if the substrings ending at that character can be segmented into words from the dictionary.

## Approach
1. Create a set to store words from `wordDict` for efficient lookup.
2. Initialize a dynamic programming array `dp` of size `len(s)+1`. `dp[i]` represents whether the substring from index 0 to i in `s` can be segmented into words from `wordDict`.
3. Iterate through each character in `s` and for each character, iterate through all possible substrings ending at that character.
4. If the substring from 0 to j is segmented (`dp[j]` is true) and the substring from j to i is in `wordDict`, set `dp[i]` to true.
5. Finally, return `dp[len(s)]` which represents whether the entire string `s` can be segmented into words from `wordDict`.

## Time Complexity
The time complexity of this approach is O(n^2), where n is the length of the string `s`. This is because we have nested loops iterating through each character and each possible substring of `s`.

## Space Complexity
O(n+m). The space complexity is O(n) where n is the length of the string `s`, due to the dynamic programming array `dp`. Additionally, the space complexity of the set storing words from `wordDict` is O(m), where m is the total number of characters in all words in `wordDict`.