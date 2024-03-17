# Count Palindromic Substrings
Given a string `s`, return the number of palindromic substrings in `s`. A substring is a contiguous sequence of characters within the string.

## Intuition

To count the number of palindromic substrings, we adopt an approach where we pick every character in the string and start expanding outwards until we don't go out-of-bound. Since picking every character and starting to expand outwards would always result in an odd number of characters, we also do the same for every pair of contiguous characters so that we also check for even length palindromes. As long as the palidromic conditions hold, then we keep incrementing the counter.

## Approach

1. Iterate through each character in the string.
2. For each character, expand outwards to find all palindromic substrings centered at that character.
3. Update the count of palindromic substrings found so far.
4. Return the total count of palindromic substrings.

## Time Complexity

The time complexity of this approach is O(n^2), where n is the length of the input string `s`. This is because for each character in the string, we expand outwards to check for palindromic substrings.

## Space Complexity

The space complexity is O(1) because we are not using any extra space that grows with the size of the input.

