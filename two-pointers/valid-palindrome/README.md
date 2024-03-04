# Valid Palindrome

The task is to determine whether a given string is a palindrome after removing non-alphanumeric characters and ignoring case sensitivity.

## Intuition:
To solve this problem, we need to preprocess the string by removing all non-alphanumeric characters and converting all characters to lowercase. After preprocessing, we can use two pointers to iterate over the string from both ends and check if it's a palindrome.

## Approach:
1. Define a mapping function `f` that takes a rune as input and returns a rune. If the input rune is not a letter or a number, it returns `-1`. Otherwise, it returns the lowercase version of the input rune.
2. Use `strings.Map()` to apply the mapping function to the input string `s`. This removes all non-alphanumeric characters and converts all characters to lowercase.
3. Initialize two pointers `l` and `r` to the start and end of the string, respectively.
4. Iterate over the string with the pointers `l` and `r`. Compare the characters at these pointers. If they are not equal, return `false`.
5. If the pointers meet (or cross each other) without encountering any inequality, return `true`.

## Time Complexity:
The time complexity of this approach is O(N), where N is the length of the input string `s`.

## Space Complexity:
The space complexity is also O(N), as we create a new string after preprocessing.

