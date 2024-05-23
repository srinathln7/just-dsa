
# Length of Last Word 
Given a string `s` consisting of words and spaces, return the length of the last word in the string. A word is defined as a sequence of non-space characters.

## Intuition
To find the length of the last word in the string, we can split the string into words using spaces as delimiters. By iterating from the end of the resulting list of words, we can find the last non-empty word and return its length.

## Approach
1. Use the `strings.Split` function to split the string `s` by spaces into a slice of words.
2. Iterate from the end of the slice to find the last non-empty word.
3. Return the length of the found word.

## Time Complexity
- **O(n)**: The `strings.Split` function traverses the entire string `s`, where `n` is the length of the string.

## Space Complexity
- **O(k)**: The extra space required for the slice of words, where `k` is the number of words in the string.
