# Problem Statement
Given a string `s`, find the length of the longest substring without repeating characters.

## Intuition
- We can use a sliding window approach to find the longest substring without repeating characters.
- We maintain a window represented by the variables `left` and `r`, where `left` is the left boundary of the window and `r` is the right boundary.
- We use a map `window` to keep track of the index of each character encountered in the string.
- **If we encounter a character that already exists in the window, we update the left boundary of the window to the next occurrence of that character.**
- We continuously update the length of the longest substring as we iterate through the string.

## Approach
1. Initialize an empty map `window` to keep track of the index of each character encountered.
2. Initialize variables `left` and `length` to keep track of the left boundary of the window and the length of the longest substring without repeating characters, respectively.
3. Iterate over the string `s` using a pointer `r`:
    - If the character `s[r]` already exists in the window and its index is greater than or equal to `left`, update `left` to the next occurrence of the character.
    - Update the index of `s[r]` in the window to `r`.
    - Update `length` to the maximum of its current value and the length of the current substring (`r - left + 1`).
4. Return `length` as the result.

## Time Complexity
The time complexity of this approach is O(n), where n is the length of the input string `s`, as we iterate over each character in the string once.

## Space Complexity
The space complexity of this approach is O(1), as the size of the map `window` is bounded by the number of characters in the English alphabet (constant space).
