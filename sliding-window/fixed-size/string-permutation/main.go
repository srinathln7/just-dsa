package main

func checkInclusion(s1 string, s2 string) bool {

	// Fixed sliding window
	n1 := len(s1)
	n2 := len(s2)

	if n2 < n1 {
		return false
	}

	subStr := s2[:n1]
	if isPermutation(s1, subStr) {
		return true
	}

	for i := n1; i < n2; i++ {
		subStr = s2[i-n1+1 : i+1]
		if isPermutation(s1, subStr) {
			return true
		}
	}

	return false
}

func isPermutation(s1, s2 string) bool {
	var (
		count [26]int
	)

	for _, ch := range s1 {
		count[ch-'a']++
	}

	for _, ch := range s2 {
		count[ch-'a']--
	}

	for i := 0; i < len(count); i++ {
		if count[i] != 0 {
			return false
		}
	}

	return true
}
