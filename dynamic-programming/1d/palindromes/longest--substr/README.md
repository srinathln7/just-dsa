# Longest Palindromic Substring

Given a string `s`, find the longest palindromic substring in `s`. You may assume that the maximum length of `s` is 1000.

## Intuition

To find the longest palindromic substring, we can adopt an approach where we pick every character in the string and start expanding outwards until we don't go out-of-bound. Since picking every character and starting to expand outwards would always result in an odd number of characters, we also do the same for every pair of contiguous characters so that we also check for even length palindromes.

## Approach

1. Iterate through each character in the string.
2. For each character, expand outwards to find the longest palindrome centered at that character.
3. Update the longest palindrome found so far.
4. Return the longest palindrome.

## Time Complexity

The time complexity of this approach is O(n^2), where n is the length of the input string `s`. This is because for each character in the string, we expand outwards to check for palindromic substrings. We have essentially reduced the palidromic check for every substring to be O(1).

## Space Complexity

The space complexity is O(1) because we are not using any extra space that grows with the size of the input.

