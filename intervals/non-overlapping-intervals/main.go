package main

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {

	// Key Idea: Sort the intervals according to their endpoints in ascending order.
	// Keep track of prevEnd and currEnd. When there is a overlap increment the count else
	// update the curr end to prev end

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] <= intervals[j][1]
	})

	var count int

	n, prevEnd := len(intervals), intervals[0][1]
	for i := 1; i < n; i++ {
		currStart, currEnd := intervals[i][0], intervals[i][1]
		// Overlap - increment the count
		if currStart < prevEnd {
			count++
		} else {
			// No overlap - update the prevEnd to currEnd and keep traversing
			prevEnd = currEnd
		}
	}

	return count
}
