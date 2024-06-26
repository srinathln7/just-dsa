package main

import "strconv"

func addBinary(a string, b string) string {

	// Key Idea: Start adding digit-by-digit from the LSB to the MSB. When we add digit-by-digit
	// sum ranges from 0,1, and 2. We can get the last binary digit by `mod 2` and any carry
	// by diving the number by 2. Prepend the resulting value with the `result` string after every iteration.
	// Add from the LSB to MSB in binary itself. Sum can range from [0,1,2]. You can obtain the
	// result value at each iteration by taking `mod 2` and the carry by `dividing` by 2.
	// IMPORTANT To postpend the append of result after every iteration and to take final carry into account.

	// Base cases
	if a == "0" {
		return b
	}

	if b == "0" {
		return a
	}

	var result string

	// Pointers to the least-significant-digit
	m, n := len(a), len(b)
	i, j := m-1, n-1

	carry := 0
	for i >= 0 || j >= 0 {
		sum := carry

		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}

		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		// sum -> 0, 1, 2
		result = strconv.Itoa(sum%2) + result
		carry = sum / 2
	}

	if carry > 0 {
		result = strconv.Itoa(carry) + result
	}

	return result
}
