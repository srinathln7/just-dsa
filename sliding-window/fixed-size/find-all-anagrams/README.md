# Find all anagrams in string

Given two strings `s` and `p`, return an array of all the start indices of `p`'s anagrams in `s`. You may return the answer in any order.

## Intuition:

Since both strings `s` and `p` only contain lowercase English letters, we can utilize a counting array for each of them to calculate their respective frequencies in the respective windows.

## Approach:

1. Initialize two counting arrays `countS` and `countP` of length 26, representing the frequencies of characters in strings `s` and `p`.
2. Iterate through the first window of length `k` in string `s` and update the counting arrays accordingly.
3. Check if the counting arrays for `s` and `p` are equal. If they are, add the starting index of the current window to the result array.
4. Slide the window one character at a time, updating the counting arrays and checking for anagrams.
5. Return the result array containing the starting indices of all anagrams of `p` in `s`.

## Time Complexity:** \(O(n)\), where \(n\) is the length of string `s`.

## Space Complexity:** \(O(1)\), since the size of the counting arrays is fixed at 26.