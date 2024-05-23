# Maximum number of balloons
Given a string `text`, determine the maximum number of times the word "balloon" can be formed using the characters in `text`.

## Intuition
To find the maximum number of times the word "balloon" can be formed, we need to count the occurrences of each character in the string `text` and consider the minimum count of the characters required to form the word "balloon". Special attention is needed for characters that appear more than once in the word "balloon" (i.e., 'l' and 'o').

## Approach
1. **Count Characters**: Count the occurrences of each character in the string `text`.
2. **Determine Minimum Count**: Calculate the minimum number of times the word "balloon" can be formed by considering the counts of the characters 'b', 'a', 'l', 'o', and 'n'. Note that 'l' and 'o' need to be divided by 2 since they appear twice in "balloon".
3. **Return the Result**: The minimum value among these counts represents the maximum number of times "balloon" can be formed.

## Time Complexity
- **O(n)**: We iterate through the string once to count character occurrences.

## Space Complexity
- **O(1)**: We use a fixed-size array to store character counts, regardless of the input size.

