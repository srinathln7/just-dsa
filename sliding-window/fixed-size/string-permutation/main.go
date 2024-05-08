package main

func checkInclusion(s1 string, s2 string) bool {

	// Key Idea: To check if s2 contains a permutation of s1, we can use a sliding window approach of fixed size k=len(s1).

	k, n := len(s1), len(s2)
	if k > n {
		return false
	}

	// Since both strings only contains lower case English letters
	countS1, countS2 := [26]int{}, [26]int{}
	for i := 0; i < k; i++ {
		countS1[s1[i]-'a']++
		countS2[s2[i]-'a']++
	}

	if isEqual(countS1, countS2) {
		return true
	}

	// Slide the window of fixed size
	for i := k; i < n; i++ {
		// Deboarding characters
		countS2[s2[i-k]-'a']--

		// Onboarding characters
		countS2[s2[i]-'a']++

		if isEqual(countS1, countS2) {
			return true
		}
	}

	return false
}

func isEqual(num1, num2 [26]int) bool {
	for i := 0; i < 26; i++ {
		if num1[i] != num2[i] {
			return false
		}
	}

	return true
}
