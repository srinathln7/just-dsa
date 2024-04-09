package main

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {

	// Key Idea: If the size of the card is not divisible by `groupSize` then return false immediately.
	// Build a freq map to keep track of the number of occurences of each number in the `hand` slice.
	// Then sort the `hand` slice since the minimum element in the hand slice can have only one position as the starting number.

	n := len(hand)
	if n%groupSize != 0 {
		return false
	}

	freq := make(map[int]int)
	for _, num := range hand {
		freq[num]++
	}

	// Sort the cards so as to start from the minimum
	sort.Ints(hand)
	for _, card := range hand {
		// For a non-empty freq of the card
		if freq[card] > 0 {
			for i := 0; i < groupSize; i++ {

				// Check for consecutive cards within the group
				if freq[card+i] == 0 {
					return false
				}

				freq[card+i]--
			}
		}
	}

	return true
}
