# Letter Combination of Phone Number
Given a string containing digits from 2-9 inclusive, we are tasked with returning all possible letter combinations that the number could represent. The mapping of digits to letters is based on the telephone keypad.

## Intuition
To solve this problem, we can use a Depth First Search (DFS) algorithm with backtracking. We'll create a map that maps each digit to its corresponding letters on the telephone keypad. Then, we'll explore all possible combinations of letters by recursively traversing through the digits and appending the corresponding letters.

## Approach
1. Define a function `letterCombinations` that takes a string `digits` as input and returns a slice of strings representing all possible letter combinations.
2. Initialize an empty result slice to store the generated combinations.
3. Define a map `phoneMap` to map each digit to its corresponding letters on the telephone keypad.
4. Implement a DFS function `dfsWithBT` that performs DFS with backtracking.
5. In `dfsWithBT`, if the length of the current combination is equal to the length of the input `digits`, append the combination to the result slice.
6. Otherwise, iterate through the letters corresponding to the current digit using the `phoneMap`.
7. For each letter, append it to the current combination and recursively call `dfsWithBT` with the next digit.
8. After the recursive call, backtrack by removing the last character from the current combination.
9. Call `dfsWithBT` with the initial empty combination and index 0.
10. Return the result slice.

## Time  Complexity
The time complexity of the DFS algorithm is O(3^N × 4^M), where N is the number of digits with three letters (e.g., 2, 3, 4, 5, 6, 8) and M is the number of digits with four letters (e.g., 7, 9). This is because for each digit with three letters, there are three recursive calls, and for each digit with four letters, there are four recursive calls.

## Space Complexity
The space complexity is also O(3^N × 4^M) due to the space used by the recursion stack and the resultant combinations.