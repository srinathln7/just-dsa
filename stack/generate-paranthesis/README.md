# Generate Parentheses
Given an integer `n`, generate all combinations of well-formed parentheses. 

## Intuition
To generate all combinations of well-formed parentheses, we can utilize backtracking. We start with an empty string and recursively add opening and closing parentheses while ensuring that the number of opening and closing parentheses are balanced.

## Approach
1. Initialize an empty slice `stack` to store the current combination of parentheses and an empty slice `res` to store all valid combinations.
2. Implement a recursive backtracking function `dfsBT` which takes two parameters: `open` representing the count of opening parentheses and `closed` representing the count of closing parentheses.
3. If both `open` and `closed` reach `n`, append the current combination of parentheses to `res`.
4. If the count of opening parentheses `open` is less than `n`, append an opening parenthesis to `stack`, call `dfsBT` recursively with `open+1` and `closed`, and then remove the appended opening parenthesis from `stack`.
5. If the count of closing parentheses `closed` is less than `open`, append a closing parenthesis to `stack`, call `dfsBT` recursively with `open` and `closed+1`, and then remove the appended closing parenthesis from `stack`.
6. Call the `dfsBT` function with initial counts of opening and closing parentheses as 0.
7. Return the `res` slice containing all valid combinations of well-formed parentheses.

## Time Complexity
The time complexity of this approach is \(O(4^n / \sqrt{n})\), where \(n\) is the given integer.

## Space Complexity
The space complexity is \(O(n)\) since we are storing the current combination of parentheses and the final result.