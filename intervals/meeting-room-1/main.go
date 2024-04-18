package main

import "sort"

func canAttendMeetings(intervals [][]int) bool {

	// Key Idea: Sort the intervals as per starting time and check for overlaps
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][0] {
			return false
		}
	}

	return true
}
