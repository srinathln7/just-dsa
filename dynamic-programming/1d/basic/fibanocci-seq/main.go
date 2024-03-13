package main

func fib(n int) int {
	current := 0
	next := 1

	if n == 0 {
		return current
	}

	if n == 1 {
		return next
	}

	for i := 2; i <= n; i++ {
		current, next = next, current+next
	}

	return next
}
