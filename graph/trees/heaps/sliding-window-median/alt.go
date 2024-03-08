package main

import (
	"container/heap"
)

type MedianFinder struct {
	maxHeap MaxHeap
	minHeap MinHeap
}

func newMedianFinder() MedianFinder {
	return MedianFinder{
		maxHeap: MaxHeap{},
		minHeap: MinHeap{},
	}
}

func (this *MedianFinder) AddNum(k int, num ...int) {

	for i := 0; i < k; i++ {
		switch {
		// Always keep the smaller values in the max heap
		case this.maxHeap.Len() == 0 || num[i] < (this.maxHeap).Top():
			heap.Push(&this.maxHeap, num[i])
		// Viceversa, keep the larger values in the minheap
		default:
			heap.Push(&this.minHeap, num[i])
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

func MedianSlidingWindow(nums []int, k int) []float64 {

	var result []float64
	n := len(nums)

	// First window
	mf := newMedianFinder()
	mf.AddNum(k, nums[0:k]...)
	result = append(result, mf.FindMedian())

	// Slide the window
	for i := k; i < n; i++ {
		window := nums[i-k+1 : i]
		mf := newMedianFinder()
		mf.AddNum(k, window[0:k]...)
		result = append(result, mf.FindMedian())
	}

	return result
}
