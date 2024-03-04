package main

import "strconv"

func compress(chars []byte) int {
	var count, resultIndex int
	var countStr string

	for i := 0; i < len(chars); i++ {
		count++

		// Only start writing when the adjacent characters are not equal to each other
		// Special case at the end of the string
		if i+1 == len(chars) || chars[i] != chars[i+1] {
			chars[resultIndex] = chars[i]
			resultIndex++

			if count > 1 {
				countStr = strconv.Itoa(count)
				for j := 0; j < len(countStr); j++ {
					chars[resultIndex] = countStr[j]
					resultIndex++
				}
			}
			// Reset count
			count = 0
		}
	}

	return resultIndex
}
