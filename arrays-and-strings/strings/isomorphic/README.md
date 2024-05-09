# Isomorphic Strings

Given two strings `s` and `t`, determine if they are isomorphic. Two strings are isomorphic if the characters in `s` can be replaced to get `t`.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.

## Intuition

To check if two strings `s` and `t` are isomorphic, we need to ensure that each character in `s` can be mapped to a unique character in `t`, and vice versa. We can use two hash maps to store the mappings from characters in `s` to characters in `t`, and vice versa. We iterate through each character in `s` and `t`, check if the mapping exists in both directions, and update the hash maps accordingly.

## Approach

1. Initialize two hash maps `charMapSet` and `revCharMapSet` to store the mappings from characters in `s` to characters in `t`, and vice versa, respectively.
2. Iterate through each character in `s` and `t`.
3. For each character `s[i]` and `t[i]`, check if the mapping exists in both `charMapSet` and `revCharMapSet`.
   - If a mapping exists in `charMapSet` for `s[i]` and it does not match `t[i]`, return `false`.
   - If a mapping exists in `revCharMapSet` for `t[i]` and it does not match `s[i]`, return `false`.
4. If no violations are found, update the hash maps with the current mapping.
5. After iterating through all characters, return `true` if no violations are found.

## Time Complexity

The time complexity of this approach is O(n), where n is the length of the input strings `s` and `t`. We iterate through each character once.

## Space Complexity

The space complexity is O(n) as well. We use two hash maps to store the mappings, each of which can store up to n mappings.
