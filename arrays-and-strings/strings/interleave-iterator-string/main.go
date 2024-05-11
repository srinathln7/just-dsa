package main

import "strings"

func printVertically(s string) []string {

	// Key Idea: Form the string array and then start interleaving elements. The results length will
	// be equal to the longest word length where each word is seperated by a space

	strArr := strings.Split(s, " ")

	// Determine the max. length
	maxlen := 0
	for _, word := range strArr {
		maxlen = max(maxlen, len(word))
	}

	// Example:
	// [["HOWARD" , "ARE", "YOU"]]
	// [["HAY"], ["ORO"], ["WEU"], ["A"], ["R"], ["D"] ]
	result := make([]string, maxlen)
	for i := 0; i < maxlen; i++ {
		for _, word := range strArr {

			// Make all words len the same before forming the result to avoid out-of-bound exception
			// This can be achieved easily by padding extra spaces to the right
			if len(word) < maxlen {
				word += strings.Repeat(" ", maxlen-len(word))
			}

			result[i] += string(word[i])
		}
	}

	// Remove the trailing spaces
	for i, word := range result {
		result[i] = strings.TrimRight(word, " ")
	}

	return result
}
