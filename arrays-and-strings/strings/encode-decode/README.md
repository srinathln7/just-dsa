# Encode and Decode strings 
Design an algorithm to encode and decode a list of strings.

## Intuition
To encode the list of strings, we need to convert each string to a format where its length is prefixed, followed by a delimiter (`#`), and then the string itself. This allows us to separate different strings and know their lengths during decoding.

## Approach
#### Encoding:
1. Iterate through each string in the input list.
2. For each string, get its length and convert it to a string.
3. Append the string length, followed by a delimiter (`#`), and then the string itself, to the result string.

#### Decoding:
1. Initialize an empty result list.
2. Iterate through the encoded string.
3. Find the length of the current string segment by scanning until the delimiter (`#`).
4. Extract the substring based on the length found in step 3 and append it to the result list.

## Time Complexity
- **Encoding**: O(n), where n is the total number of characters in all strings combined.
- **Decoding**: O(n), where n is the length of the encoded string.

## Space Complexity
- **Encoding**: O(n), where n is the total number of characters in all strings combined (space used by the result string).
- **Decoding**: O(n), where n is the number of strings in the result list.