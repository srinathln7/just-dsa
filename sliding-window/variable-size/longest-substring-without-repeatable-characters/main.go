package main

func lengthOfLongestSubstring(s string) int {

	// Valid window mapping the character in each string to its respective index
	window := make(map[byte]int)

	var left, length int
	for r := 0; r < len(s); r++ {
		if index, exists := window[s[r]]; exists && index >= left {
			// Index represents the repeated character's index
			// Whenever a repeated character occurs move your left pointer to the next occurence of
			// the index to keep the window valid. NOTE: your left pointer should always trail behind or
			// be equal to the index and never greater.
			left = index + 1
		}

		// Create/Update the characters index
		window[s[r]] = r
		length = max(length, r-left+1)
	}

	return length
}
