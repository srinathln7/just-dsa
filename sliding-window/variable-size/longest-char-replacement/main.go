package main

func characterReplacement(s string, k int) int {

	// KEY IDEA: We init two pointers `l` and `r` and keep moving `r` pointer forward.
	// As move along, we need to keep track of the character's freq (char with the most frequency can be inferred from this)
	// Then the number of operations required to make other characters similar in that window can be calculated by
	// #operations_reqd = #window_size - #max_repeated_character. If this operation exceeds `k`, we start shrinking the window
	// from the left since oly atmost `k` operations are allowed.

	var maxCount, maxLength int
	countWindow := make(map[byte]int)

	l := 0
	for r := 0; r < len(s); r++ {
		countWindow[s[r]]++

		// maxCount represents the maximum number of repeated characters in a window
		maxCount = max(maxCount, countWindow[s[r]])

		// Number of operations required to make all the characters in the window the same and
		// if that exceeds `k` start shrinking the window from the left
		num_of_ops := (r - l + 1) - maxCount
		if num_of_ops > k {
			countWindow[s[l]]--
			l++
		}

		maxLength = max(maxLength, r-l+1)
	}

	return maxLength
}
