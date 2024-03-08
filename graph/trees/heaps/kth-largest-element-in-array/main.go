package main

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func findKthLargest(nums []int, k int) int {

	// Key Idea: If we maintain a min heap of size `k`, kth largest element will be located at the root of the heap
	// Remember: Larget => min Heap , Smallest => Max Heap

	minHeap := &MinHeap{}
	for _, num := range nums {
		heap.Push(minHeap, num)

		// Maintain a constant heap size k
		if (*minHeap).Len() > k {
			heap.Pop(minHeap)
		}
	}

	return (*minHeap)[0]
}
