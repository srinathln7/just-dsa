package main

func findAnagrams(s string, p string) []int {

	// Key Idea: Since both strings `s` and `p` only contains lower-case english letters, we can declare a count array
	// for each of them and calculate its respective frequencies in the respective windows

	n, k := len(s), len(p)
	if k > n {
		return nil
	}

	countS, countP := [26]int{}, [26]int{}

	var result []int

	//First window: Build the count for string `p` and the count of string 's' in the first window
	for i := 0; i < k; i++ {
		countP[p[i]-'a']++
		countS[s[i]-'a']++
	}

	if isEqual(countS, countP) {
		result = append(result, 0)
	}

	// Slide the window
	for i := k; i < n; i++ {

		// Deboarding character with index `i-k` from string `s` in the current window
		countS[s[i-k]-'a']--

		// Onboarding character with index `i` from string `s` in the current window
		countS[s[i]-'a']++

		if isEqual(countS, countP) {
			result = append(result, i-k+1)
		}
	}

	return result
}

func isEqual(a, b [26]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
