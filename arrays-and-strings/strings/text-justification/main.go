package main

import "strings"

func fullJustify(words []string, maxWidth int) []string {

	// Key Idea: Distribute spaces evenly with bias for left justification based on cyclic modulo distribution
	// At any given point in time, keep track of currentLine (of type `string`) and currentLength (representing
	// the length of the non-space characters) and modify accordingly

	var result []string
	var currLine []string
	var currLen int

	for _, word := range words {

		// currLen: Total length of all words currently in the line.
		// len(currLine): Represents the minimum number of spaces needed between the words (one less than the number of words).
		// len(word): Length of the current word we are trying to add.

		if currLen+len(currLine)+len(word) > maxWidth {
			// We have to distribute a total of `maxWidth-currLen` spaces b/w `len(currLine)` words
			// i.e. We can atmost have `len(currLine)-1` space in-total.
			for i := 0; i < maxWidth-currLen; i++ {
				currLine[i%max(1, len(currLine)-1)] += " "
			}
			result = append(result, strings.Join(currLine, ""))

			// Reset after the end of every line
			currLine, currLen = []string{}, 0
		}

		currLine = append(currLine, word) // Rep. the number of space characters req'd
		currLen += len(word)              // Rep. length of non-space characters
	}

	// Account for last line. This line constructs the last line of the text to be left-justified
	// and ensures that the total length of the line equals maxWidth by adding the necessary padding
	// spaces on the right. `len(currLine)-1` represents the number of spaces that are already intoduced by
	// `strings.Join` function.
	result = append(result, strings.Join(currLine, " ")+strings.Repeat(" ", maxWidth-currLen-(len(currLine)-1)))
	return result
}
