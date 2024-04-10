package main

func generateParenthesis(n int) []string {

	// Key Idea: We need to draw a decision tree to generate all possible combinations. To do this,
	// we implement standard DFS with backtracking solution

	var stack []rune
	var res []string

	var dfsBT func(int, int)
	dfsBT = func(open, closed int) {
		if open == n && closed == n {
			res = append(res, string(stack))
			return
		}

		if open < n {
			stack = append(stack, '(')
			dfsBT(open+1, closed)
			stack = stack[:len(stack)-1]
		}

		if closed < open {
			stack = append(stack, ')')
			dfsBT(open, closed+1)
			stack = stack[:len(stack)-1]
		}
	}

	dfsBT(0, 0)
	return res
}
