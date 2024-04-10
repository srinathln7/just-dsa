package main

func isValid(s string) bool {

	// Key Idea: Let us use a stack to determine if the given string is a valid parentheses or not
	// Every time we encounter a open braces we push the corresponding closed braces to the stack and
	// everytime we encounter a closed brace we pop the top from stack and check if it is the same.
	// By the end of the iteration, the stack should be empty

	var stack []rune
	for _, ch := range s {
		switch ch {
		case '(':
			stack = append(stack, ')')
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		default:
			if len(stack) == 0 || ch != stack[len(stack)-1] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}
