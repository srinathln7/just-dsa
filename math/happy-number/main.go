package main

func isHappy(n int) bool {
	// Define a function to compute the next number
	nextNumber := func(num int) int {
		sum := 0
		for num > 0 {
			digit := num % 10
			sum += digit * digit
			num /= 10
		}
		return sum
	}

	// Initialize slow and fast pointers to n
	slow, fast := n, n

	// Use Floyd's cycle-finding algorithm
	for {
		slow = nextNumber(slow)             // Move slow one step forward
		fast = nextNumber(nextNumber(fast)) // Move fast two steps forward

		// If slow reaches 1, n is a happy number
		if slow == 1 {
			return true
		}

		// If slow equals fast, n is not a happy number (cycle detected)
		if slow == fast {
			return false
		}
	}
}
