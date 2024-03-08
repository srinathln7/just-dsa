package main

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

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

func (h MinHeap) Top() int {
	return h[0]
}

func (h MinHeap) Empty() bool {
	return len(h) == 0
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

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

func (h MaxHeap) Top() int {
	return h[0]
}

func (h MaxHeap) Empty() bool {
	return len(h) == 0
}

func getMedian(minHeap *MinHeap, maxHeap *MaxHeap, k int) float64 {
	if k%2 == 0 {
		return float64(minHeap.Top()+maxHeap.Top()) / 2
	}

	return float64(maxHeap.Top())
}

func medianSlidingWindow(nums []int, k int) []float64 {
	// Initialize two heaps
	minHeap := &MinHeap{}
	maxHeap := &MaxHeap{}

	// Map to keep track of elements moving out of the window
	outOfWindow := make(map[int]int)

	// Slice to store the result
	res := []float64{}

	// Iterate through the first window to populate the heaps
	for i := 0; i < k; i++ {
		heap.Push(maxHeap, nums[i])
	}

	// Balance the heaps
	for i := 0; i < k/2; i++ {
		heap.Push(minHeap, heap.Pop(maxHeap))
	}

	// Calculate median for the first window and append to result
	res = append(res, getMedian(minHeap, maxHeap, k))

	// Slide the window
	for i := k; i < len(nums); i++ {
		// Calculate the outgoing and incoming elements
		outgoing, incoming := nums[i-k], nums[i]

		// Update the outOfWindow map
		if _, ok := outOfWindow[incoming]; ok {
			outOfWindow[outgoing]++
		} else {
			outOfWindow[outgoing] = 1
		}

		// Now until this the heap is balance. Once we start adding elements and removing elements to/from the heap by sliding the window,
		// heap will become unbalanced. So we need to start balancing the heap.
		// REMEMBER: A balanced heap will have atmost difference of 1 b/w minHeap and maxHeap
		balance := 0

		// Check if the outgoing element is lesser than the max heap top. Then we decrement the balance.
		// We choose this convention because remember we started filling the max heap first. We can choose any valid convention.
		// If we have odd number of elements, the maxheap will always have more elements than the min heap.
		switch {
		case outgoing <= (*maxHeap).Top():
			balance--
		default:
			balance++
		}

		// Push the incoming element to its corresponding heap and update the balance accordingly
		switch {
		case !(*maxHeap).Empty() && incoming <= (*maxHeap).Top():
			heap.Push(maxHeap, incoming)
			balance++
		default:
			heap.Push(minHeap, incoming)
			balance--
		}

		// Balance the heap by restoring `balance` to zero, if not already
		switch {
		// max. heap has lesser elements than min. heap
		case balance < 0:
			heap.Push(maxHeap, heap.Pop(minHeap))

		// min. heap has more elements than max. heap
		case balance > 0:
			heap.Push(minHeap, heap.Pop(maxHeap))
		}

		// Remove out-of-window elements from the heaps
		for !(*maxHeap).Empty() && outOfWindow[maxHeap.Top()] > 0 {
			outOfWindow[maxHeap.Top()]--
			heap.Pop(maxHeap)
		}

		for !(*minHeap).Empty() && outOfWindow[minHeap.Top()] > 0 {
			outOfWindow[minHeap.Top()]--
			heap.Pop(minHeap)
		}

		// Calculate median and append to result
		res = append(res, getMedian(minHeap, maxHeap, k))
	}

	return res
}
