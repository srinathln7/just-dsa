package main

import (
	"container/heap"
)

type Element struct {
	val  int
	freq int
}

type MinHeap []Element

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].freq < h[j].freq }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Top() Element       { return h[0] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Element))
}

func (h *MinHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {

	// Key Idea: Build the freq table and then maintain a min heap of constant size `k` to store top `k` freq elements/
	// Building the `freq` table will let us know how many unique elements are there in the map. Then we range over the
	// this map and keep pushing to the min heap until size `k`. Once it reaches size `k`, we keep pushing and popping
	// simulataneously. Time complexity O(n. log k)

	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	minHeap := &MinHeap{}

	// Every key is unique as we loop over the `freq` map
	for key, val := range freq {
		// Push to min heap and maintain a const heap size `k`
		heap.Push(minHeap, Element{val: key, freq: val})
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}

	result := make([]int, k)
	for i, element := range *minHeap {
		result[i] = element.val
	}

	return result
}
