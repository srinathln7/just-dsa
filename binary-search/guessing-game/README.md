# Guessing Game

You are playing a guessing game with a computer. The computer picks a number between 1 and n (inclusive), and you need to guess it. Each time you guess a number, the computer will tell you if the guess is higher, lower, or correct.

**Intuition:**
Given that the target number is between 1 and n, and the guess response can be used to narrow down the search space, we can use binary search to efficiently find the target number. **Remember** this game is designed to be always eventually won provided the input parameters are right.

**Approach:**
1. Initialize two pointers, `l` and `r`, to 1 and n, respectively, representing the range of possible numbers.
2. Perform binary search until `l` is less than or equal to `r`:
   - Calculate the middle number `try` as `(l + r) / 2`.
   - Use the `guess` function to determine if `try` is the correct number:
     - If the guess is correct (returns 0), return `try` as the result.
     - If the guess indicates that `try` is higher (returns -1), update `r` to `try - 1`.
     - If the guess indicates that `try` is lower (returns 1), update `l` to `try + 1`.
3. If the loop exits without finding the correct number, return -1 to indicate that the number was not found.

**Time Complexity:**
The time complexity of this algorithm is O(log n) because we use binary search to narrow down the search space.

**Space Complexity:**
The space complexity of this algorithm is O(1) because we only use a constant amount of extra space for variables regardless of the size of the input.

This code has been modified from the previous version to return -1 if the number is not found within the range 1 to n. This ensures that the function always returns a valid integer, even if the target number is not present in the given range.
