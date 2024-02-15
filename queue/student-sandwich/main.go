package main

func countStudents(students []int, sandwiches []int) int {
	var count [2]int
	n := len(students)

	// Calculate how many students want circular and square sandwiches
	for i := 0; i < n; i++ {
		count[students[i]]++
	}

	// Range over sandwich and decrement the count as students take sandwich
	for j := 0; j < n; j++ {
		// No further sandwich left
		if count[sandwiches[j]] == 0 {
			break
		}
		count[sandwiches[j]]--
	}

	return count[0] + count[1]
}
