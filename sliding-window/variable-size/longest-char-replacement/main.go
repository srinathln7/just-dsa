package main

func characterReplacement(s string, k int) int {

	// KEY IDEA: For a given window, we need to keep track of the character with the most frequency
	// Then the number of operations required to make other characters similar in that window can be calculated by
	// #operations_reqd = #window_size - #max_repeated_character And this operation should never exceed `k`

	var maxCount, maxLength int
	countWindow := make(map[byte]int)

	l := 0
	for r := 0; r < len(s); r++ {
		countWindow[s[r]]++

		// maxCount represents the maximum number of repeated characters in a window
		maxCount = max(maxCount, countWindow[s[r]])

		// Number of operations required to make all the characters in the window the same and
		// if that exceeds `k` shrink the given window
		num_of_ops := (r - l + 1) - maxCount
		if num_of_ops > k {
			countWindow[s[l]]--
			l++
		}

		maxLength = max(maxLength, r-l+1)
	}

	return maxLength
}
