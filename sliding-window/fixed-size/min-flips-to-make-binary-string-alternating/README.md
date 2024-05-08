# Min. flips to make binary string alternating 

You are given a binary string `s`. You are allowed to perform two types of operations on the string in any sequence:

1. **Type-1**: Remove the character at the start of the string `s` and append it to the end of the string.
2. **Type-2**: Pick any character in `s` and flip its value, i.e., if its value is '0' it becomes '1' and vice-versa.

Return the minimum number of type-2 operations you need to perform such that `s` becomes alternating.

The string is called alternating if no two adjacent characters are equal.

## Intuition

The problem involves finding the minimum number of operations required to make the binary string alternating. We can notice that for a given length of the string, there are only two possibilities of an alternating string: either starting with '0' or starting with '1'. Therefore, we can compare the difference between the original string and these two alternating strings and take the minimum of the two. By appending the string to itself, we can turn this problem into a fixed-sliding window problem, where we consider a window of size `n` and calculate the differences.

## Approach

1. Append the string `s` to itself to create a sliding window of size `2n`.
2. Create two alternating strings: one starting with '0' and the other starting with '1'.
3. Calculate the difference between the original string `s` and these two alternating strings for the first window.
4. Slide the window of size `n` and update the differences based on the characters that enter or leave the window.
5. Track the minimum number of differences encountered during the sliding window process.

## Time Complexity

The time complexity of this approach is O(n), where n is the length of the string `s`. This is because we only need to iterate over the string twice: once to initialize the sliding window and once to slide it.

## Space Complexity

The space complexity of this approach is O(n), where n is the length of the string `s`. This is because we create two arrays of size `2n` to store the alternating strings.
