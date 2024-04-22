package main

func dailyTemperatures(temperatures []int) []int {

	// Key Idea: We start from reverse and we observe that the if we start from reverse, the first element to be pushed
	// to the stack will always be zero. This is because there is no element to the right of it and hence there can be no days
	// warmer that that. The idea is to traverse the temperatures array from right to left and maintain a stack that keeps
	// track of indices of elements whose temperatures we have not yet found a warmer day for. Whenever we encounter a temperature
	// higher than the temperature of the current element, we pop indices from the stack and calculate the number of days to wait for
	// a warmer temperature.

	n := len(temperatures)
	if n == 1 {
		return []int{0}
	}

	stack := make([]int, 0)
	result := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		currTemp := temperatures[i]

		// If the stack is non-empty and we have encountered a temperature that is warmer than all the subsequent days
		// pop from the stack. This is because for all the other days going from right to left, we have now found the
		// warmest day.
		for len(stack) > 0 && currTemp >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		// If the stack is not empty - result is captured by the subracting the day at which
		// at you which you will encounter the warming day minus your current day => This will
		// give the number of days we have to wait for
		if len(stack) > 0 {
			result[i] = stack[len(stack)-1] - i
		}

		// Append the index to the stack at every iteration
		stack = append(stack, i)
	}

	return result
}
