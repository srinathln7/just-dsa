package main

import (
	"container/heap"
	"sort"
)

type Element struct {
	val  int
	diff int
}

type MaxHeap []Element

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	if h[i].diff == h[j].diff {
		return h[i].val > h[j].val
	}
	return h[i].diff > h[j].diff
}

func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Top() Element  { return h[0] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(Element))
}

func (h *MaxHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	(*h) = (*h)[:n-1]
	return x
}

func findClosestElements(arr []int, k int, x int) []int {

	n := len(arr)
	if n == k {
		return arr
	}

	result := make([]int, k)
	maxHeap := &MaxHeap{}
	for _, num := range arr {
		element := Element{val: num, diff: abs(num - x)}
		heap.Push(maxHeap, element)

		// Maintain a constant heap of size `k`
		if (*maxHeap).Len() > k {
			heap.Pop(maxHeap)
		}
	}

	for i := k - 1; i >= 0; i-- {
		element := heap.Pop(maxHeap).(Element)
		result[i] = element.val
	}

	sort.Ints(result)
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
