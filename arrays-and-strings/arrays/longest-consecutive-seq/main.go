package main

func longestConsecutive(nums []int) int {
	// Key Idea: We know that the start of the sequence has no left neighbour.
	// Every time, we encounter a start of the sequence, we check if the next consequective element exists or not
	seqMap := make(map[int]bool)
	for _, num := range nums {
		seqMap[num] = true
	}

	var maxLen int
	for _, num := range nums {
		// If no left seq exists, then check if a right seq can be formed
		if !seqMap[num-1] {
			seqLen := 1
			for seqMap[num+1] {
				seqLen++
				num++
			}
			maxLen = max(maxLen, seqLen)
		}
	}

	return maxLen
}
