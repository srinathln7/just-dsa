package main

import "strings"

func lengthOfLastWord(s string) int {

	strArr := strings.Split(s, " ")
	n := len(strArr)

	i := n - 1
	for i >= 0 && len(strArr[i]) == 0 {
		i--
	}

	return len(strArr[i])
}
