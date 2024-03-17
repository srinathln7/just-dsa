package main

import (
	"strings"
	"unicode"
)

func PrintVertically(s string) []string {
	// Convert given string to string array
	strArray := strings.Split(s, " ")
	maxLength := 0
	for _, word := range strArray {
		if len(word) > maxLength {
			maxLength = len(word)
		}
	}

	// Pad the string to account for spaces
	// It is important to pad the spaces to the right since trailing spaces can later be removed
	for i, word := range strArray {
		strArray[i] = word + strings.Repeat(" ", maxLength-len(word))
	}

	byteArray := make([][]byte, len(strArray))
	for i, str := range strArray {
		byteArray[i] = []byte(str)
	}

	resultByteArray := interLeave(byteArray)

	// Convert back to string array
	var result []string
	var resultStr string
	for _, bytes := range resultByteArray {
		resultStr = strings.TrimRightFunc(string(bytes), unicode.IsSpace)
		result = append(result, resultStr)
	}

	return result
}

type Iterator struct {
	lists []byte
	index int
}

func newIterator(lists []byte) *Iterator {
	return &Iterator{
		lists: lists,
	}
}

func (i *Iterator) hasNext() (byte, bool) {
	if i.index >= len(i.lists) {
		return 32, false
	}

	val := i.lists[i.index]
	i.index++
	return val, true
}

func interLeave(lists [][]byte) [][]byte {
	var results [][]byte
	iterators := make([]*Iterator, len(lists)) // Key idea is to make this a POINTER to the iterator struct

	// Key Idea: is to loop one row at a time and within each row we loop one char at a time
	// To do this we create a wrapper around iterator for each row in the lists
	for i, list := range lists {
		iterators[i] = newIterator(list)
	}

	for {
		var elements []byte // Slice containg interleaved elements for the current iteration
		var anyRemaining bool

		for _, iterator := range iterators {
			val, isPresent := iterator.hasNext()
			if isPresent {
				elements = append(elements, val)
				anyRemaining = true
			}
		}

		if !anyRemaining {
			break
		}
		results = append(results, elements)
	}
	return results
}
