package main

import "math"

func minWindow(s string, t string) string {

	// Key Idea: Sliding window of variable size. Two pointer approach `l`  and `r`. Move `r` until we first encounter a valid
	// window where all chars in `t` are contained in that window. Then, lets adopt a greedy approach and start shrinking the
	// window by moving the left pointer to find the minimum valid window

	// Base cases:

	// When the first string has fewer chars than the second one, it is impossible for `s` to conatin all chars in `t`
	if len(s) < len(t) {
		return ""
	}

	// When two strings are equal that is the min. substring
	if s == t {
		return s
	}

	// Build a freq. table for string `t`
	freq := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		freq[t[i]]++
	}

	// Check for the valid window
	var minStart, matchCount int
	var minLen int = math.MaxInt32

	l := 0
	for r := 0; r < len(s); {
		// Expand the window
		freq[s[r]]--
		if freq[s[r]] >= 0 {
			matchCount++
		}
		r++

		// GREEDY: Try to shrink the window as much as possible
		for matchCount == len(t) {
			if r-l < minLen {
				minLen = r - l
				minStart = l
			}
			freq[s[l]]++
			if freq[s[l]] > 0 {
				matchCount--
			}
			l++
		}
	}

	// If despite running the above algo. minLen has not changed, it means, there is no matching string
	if minLen == math.MaxInt32 {
		return ""
	}

	return s[minStart : minStart+minLen]
}
