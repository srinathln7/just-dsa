package main

import "strconv"

func evalRPN(tokens []string) int {
	var stack []int

	for _, token := range tokens {
		switch token {
		case "+":
			n := len(stack)
			stack[n-2] += stack[n-1]
			stack = stack[:n-1]
		case "-":
			n := len(stack)
			stack[n-2] -= stack[n-1]
			stack = stack[:n-1]
		case "*":
			n := len(stack)
			stack[n-2] *= stack[n-1]
			stack = stack[:n-1]
		case "/":
			n := len(stack)
			stack[n-2] /= stack[n-1]
			stack = stack[:n-1]
		default:
			val, _ := strconv.Atoi(token)
			stack = append(stack, val)
		}
	}

	return stack[0]
}
