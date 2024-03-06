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

func (i *Iterator) hasNext() bool {
	if i.index >= len(i.lists) {
		return false
	}
	return true
}

func (i *Iterator) next() int {
	val := i.lists[i.index]
	i.index++
	return val
}

func interLeave(lists [][]int) [][]int {

	// Key Idea: is to loop one row at a time and within each row we loop one element at a time
	// To do this we create a wrapper around iterator for each row in the lists

	// MISTAKE: `results` length need not always be same as the length of the lists.
	// It will be equal to the length of the maximum number of elements among all the lists
	// Consider the case where all the elements in the list is []
	//results := make([][]int, len(lists))

	var results [][]int
	iterators := make([]*Iterator, len(lists)) // Key idea is to make this a POINTER to the iterator struct
	for i, list := range lists {
		iterators[i] = newIterator(list)
	}

	for {
		var elements []int    // Slice containg interleaved elements for the current iteration
		var anyRemaining bool // Returns false only when all elements in the iterator are exhausted
		for _, iterator := range iterators {
			if iterator.hasNext() {
				elements = append(elements, iterator.next())
				anyRemaining = true
			}
		}

		// When all elements in all the iterators are exhausted
		if !anyRemaining {
			break
		}
		results = append(results, elements)
	}

	return results
}

func main() {
	lists := [][]int{{1}, {2, 3}, {4, 5, 6}}
	interleaved := interLeave(lists)
	fmt.Println(lists)
	fmt.Println(interleaved) // Output: [1 2 4] [3 5] [6]]
}
