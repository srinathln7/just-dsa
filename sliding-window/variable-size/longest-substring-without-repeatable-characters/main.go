package main

func lengthOfLongestSubstring(s string) int {

	// Valid window mapping the character in each string to its respective index
	window := make(map[byte]int)

	var left, length int
	for r := 0; r < len(s); r++ {
		if indexRepeat, exists := window[s[r]]; exists && indexRepeat >= left {
			// Index represents the repeated character's index
			// Whenever a repeated character occurs move your left pointer to the next occurence of
			// the index to keep the window valid.
			// NOTE: once you have updated your left pointer, the next repeated index should always be greater than or equal
			// the current left index. Eg: Think of string "abba" with max non-repeatable character length = 2
			left = indexRepeat + 1
		}

		// Create/Update the characters index
		window[s[r]] = r
		length = max(length, r-left+1)
	}

	return length
}
