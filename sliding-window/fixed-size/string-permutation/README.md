# Permuatation Inclusion

You are given two strings, `s1` and `s2`. You need to determine if `s2` contains a permutation of `s1`.

## Intuition

To check if `s2` contains a permutation of `s1`, we can use a sliding window approach with a fixed size equal to the length of `s1`.

## Approach

1. Initialize two arrays `countS1` and `countS2` of size 26, representing the count of each character in `s1` and `s2` respectively.
2. Calculate the counts of characters in the first window of size `k` (length of `s1`) for both `s1` and `s2`.
3. Check if the counts of characters in `s1` and `s2` are equal. If they are, return `true`.
4. Slide the window of size `k` over `s2` by removing the leftmost character and adding the next character to the window.
5. Recalculate the counts of characters in the updated window and compare them with the counts of `s1`.
6. Repeat steps 3-5 until all windows of size `k` are checked or a permutation is found.

## Time Complexity

The time complexity of this approach is O(n), where n is the length of `s2`. This is because we only need to iterate over `s2` once to perform the sliding window operation.

## Space Complexity

The space complexity of this approach is O(1). Although we use two arrays of size 26 (`countS1` and `countS2`), the space usage remains constant and does not depend on the input size.


