package main

import (
	"strings"
)

func printvertically(s string) []string {
	strArr := strings.Split(s, " ")

	// Determine max length
	maxLength := 0
	for _, str := range strArr {
		if len(str) > maxLength {
			maxLength = len(str)
		}
	}

	// Pad whitespaces
	for i, str := range strArr {
		if len(str) < maxLength {
			strArr[i] = str + strings.Repeat(" ", maxLength-len(str))
		}
	}
	// [["HOWARD" , "ARE", "YOU"]]
	// [["HAY"], ["ORO"], ["WEU"], ["A"], ["R"], ["D"] ]
	result := make([]string, maxLength)
	for i := 0; i < maxLength; i++ {
		for _, word := range strArr {
			result[i] += string(word[i])
		}
	}

	// Trim the trailing right spaces
	for i := 0; i < maxLength; i++ {
		result[i] = strings.TrimRight(result[i], " ")
	}

	return result
}
