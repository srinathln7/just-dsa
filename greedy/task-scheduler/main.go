package main

import "sort"

func leastInterval(tasks []byte, n int) int {

	// Key Idea: We calculate the freq of each tasks and sort them in descending order. This tells us about the max. freq of each task.
	// We start with the worst case sceanario - we calculate the total number of idle slots required to space these tasks apart. After
	// that, we try to optimise the total number of idle slots by subrating it based on the min(task_freq, maxfreq-1). It is possible that
	// the number of idle slots becomes negative (when we have more tasks than avaiable idle slots) and in that case, we reset it to zero
	// since negative values has no meaning in this context. Finally, we caulcalte the sum of all tasks and idle slots

	// Build the freq set
	freqSet := make(map[byte]int)
	for _, task := range tasks {
		freqSet[task]++
	}

	// Sort the freq set in decreasing order
	sorted, idx := make([]int, len(freqSet)), 0
	for _, freq := range freqSet {
		sorted[idx] = freq
		idx++
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] > sorted[j]
	})

	maxFreq := sorted[0]

	// Worst case sceanario - Tasks with max. freq are spaced `n` intervals apart
	idleSlots := (maxFreq - 1) * n

	// Optimize (Minimize) the number of idle slots required
	for i := 1; i < len(sorted); i++ {
		idleSlots -= min(sorted[i], maxFreq-1)
	}

	// Reset idle slots to zero if negative. Negative slots occur when there are more tasks than idle slots which
	// essentially means that we can run all tasks in the given order satisfying the cool down constraint without any
	// idle slots
	idleSlots = max(0, idleSlots)

	// As any way - we need to perform all tasks + any non-negative idle slots
	return len(tasks) + idleSlots
}
