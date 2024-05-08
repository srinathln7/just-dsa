package main

func findRepeatedDnaSequences(s string) []string {

	// Key Idea: Sliding window of fixed size = 10

	n := len(s)
	if n < 10 {
		return nil
	}

	var result []string
	dnaSet := make(map[string]bool)
	resultSet := make(map[string]bool)

	// First window
	dnaSet[s[:10]] = true

	// Slide the window of fixed size=10
	for i := 10; i < n; i++ {
		if dnaSet[s[i-10+1:i+1]] && !resultSet[s[i-10+1:i+1]] {
			resultSet[s[i-10+1:i+1]] = true
			result = append(result, s[i-10+1:i+1])
		} else {
			dnaSet[s[i-10+1:i+1]] = true
		}
	}

	return result
}
