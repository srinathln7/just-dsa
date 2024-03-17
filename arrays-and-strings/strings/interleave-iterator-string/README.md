**Problem:**

Given a string, we need to print the words vertically such that each word occupies one column. Trailing spaces should be trimmed.

**Intuition:**

To print the words vertically, we first split the string into individual words. We then determine the maximum length of the words to find the number of columns required. Next, we pad each word with spaces to match the maximum length. Finally, we construct the vertical output by appending characters from each word at their respective positions in the columns.

**Approach:**

1. Split the input string into individual words.
2. Determine the maximum length among all words to find the number of columns required.
3. Pad each word with spaces to match the maximum length.
4. Construct the vertical output by appending characters from each word at their respective positions in the columns.
5. Trim trailing right spaces from each column.
6. Return the vertical output as a slice of strings.

**Time Complexity:**

Let `n` be the total number of characters in the input string, and `m` be the number of words.

- Splitting the string: O(n).
- Finding the maximum length: O(m).
- Padding words: O(m).
- Constructing the vertical output: O(n).
- Trimming trailing spaces: O(n).

The overall time complexity is O(n).

**Space Complexity:**

- Space required for the output slice: O(m), where m is the number of columns (equal to the maximum length of words).
- Space required for intermediate variables: O(n).

The overall space complexity is O(n).