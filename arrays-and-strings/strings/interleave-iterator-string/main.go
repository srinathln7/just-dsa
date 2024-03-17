package main

import "strings"

func printVertically(s string) []string {

	// Key Idea: To convert the given string into a string array seperated by spaces. The length of the
	// resulting array will be the length of the longest word in the formed string array

	strArr := strings.Split(s, " ")

	// Determine max length of the result array
	maxLength := 0
	for _, str := range strArr {
		if len(str) > maxLength {
			maxLength = len(str)
		}
	}

	// Pad whitespaces to the right i.e. trailing white spaces
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
