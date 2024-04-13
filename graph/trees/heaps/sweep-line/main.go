package main

import (
	"container/heap"
)

type Event int

const (
	open Event = iota
	visit
	close
)

type Plane struct {
	x     int
	event Event
}

type MinHeap []Plane

func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Plane)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	o := old[n-1]
	*h = old[:n-1]
	return o
}
func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i].x == h[j].x {
		return h[i].event < h[j].event
	}

	return h[i].x < h[j].x
}

func (h MinHeap) Top() Plane     { return h[0] }
func (h *MinHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func fullBloomFlowers(flowers [][]int, people []int) []int {

	n := len(people)
	result := make([]int, n)

	minHeap := &MinHeap{}
	for _, flower := range flowers {
		heap.Push(minHeap, Plane{x: flower[0], event: open})
		heap.Push(minHeap, Plane{x: flower[1], event: close})
	}

	for i := 0; i < n; i++ {
		heap.Push(minHeap, Plane{x: people[i], event: visit})
	}

	countSet := make(map[int]int)
	count := 0
	for len(*minHeap) > 0 {

		if len(countSet) == n {
			break
		}

		plane := heap.Pop(minHeap).(Plane)
		switch plane.event {
		case 0:
			count++
		case 2:
			count--
		case 1:
			countSet[plane.x] = count
		}
	}

	for i := 0; i < n; i++ {
		result[i] = countSet[people[i]]
	}

	return result
}
