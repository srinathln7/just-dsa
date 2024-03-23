# Valid Sudoku

Determine if a given 9x9 Sudoku board is valid according to Sudoku rules.

## Intuition
1. Iterate over each row and ensure that there are no repeated digits.
2. Iterate over each column and ensure that there are no repeated digits.
3. Iterate over each 3x3 sub-box and ensure that there are no repeated digits.
4. If all conditions are satisfied, the Sudoku board is valid.

## Approach
1. Implement a function `isValidBoard` to validate rows and columns.
2. Implement a function `isValidSubBox` to validate 3x3 sub-boxes.
3. Validate all rows and columns using `isValidBoard`.
4. Validate all 3x3 sub-boxes using `isValidSubBox`.
5. Transpose the board to validate columns (optional).
6. Return true if all validations pass, indicating a valid Sudoku board.

## Time Complexity
O(1) - Since board size is fixed (9 * 9)

## Space Complexity
O(1) - Since board size is fixed (9 * 9)

