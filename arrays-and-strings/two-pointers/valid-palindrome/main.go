package main

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {

	f := func(ch rune) rune {
		if !unicode.IsLetter(ch) && !unicode.IsNumber(ch) {
			return -1
		}
		return unicode.ToLower(ch)
	}

	// Map returns a copy of the string s with all its characters modified according to the mapping function.
	// If mapping returns a negative value, the character is dropped from the string with no replacement.
	s = strings.Map(f, s)
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}
