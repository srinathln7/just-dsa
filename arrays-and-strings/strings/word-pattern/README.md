# [Word Pattern](https://leetcode.com/problems/word-pattern/)
Given a pattern and a string `s`, determine if `s` follows the same pattern. Here, `follow` means a bijection between a letter in `pattern` and a word in `s`.

## Intuition
To check if `s` follows the pattern, we can map each character in the pattern to a word in the string `s` and vice versa. If we encounter any inconsistencies during this mapping, the string `s` does not follow the pattern.

## Approach
1. **Split String**: Split the string `s` into words.
2. **Check Length**: If the number of words in `s` is not equal to the length of the pattern, return `false`.
3. **Mapping**:
   - Use two hash maps:
     - `chMap` to map characters from the pattern to words.
     - `strMap` to map words to characters from the pattern.
   - Iterate through the pattern and words simultaneously:
     - If a character in the pattern is already mapped to a different word or vice versa, return `false`.
     - Otherwise, establish the mapping.

## Time Complexity
- **O(n)**: We iterate through the pattern and words once, where `n` is the length of the pattern or the number of words.

## Space Complexity
- **O(k)**: We use hash maps to store the mappings, where `k` is the number of unique characters in the pattern or unique words in the string.
