package main

func plusOne(digits []int) []int {
	n := len(digits)

	// Iterate from the least significant digit to the most significant digit
	for i := n - 1; i >= 0; i-- {
		// Increment the current digit by 1
		digits[i]++

		// If the digit becomes 10, set it to 0 and propagate the carry
		if digits[i] == 10 {
			digits[i] = 0
		} else {
			// If the digit is less than 10, no need to propagate carry, return the result
			return digits
		}
	}

	// If the carry propagates to the most significant digit, append a new digit (1) to the beginning
	return append([]int{1}, digits...)
}
