package main

func getSum(a int, b int) int {

	// Key Idea: After doing a set of simple example problems, we observe that the actual sum is formed by XORing the digits
	// and the carry is obtained by BITWISE AND ops and left shifting it by one. We continue this process until the carry is 0

	var sum, carry int
	for b != 0 {
		// Calculate sum without considering carry
		sum = a ^ b

		// Calculate carry (bits that need to be added to the next position)
		carry = (a & b) << 1

		// Update a to sum and b to carry shifted by 1
		// Repeat the process until there is no carry
		a = sum
		b = carry
	}

	return a
}
