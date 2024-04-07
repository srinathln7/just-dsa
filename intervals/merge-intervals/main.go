package main

import "sort"

func merge(intervals [][]int) [][]int {

	// Key Idea: Sort the interval based on the start point of the intervals.
	// Form a new slice and keep adding the elements when there is no overlap to this new slice.
	// If there is a overlap, update the end point of the merged slice DYNAMICALLY.

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})

	var merged [][]int
	for _, interval := range intervals {
		// No overlap
		if len(merged) == 0 || interval[0] > merged[len(merged)-1][1] {
			merged = append(merged, interval)
		} else {
			// If there is a overlap, update the end point of the interval dynamically
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], interval[1])
		}
	}

	return merged
}
