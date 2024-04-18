package main

import (
	"fmt"
	"sort"
)

// Definition of Interval
type Interval struct {
	Start, End int
}

func MinMeetingRooms(intervals []*Interval) int {
	// Write your code here

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].End < intervals[j].End
	})

	// Remove duplicates
	idx := 1
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start == intervals[i-1].Start && intervals[i].End == intervals[i-1].End {
			continue
		}
		intervals[idx] = intervals[i]
		idx++
	}
	intervals = intervals[:idx]

	// Min number of intervals we need to remove to make all non-overlapping
	var count int
	prevEnd := intervals[0].End
	for i := 1; i < len(intervals); i++ {
		currStart, currEnd := intervals[i].Start, intervals[i].End
		// Overlap
		if currStart < prevEnd {
			count++
		} else {
			prevEnd = currEnd
		}
	}

	// Count represents the total number of overlapping intervals and every overlapping interval requires a seperate conference room.
	// Plus one is for all the other non-overlapping intervals
	return 1 + count
}

func main() {
	// Example usage
	intervals := []*Interval{
		{Start: 1, End: 2},
		{Start: 1, End: 2},
		{Start: 1, End: 2},
		{Start: 1, End: 2},
		{Start: 1, End: 2},
	}
	fmt.Println("Minimum number of conference rooms required:", MinMeetingRooms(intervals))

	// Example usage
	intervals = []*Interval{
		{Start: 0, End: 30},
		{Start: 5, End: 10},
		{Start: 15, End: 20},
	}
	fmt.Println("Minimum number of conference rooms required:", MinMeetingRooms(intervals))

	// Example usage
	intervals = []*Interval{
		{Start: 2, End: 7},
	}
	fmt.Println("Minimum number of conference rooms required:", MinMeetingRooms(intervals))
}
