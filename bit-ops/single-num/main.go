package main

func singleNumber(nums []int) int {
	// Key Idea: Use the properties of XOR
	// XORing any number with itself results in 0.
	// XORing any number with 0 results in the number itself.
	// XORing is commutative and associative, meaning the order of operands does not matter.

	var result int
	for _, num := range nums {
		result ^= num
	}

	return result
}
