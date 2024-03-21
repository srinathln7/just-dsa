# Minimum Window Substring 

Given two strings `s` and `t`, the task is to find the minimum window substring of `s` that contains all characters of `t`. If no such substring exists, return an empty string "".

## Intuition

The approach involves using a sliding window of variable size with a two-pointer approach (`left` and `right`). We move the `right` pointer until we encounter the first valid window where all characters in `t` are contained. Then, we adopt a greedy approach and start shrinking the window by moving the `left` pointer to find the minimum valid window.

## Approach

1. **Base Cases**:
   - If the length of `s` is less than the length of `t`, it is impossible for `s` to contain all characters in `t`. Return an empty string.
   - If `s` is equal to `t`, return `s` as it is the minimum substring.

2. Build a frequency table for string `t` to store the frequency of each character.

3. Initialize variables `minStart`, `matchCount`, and `minLen`.
   - `minStart` will store the starting index of the minimum window substring.
   - `matchCount` will keep track of the number of characters in `s` that match characters in `t`.
   - `minLen` will store the length of the minimum window substring, initialized to the maximum possible value.

4. Use two pointers `left` and `right` to define the window. Initialize `left` to 0 and iterate `right` from 0 to the end of `s`.

5. **Expand the window**:
   - Decrease the frequency of `s[right]` in the frequency table.
   - If the frequency becomes zero or positive, increment `matchCount`.

6. **Greedy Shrinking**:
   - While `matchCount` equals the length of `t`, shrink the window by moving the `left` pointer:
     - Update `minLen` and `minStart` if the current window size is smaller.
     - Increase the frequency of `s[left]` in the frequency table.
     - If the frequency becomes positive, decrement `matchCount`.
     - Increment `left`.

7. If `minLen` remains equal to the maximum possible value (`math.MaxInt32`), return an empty string as no valid window was found.

8. Return the substring of `s` from `minStart` to `minStart + minLen`.

## Time Complexity

The time complexity of this algorithm is O(n), where n is the length of string `s`. Both the `left` and `right` pointers traverse the string `s` only once.

## Space Complexity

The space complexity is O(1) since the frequency table has a fixed size of 256 characters, and the other variables used are of constant size.