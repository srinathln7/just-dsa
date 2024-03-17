package main

func longestPalindrome(s string) string {

	// Key Idea: Traditionally, to find if a given string is palidrome or not, we start at the extreme ends and expand inwards.
	// To find the length of the longest palindromic substring, we adopt a reverse approach where we pick every character in the
	// string and start expanding outwards until we don't go out-of-bound. Since picking every character and starting to expand
	// outwards would always result in odd number of characters, we also do the same for every pair of contagious characters
	// so that we also check for even length

	var substr, longest string
	n := len(s)
	for i := 0; i < n; i++ {
		// Odd length
		l, r := i, i
		substr = getPalindromicSubStr(s, l, r)
		if len(substr) > len(longest) {
			longest = substr
		}

		// Even length
		l, r = i, i+1
		substr = getPalindromicSubStr(s, l, r)
		if len(substr) > len(longest) {
			longest = substr
		}
	}

	return longest
}

func getPalindromicSubStr(s string, l, r int) string {
	var substr string
	n := len(s)
	for l >= 0 && r < n && s[l] == s[r] {
		substr = s[l : r+1]
		l--
		r++
	}

	return substr
}
