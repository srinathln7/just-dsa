# Problem Statement
Given a list of strings `tokens` representing an arithmetic expression in Reverse Polish Notation (RPN) format, evaluate the expression and return the result.

## Intution
We use a stack to push numbers. The minute we enconter a operator, we pop the last two elements, calculate the intermediate result and push this result back to
the stack. At the end of the computation, we return the top value.

## Approach
1. We iterate through each token in the `tokens` list.
2. If the token is an operand, we convert it to an integer and push it onto the stack.
3. If the token is an operator (`+`, `-`, `*`, `/`), we pop the last two elements from the stack, perform the operation, and push the result back onto the stack.
4. After processing all tokens, the only element remaining in the stack is the result of the expression, which we return.

## Time Complexity
The time complexity of this approach is O(n), where n is the number of tokens in the input list.

## Space Complexity
The space complexity is also O(n), where n is the number of tokens in the input list, as we use a stack to store intermediate results.
