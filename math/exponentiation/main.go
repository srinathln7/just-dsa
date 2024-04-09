package main

func myPow(x float64, n int) float64 {

	// Key Idea: Straight-forward exponentiation would not work for larger numbers
	// To accomodate large numbers, we need to implement binary exponentiation which implements exponentiation in log(n) time.
	// To do this, we square the number and divide `n` by 2 at every iteration. We continue this process until `n` is zero

	if x == 0.0 {
		return 0.0
	}

	if n == 0 || x == 1.0 {
		return 1.0
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	result := 1.0
	for n > 0 {

		// If `n` is odd number
		// As we keep dividing `n` by 2, eventually `n` will become 1
		// This will accumulate the result variable
		if n&1 == 1 {
			result *= x
		}

		// Square the number
		x *= x

		// Divide `n` by 2
		n >>= 1
	}

	return result
}
