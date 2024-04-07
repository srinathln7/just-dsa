package main

func insert(intervals [][]int, newInterval []int) [][]int {

	// Key Idea: To first add all the intervals that come before the new interval (curr_interval_end < new_start)
	// Merge the overlapping intervals and then add the remaining intervals.
	// New intervals start will be min of the start value and end will be the max of the end value.

	var result [][]int
	i, n := 0, len(intervals)
	newIntervalStart, newIntervalEnd := newInterval[0], newInterval[1]

	// Add all intervals before the `newInterval`
	for i < n && intervals[i][1] < newIntervalStart {
		result = append(result, intervals[i])
		i++
	}

	// Merge the overlapping intervals
	for i < n && intervals[i][0] <= newIntervalEnd {
		newIntervalStart = min(intervals[i][0], newIntervalStart)
		newIntervalEnd = max(intervals[i][1], newIntervalEnd)
		i++
	}
	result = append(result, []int{newIntervalStart, newIntervalEnd})

	// Add the remaining intervals
	for i < n {
		result = append(result, intervals[i])
		i++
	}

	return result
}
