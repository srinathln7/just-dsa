package main

import "strings"

func wordPattern(pattern string, s string) bool {

	chMap := make(map[byte]string)
	strMap := make(map[string]byte)

	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}

	n := len(pattern)
	for i := 0; i < n; i++ {

		if word, exists := chMap[pattern[i]]; exists && word != words[i] {
			return false
		}

		if ch, exists := strMap[words[i]]; exists && ch != pattern[i] {
			return false
		}

		chMap[pattern[i]] = words[i]
		strMap[words[i]] = pattern[i]
	}

	return true
}
