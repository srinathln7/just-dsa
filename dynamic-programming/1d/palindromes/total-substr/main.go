package main

func countSubstrings(s string) int {

	// Key Idea: Traditionally, to find if a given string is palidrome or not, we start at the extreme ends and expand inwards.
	// To find the length of the longest palindromic substring, we adopt a reverse approach where we pick every character in the
	// string and start expanding outwards until we don't go out-of-bound. Since picking every character and starting to expand
	// outwards would always result in odd number of characters, we also do the same for every pair of contagious characters
	// so that we also check for even length

	var count int
	n := len(s)
	for i := 0; i < n; i++ {
		// Odd length
		l, r := i, i
		getPalindromicSubStr(s, l, r, &count)

		// Even length
		l, r = i, i+1
		getPalindromicSubStr(s, l, r, &count)
	}

	return count
}

func getPalindromicSubStr(s string, l, r int, count *int) {
	n := len(s)
	for l >= 0 && r < n && s[l] == s[r] {
		l--
		r++
		*count++
	}
}
