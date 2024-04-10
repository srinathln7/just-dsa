package main

import (
	"strconv"
	"strings"
)

type Solution struct{}

// Function to encode list of strings to a single string
func (s *Solution) Encode(strs []string) string {

	// Key Idea: Since we cannot use any extra space or data structure for encoding
	// let us encode the number of characters in each string array along with the delimited (#)
	// with the string itself. This will help us determine the length of each characters when
	// string is decoded back

	var builder strings.Builder
	for _, str := range strs {
		builder.WriteString(strconv.Itoa(len(str)))
		builder.WriteByte('#')
		builder.WriteString(str)
	}
	return builder.String()
}

// Function to decode the encoded string back to the list of strings
func (s *Solution) Decode(sEncoded string) []string {

	var result []string
	i := 0
	for i < len(sEncoded) {
		j := i

		// All characters before the delimiter represent the length
		for sEncoded[j] != '#' {
			j++
		}

		length, _ := strconv.Atoi(sEncoded[i:j])
		i = j + 1
		j = i + length
		result = append(result, sEncoded[i:j])
		i = j
	}

	return result
}
