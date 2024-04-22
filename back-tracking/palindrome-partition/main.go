package main

import "fmt"

func partition(s string) [][]string {

	// Key Idea: We run a simple DFS with back tracking algorithm to output all possible partitions

	var result [][]string

	//  IMPORTANT: To declare palindrome here as whatever slice we are backtracking needs to be outside the recursive function
	var palindrome []string

	var dfs func(idx int)
	dfs = func(startIdx int) {

		if startIdx >= len(s) {
			result = append(result, append([]string{}, palindrome...))
			return
		}

		for endIdx := startIdx; endIdx < len(s); endIdx++ {
			if isPalindrome(s, startIdx, endIdx) {
				// Append the palindrome
				palindrome = append(palindrome, s[startIdx:endIdx+1])

				fmt.Println("palindrome", palindrome)
				// Recurse
				dfs(endIdx + 1)

				// Backtrack
				palindrome = palindrome[:len(palindrome)-1]
			}
		}
	}

	// Kickstart
	dfs(0)

	return result
}

func isPalindrome(s string, l, r int) bool {

	for l < r {
		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}
