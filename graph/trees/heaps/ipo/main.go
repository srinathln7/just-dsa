package main

import (
	"container/heap"
	"sort"
)

type Project struct {
	capital int
	profit  int
}

type MaxHeap []Project

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].profit > h[j].profit }
func (h MaxHeap) Empty() bool        { return len(h) == 0 }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Top() Project       { return h[0] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(Project)) }
func (h *MaxHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {

	// KEY IDEA: Create a project struct with profits and capital info.
	// Sort the project according to its capital in ascending order and continue pushing the project to the max heap
	// (maintained as per the projects profit.) until your current capital can afford that project. After that,
	// pop the project until k reaches 0 or no further project is left on the heap.

	n := len(profits)
	maxHeap := &MaxHeap{}

	// Push the project to max heap
	projects := make([]Project, n)
	for i := 0; i < n; i++ {
		projects[i] = Project{capital: capital[i], profit: profits[i]}
	}

	// Sort the project in ascending order as per the capital
	sort.Slice(projects, func(i, j int) bool {
		return projects[i].capital < projects[j].capital
	})

	i := 0
	for k > 0 {
		// Keep pushing to the max heap until we can afford the project with our current capital
		for i < len(projects) && w >= projects[i].capital {
			heap.Push(maxHeap, projects[i])
			i++
		}

		// If there is no projects left in the heap
		if len(*maxHeap) == 0 {
			break
		}

		currentProject := (*maxHeap).Top()
		w += currentProject.profit
		heap.Pop(maxHeap)

		k--
	}

	return w
}
