# Anagram Checker
Given two strings `s` and `t`, determine if they are anagrams of each other.

## Intuition
An anagram of a string is another string that contains the same characters, only the order of characters can be different. To check if two strings are anagrams, we can compare the frequency of characters in each string.

## Approach
1. Initialize a frequency table for the first string `s`.
2. Iterate through each character in `s` and increment its frequency in the table.
3. Iterate through each character in `t` and decrement its frequency in the table.
4. If the two strings are anagrams, all frequencies in the table should be zero.
5. Iterate through the frequency table, and if any value is non-zero, return false.
6. If all values in the frequency table are zero, return true.

## Time Complexity
Let `m` be the length of string `s` and `n` be the length of string `t`.
- Constructing the frequency table for `s` takes O(m) time.
- Constructing the frequency table for `t` takes O(n) time.
- Checking if two strings are anagrams by comparing the frequency tables takes O(m + n) time.
- Overall, the time complexity is O(m + n).

## Space Complexity
- We use a frequency table of size O(1) (since it contains a fixed number of unique characters - only lowercase english letters).

