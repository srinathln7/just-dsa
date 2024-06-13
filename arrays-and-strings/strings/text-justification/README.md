# [Text Justification](https://leetcode.com/problems/text-justification/description/?envType=list&envId=ps4mcs9r)
Given an array of words and a maximum width, format the text such that each line has exactly `maxWidth` characters and is fully (left and right) justified. You should pack your words in a greedy approach; that is, pack as many words as you can in each line. Pad extra spaces when necessary so that each line has exactly `maxWidth` characters.

Extra spaces between words should be distributed as evenly as possible. If the number of spaces on a line does not divide evenly between words, the empty slots on the left will be assigned more spaces than the slots on the right.

For the last line of text, it should be left-justified and no extra space is inserted between words.

## Intuition
The idea is to build the result line by line, adding words one by one until the current line cannot accommodate the next word without exceeding the maximum width. At that point, spaces are distributed evenly between the words in the current line to ensure full justification. For the last line, words are left-justified and any remaining spaces are added to the end of the line.

## Approach
1. **Initialization**:
   - `result`: A slice to store the final justified lines.
   - `currLine`: A slice to accumulate words for the current line.
   - `currLen`: An integer to track the total length of words in the current line.

2. **Process each word**:
   - For each word in the input array:
     - Check if adding the word to the current line would exceed the `maxWidth`.
     - If it would, distribute spaces evenly between the words in `currLine` using a cyclic modulo approach to ensure left bias in space distribution.
     - Append the fully justified line to `result` and reset `currLine` and `currLen` for the next line.
     - Add the current word to `currLine` and update `currLen`.

3. **Handle the last line**:
   - After processing all words, handle the last line separately to ensure it is left-justified. Join the words with a single space and pad the remaining space to the right to match `maxWidth`.

4. **Return the result**:
   - Return the `result` slice containing all the justified lines.

## Time and Space Complexity
- **Time Complexity**: `O(n)`, where `n` is the total number of characters in all words combined. Each word is processed once, and space distribution is handled in linear time relative to the number of words in each line.
- **Space Complexity**: `O(n)`, where `n` is the total number of characters in all words combined. This is because the result slice and intermediate structures hold the words and spaces for each line.

