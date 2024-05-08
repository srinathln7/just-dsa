# Rpeated DNA Sequence

You are given a string `s` that represents a DNA sequence. You need to find all the 10-letter-long sequences (substrings) that occur more than once in the DNA molecule.

## Intuition

To identify repeated sequences within the DNA, we can use a sliding window approach. We'll iterate over the string and keep track of all 10-letter-long substrings encountered so far. If we encounter a substring that has already been seen, we add it to the result.

## Approach

1. Initialize an empty set `seen` to store all the 10-letter-long substrings encountered so far.
2. Initialize an empty set `result` to store the substrings that occur more than once.
3. Iterate over the string `s` using a sliding window of size 10.
4. For each window, extract the 10-letter-long substring and check if it's already in the `seen` set.
5. If the substring is already in `seen`, add it to the `result` set.
6. Otherwise, add the substring to the `seen` set.
7. After iterating over the entire string, return the elements of the `result` set.

## Time Complexity

The time complexity of this approach is O(n), where n is the length of the input string `s`. This is because we iterate over the string once using a sliding window of fixed size.

## Space Complexity

The space complexity of this approach is O(n), where n is the length of the input string `s`. This is because we use sets to store the substrings encountered so far (`seen` set) and the substrings that occur more than once (`result` set).

