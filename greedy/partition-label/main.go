package main

func partitionLabels(s string) []int {

	// Key Idea: Track the last occuring index of every character in a set and keep partioning accordingly

	lastIdxSet := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		lastIdxSet[s[i]] = i
	}

	var result []int

	left, lastIdx := 0, 0

	// Starting partioning
	for right := 0; right < len(s); right++ {
		currCharLastIdx := lastIdxSet[s[right]]
		if currCharLastIdx > lastIdx {
			lastIdx = currCharLastIdx
		}

		if lastIdx == right {
			result = append(result, right-left+1)
			left = right + 1
		}
	}

	return result
}
