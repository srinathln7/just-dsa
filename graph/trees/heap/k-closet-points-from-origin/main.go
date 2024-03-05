package main

import (
	"container/heap"
	"math"
)

type Point struct {
	X    int
	Y    int
	dist float64
}

type MinHeap []Point

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].dist < h[j].dist }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Top() Point          { return h[0] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Point)) }
func (h *MinHeap) Pop() any {
	n := len(*h)
	top := (*h)[n-1]
	(*h) = (*h)[:n-1]
	return top
}

func getPointWithEuclideanDist(point *Point) *Point {
	point.dist = math.Sqrt(float64(point.X*point.X + point.Y*point.Y))
	return point
}

func kClosest(points [][]int, k int) [][]int {
	// Calculate the Eucledian distance for every point and store the resulting point in the min-heap
	// To return `k` closest points pop the heap `k` times
	kclosest := make([][]int, k)

	// Start populating the min-heap
	minHeap := &MinHeap{}
	for _, point := range points {
		euclideanPoint := getPointWithEuclideanDist(&Point{X: point[0], Y: point[1]})
		heap.Push(minHeap, *euclideanPoint)
	}

	// Pop the first k out
	for i := 0; i < k; i++ {
		euclideanPoint := heap.Pop(minHeap).(Point) // Do not forget to cast back
		kclosest[i] = []int{euclideanPoint.X, euclideanPoint.Y}
	}

	// Alternatively, Since order does not matter => We can use max heap as well
	// Advantage is we dont have to store `n` points but only `k` points at a time
	// Only the length of the heap is greater than `k` start popping and the farther points will be removed
	return kclosest
}
