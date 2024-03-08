package main

import (
	"container/heap"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Top() int           { return h[0] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Top() int           { return h[0] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianFinder struct {
	maxHeap MaxHeap
	minHeap MinHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		maxHeap: MaxHeap{},
		minHeap: MinHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {

	switch {
	// Always keep the smaller values in the max heap
	case this.maxHeap.Len() == 0 || num < (this.maxHeap).Top():
		heap.Push(&this.maxHeap, num)
	// Viceversa, keep the larger values in the minheap
	default:
		heap.Push(&this.minHeap, num)
	}

	// Rebalance heaps after adding every num if necessary
	switch {
	case this.maxHeap.Len()-this.minHeap.Len() > 1:
		heap.Push(&this.minHeap, this.maxHeap.Top())
		heap.Pop(&this.maxHeap)
	case this.minHeap.Len()-this.maxHeap.Len() > 1:
		heap.Push(&this.maxHeap, this.minHeap.Top())
		heap.Pop(&this.minHeap)
	}

}

func (this *MedianFinder) FindMedian() float64 {
	switch {
	case this.maxHeap.Len() == this.minHeap.Len():
		return float64(this.maxHeap.Top()+this.minHeap.Top()) / 2
	case this.maxHeap.Len() > this.minHeap.Len():
		return float64(this.maxHeap.Top())
	default:
		return float64(this.minHeap.Top())
	}
}
