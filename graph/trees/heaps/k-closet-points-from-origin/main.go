package main

import (
	"container/heap"
	"math"
)

type Point struct {
	idx  int // Store the point's idx rather than the actual point. This is much efficient
	dist float64
}

type MaxHeap []Point

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i].dist > h[j].dist }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Top() Point          { return h[0] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(Point)) }
func (h *MaxHeap) Pop() any {
	n := len(*h)
	top := (*h)[n-1]
	(*h) = (*h)[:n-1]
	return top
}

func getEuclideanDist(x, y int) float64 {
	return math.Sqrt(float64(x*x + y*y))
}

func kClosest(points [][]int, k int) [][]int {

	// Key Idea: Maintain a max heap of size `k`

	result := make([][]int, k)
	maxHeap := &MaxHeap{}
	for idx, point := range points {
		dist := getEuclideanDist(point[0], point[1])
		heap.Push(maxHeap, Point{idx: idx, dist: dist})

		// When heap size exceeds size `k` - start popping out so that the heap maintains the `k` closest points at any given index
		if idx >= k {
			heap.Pop(maxHeap)
		}
	}

	// Fill in the result
	for i := 0; i < k; i++ {
		point := heap.Pop(maxHeap).(Point)
		result[i] = points[point.idx]
	}

	return result
}
