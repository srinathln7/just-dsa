package main

import "strconv"

func addBinary(a string, b string) string {
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

		// At every iteration: sum ranges from -> 0, 1, 2
		result = strconv.Itoa(sum%2) + result
		carry = sum / 2
	}

	// Any remaining carry after the last iteration
	if carry > 0 {
		result = strconv.Itoa(carry) + result
	}

	return result
}
