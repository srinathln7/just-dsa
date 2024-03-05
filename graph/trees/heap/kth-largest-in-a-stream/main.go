package main

import "container/heap"

// Implement the heap interface
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

type KthLargest struct {
	minHeap *MinHeap
	k       int
}

func Constructor(k int, nums []int) KthLargest {
	kth := &KthLargest{
		minHeap: &MinHeap{},
		k:       k,
	}

	for _, num := range nums {
		kth.Add(num)
	}

	return *kth
}

func (this *KthLargest) Add(val int) int {
	// First push the value to the heap
	heap.Push(this.minHeap, val)

	// Key Idea: We can obtain the Kth largest element in a stream by always maintaining a min heap of size `k`
	// In a minHeap of size `k`, the kth largest element will be the first smallest element which will always
	// be located at the root of the heap
	for len(*(this.minHeap)) > this.k {
		heap.Pop(this.minHeap)
	}

	//minheap := *this.minHeap
	return (*this.minHeap)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
