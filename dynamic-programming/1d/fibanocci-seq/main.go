package main

func fib(n int) int {
	current, next := 0, 1
	for i := 0; i < n; i++ {
		current, next = next, current+next
	}

	return current
}
