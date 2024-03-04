package main

import (
	"fmt"
)

type Iterator struct {
	lists []int
	index int
}

func newIterator(lists []int) *Iterator {
	return &Iterator{
		lists: lists,
	}
}

func (i *Iterator) hasNext() (int, bool) {
	if i.index >= len(i.lists) {
		return -1, false
	}

	val := i.lists[i.index]
	i.index++
	return val, true
}

func interLeave(lists [][]int) [][]int {
	// MISTAKE: `results` length need not always be same as the length of the lists.
	// Consider the case where all the elements in the list is []
	//results := make([][]int, len(lists))

	var results [][]int
	iterators := make([]*Iterator, len(lists)) // Key idea is to make this a POINTER to the iterator struct

	// Key Idea: is to loop one row at a time and within each row we loop one element at a time
	// To do this we create a wrapper around iterator for each row in the lists
	for i, list := range lists {
		iterators[i] = newIterator(list)
	}

	for {
		var elements []int    // Slice containg interleaved elements for the current iteration
		var anyRemaining bool // Returns false only when all elements in the iterator are exhausted
		for _, iterator := range iterators {
			val, isPresent := iterator.hasNext()
			if isPresent {
				elements = append(elements, val)
				anyRemaining = true
			}
		}

		// When all elements are exhausted in the iterator
		if !anyRemaining {
			break
		}
		results = append(results, elements)
	}
	return results
}

func main() {
	lists := [][]int{{2, 3}, {1, 4, 5}, {}}
	interleaved := interLeave(lists)
	fmt.Println(lists)
	fmt.Println(interleaved) // Output: [[1 4 6] [2 5] [3]]
}
