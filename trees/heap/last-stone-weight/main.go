package main

import "container/heap"

func lastStoneWeight(stones []int) int {

	// Key Idea: We need a data structure where we get the largest (heaviest) element at the top
	// Every time we pop it, we need the second heaviest element to be at the top => Max Heap

	maxheap := &MaxHeap{}

	// Push all the stones to the max heap O(n logn)
	for _, stone := range stones {
		heap.Push(maxheap, stone)
	}

	// Repeat the process until there is only one or no stone left
	for maxheap.Len() > 1 {
		// Get the 2 heaviest stones
		stone1 := heap.Pop(maxheap).(int)
		stone2 := heap.Pop(maxheap).(int)

		newStoneWeight := stone1 - stone2
		if newStoneWeight != 0 {
			heap.Push(maxheap, newStoneWeight)
		}
	}

	// If all stones are destroyed then return 0
	if maxheap.Len() == 0 {
		return 0
	}

	return (*maxheap)[0]
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}
