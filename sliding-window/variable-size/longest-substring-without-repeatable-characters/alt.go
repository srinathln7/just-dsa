package main

import "math"

func altlengthOfLongestSubstring(s string) int {
	// longest unique substring => sliding window of variable size
	// "abcabcbb"

	if len(s) == 0 {
		return 0
	}

	var length int = math.MinInt32
	var subStr string

	l := 0
	for r, ch := range s {
		// Keep accumulating
		subStr += string(ch)
		// Maximize the window
		for !isUniqueString(subStr) {
			length = max(length, r-l)
			subStr = subStr[1:]
			l++
		}
	}

	if length == math.MinInt32 {
		return len(s)
	}

	return max(length, len(subStr))
}

func isUniqueString(s string) bool {
	strMap := make(map[string]struct{})
	for _, ch := range s {
		if _, exists := strMap[string(ch)]; exists {
			return false
		}

		strMap[string(ch)] = struct{}{}
	}

	return true
}
