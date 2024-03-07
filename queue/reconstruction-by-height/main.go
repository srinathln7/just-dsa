package main

import "sort"

func reconstructQueue(people [][]int) [][]int {

	// Key Idea: The goal is to reconstruct the queue according to people's height AND the number of people who are of the same or greater height
	// standing in front of them

	// First step : We sort the people based on their height in descending order.  However, if two people have the same height, then we sort them
	// according to the number of people standing infront of them in ascending order. Why descending order for people with different heights and ascending order for ppl
	// with same height?  This is because since a valid queue can always be reconstructed,  there will always be one tallest person who will have zero members taller than them standing
	// in front of them. Consider ex: [[7,0], [7,1]]. We need this so that  after we sort the people, the first element will have always zero members standing in front of them.

	sort.Slice(people, func(i, j int) bool {

		// Peoplw with same height - sort them in ascending order according to the number of ppl standing infront of them
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}

		// People with different height - sort them according to their height in descending order
		return people[i][0] > people[j][0]
	})

	// Operate this sorted people and fill up the queue
	var queue [][]int

	for _, p := range people {
		pplInFront := p[1]

		// Add the current person to the queue
		queue = append(queue, p)

		// MOST IMPORTANT: Prior to inserting the person at the appropriate index move the elements to the right of this index by one unit
		// Rem: This will pop out the last element and that is totally fine because the last element is only an array with single element which is not necessary
		// Refer to `pre-req.go` for any further clarifications
		copy(queue[pplInFront+1:], queue[pplInFront:])

		// People with Zero ppl in front will inserted at zeroth index, 1 ppl in front will be inserted at first index and so on
		queue[pplInFront] = p
	}

	return queue

}
