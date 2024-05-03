# Valid Palindrome 
Determine whether a given string is a palindrome, ignoring non-alphanumeric characters and case sensitivity.

## Intuition
To determine if a string is a palindrome, we can use two pointers approach. We start from both ends of the string and compare characters at each position. If the characters at corresponding positions are equal, we continue moving towards the center. If at any point, the characters are not equal, we return `false`, indicating that the string is not a palindrome. If the iteration completes without finding any unequal characters, we return `true`, indicating that the string is a palindrome.

## Approach
1. Define a filter function that takes a rune (Unicode code point) as input and returns a modified rune.
2. Inside the filter function, use the `unicode.IsLetter` and `unicode.IsNumber` functions to check if the rune is a letter or a number.
3. If the rune is neither a letter nor a number, return `-1` to drop it from the string.
4. Otherwise, convert the rune to lowercase using `unicode.ToLower` and return it.
5. Use the `strings.Map` function to apply the filter function to every character in the string.
6. Initialize two pointers `l` and `r` to the start and end indices of the filtered string, respectively.
7. Iterate over the string using the two pointers approach:
   - If the characters at positions `l` and `r` are not equal, return `false`.
   - Otherwise, continue moving towards the center by incrementing `l` and decrementing `r`.
8. If the iteration completes without finding any unequal characters, return `true`.

## Time Complexity
The time complexity of this approach is O(n), where n is the length of the input string. This is because we iterate through the string once to filter out non-alphanumeric characters and once again to compare characters.

## Space Complexity
The space complexity of this approach is O(n), where n is the length of the input string. This is because we create a new string with filtered characters, which can be at most the same length as the input string.