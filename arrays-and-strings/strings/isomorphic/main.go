package main

func isIsomorphic(s string, t string) bool {
	// Key Idea: Since we need to preserve one-to-one mapping i.e. no two characters can map to the same character set
	// we can maintain two hashmaps and check if the string is isomorphic or not

	n := len(s)
	charMapSet := make(map[byte]byte)
	revCharMapSet := make(map[byte]byte)

	for i := 0; i < n; i++ {

		// Check for one-to-one mapping in charMapSet
		if val, exists := charMapSet[s[i]]; exists && val != t[i] {
			return false
		}

		// Check for one-to-one mapping in recCharMapSet
		if val, exists := revCharMapSet[t[i]]; exists && val != s[i] {
			return false
		}

		charMapSet[s[i]] = t[i]
		revCharMapSet[t[i]] = s[i]
	}

	return true
}
