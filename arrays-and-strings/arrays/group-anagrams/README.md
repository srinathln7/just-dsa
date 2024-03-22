# Group Anagrams
Given an array of strings `strs`, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

## Intuition:
We can efficiently group anagrams together by using a hashmap. Each key in the hashmap will be a sorted version of a string, and the corresponding value will be a list of strings that are anagrams of each other.

## Approach:
1. Iterate through each string in the input array `strs`.
2. Sort the characters of each string and use the sorted version as the key in the hashmap.
3. Append the original string to the corresponding group in the result slice.
4. Finally, filter out any empty groups from the result slice and return the grouped anagrams.

## Time Complexity:
- Sorting each string takes O(m * log m) time, where `m` is the maximum length of a string.
- Since we iterate through each string in the input array once, the overall time complexity is O(n * m * log m), where `n` is the number of strings in the input array `strs`.

## Space Complexity:
- The space complexity is O(n), where `n` is the number of strings in the input array `strs`, for storing the result and the hashmap.


