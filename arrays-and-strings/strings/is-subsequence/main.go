package main

func isSubsequence(s string, t string) bool {

	// Key Idea: Use a two pointer approach
	m, n := len(s), len(t)
	if m > n {
		return false
	}

	if s == t || s == "" {
		return true
	}

	idx := 0
	for i := 0; i < len(t); i++ {
		// Cannot iterate any further
		if idx == len(s) {
			break
		}

		// Count the number of common characters ensuring the same relative order with the use of the loop
		if s[idx] == t[i] {
			idx++
		}
	}

	return idx == len(s)
}
