package main

func missingNumber(nums []int) int {
	// Key Idea: We can use properties of XOR to perform this operation.
	// XOR is commutative and associative and XOR of a number with itself gives zero.
	// XOR of a number with zero gives the number

	var missing int
	for i, num := range nums {
		missing ^= num ^ (i + 1)
	}
	return missing
}
