# String Compression

The task is to compress the given array of characters `chars` in-place such that the compressed array contains groups of consecutive characters with their counts.

## Intuition
- We'll iterate through the array, keeping track of the count of consecutive characters.
- Whenever a different character is encountered, or the end of a sequence is reached, we'll compress the sequence by writing the character followed by its count.

## Approach
1. Initialize variables `count` and `resultIndex` to track the count of consecutive characters and the index where the compressed characters should be written, respectively.
2. Iterate through the `chars` array:
   - If the current character is the same as the previous one, increment the `count`.
   - If the current character is different from the previous one or it's the end of a sequence:
     - Write the character at `resultIndex`.
     - If `count` is greater than 1, convert it to a string and write its digits at consecutive indices starting from `resultIndex + 1`.
     - Update `resultIndex` accordingly.
     - Reset `count` to 1.
3. Return the length of the compressed array.

## Time Complexity
- The time complexity of this solution is O(n), where n is the length of the input `chars` array. We iterate through the array once to compress it.

## Space Complexity
- The space complexity is O(1) because we compress the characters in-place without using any additional data structures other than a few variables. Therefore, the space used is constant and does not depend on the size of the input.

#### Overall
This solution efficiently compresses the characters in the given array in-place with constant space complexity.