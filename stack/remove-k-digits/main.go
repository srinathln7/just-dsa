package main

func removeKdigits(num string, k int) string {

	// Key Idea: We use monotonic stack technique to solve this problem where we try to maintain the stack
	// in increasing order. If this is the case, then we can stack removing the digits from the right.
	// If the numbers are in decreasing order, then we do the opposite by starting to remove from the left.
	// For ex: 12345, k =3 => smallest number is 12 removing from the right side
	// For ex: 54321, k=3 => smallest number is 21 removing from the left side
	// Therefore if we can maintain a monotonic stack of increasing order we can start poping the element as long
	// as k is zero

	// Base case: If the string length is lesser than or equal to k, then we can remove all the characters
	n := len(num)
	if n <= k {
		return "0"
	}

	var stack []rune
	for _, ch := range num {

		// To get the minimum number: for a non-empty stack and if the incoming character is
		//lesser than the stack's top then stack is no longer in increasing order. Hence we need to pop
		// To get maximum number, do the opposite
		for k > 0 && len(stack) > 0 && ch < stack[len(stack)-1] {

			// Pop from the stack
			stack = stack[:len(stack)-1]
			k--
		}

		// No use appending leading zeros since they dont add any value
		if len(stack) > 0 || ch != '0' {
			stack = append(stack, ch)
		}
	}

	// Since now the stack is is increasing order, if k is still non-zero
	// then we have to remove the remaining digits from the MSB until either the
	// stack is empty is k is zero
	for len(stack) > 0 && k > 0 {
		stack = stack[:len(stack)-1]
		k--
	}

	// Trim leading zeros - This is not needed since we are not appending any leading zeros
	//result := strings.TrimLeft(string(stack), "0")
	//result, idx := string(stack), 0
	// for _, ch := range result {
	//     if ch == '0' {
	//         idx++
	//     } else {
	//         break
	//     }
	// }
	// result = result[idx:]

	result := string(stack)
	if result == "" {
		return "0"
	}

	return result
}
