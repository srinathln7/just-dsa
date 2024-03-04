package main

func rotateString(s string, goal string) bool {
	if len(s) > len(goal) {
		return false
	}

	return isSubString(s, goal+goal)
}

func isSubString(shortStr, longStr string) bool {
	if len(shortStr) > len(longStr) {
		return false
	}

	if shortStr == longStr {
		return true
	}

	m := len(shortStr)
	n := len(longStr)

	// Fixed Window
	subStr := longStr[0:m]
	if shortStr == subStr {
		return true
	}

	for i := m; i < n; i++ {
		subStr = longStr[i-m+1 : i+1]
		if shortStr == subStr {
			return true
		}
	}

	return false
}
