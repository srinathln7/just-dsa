package main

func letterCombinations(digits string) []string {

	// Key Idea: Use a map to map the corresponding letter to the respective numbers
	// Run a DFS with BT algorith to generate all possible combinations

	n := len(digits)
	if n == 0 {
		return nil
	}

	var result []string
	phoneMap := make(map[byte]string)
	phoneMap['2'] = "abc"
	phoneMap['3'] = "def"
	phoneMap['4'] = "ghi"
	phoneMap['5'] = "jkl"
	phoneMap['6'] = "mno"
	phoneMap['7'] = "pqrs"
	phoneMap['8'] = "tuv"
	phoneMap['9'] = "wxyz"

	dfsWithBT(phoneMap, digits, "", 0, &result)
	return result
}

func dfsWithBT(phoneMap map[byte]string, digits string, comb string, idx int, result *[]string) {
	if len(comb) == len(digits) {
		*result = append(*result, comb)
		return
	}

	for _, char := range phoneMap[digits[idx]] {
		// Append
		comb += string(char)

		// Recurse
		dfsWithBT(phoneMap, digits, comb, idx+1, result)

		// Backtrack
		comb = comb[:len(comb)-1]

	}
}
