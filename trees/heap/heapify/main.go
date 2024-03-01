package main

import "fmt"

type Heap interface {
	Push(x int)
	Pop() int
}

type MinHeap struct {
	arr []int
}

func (h *MinHeap) Push(x int) {
	h.arr = append(h.arr, x)
	// Restore the heap property
	heapifyUp(h, len(h.arr)-1)
}

func (h *MinHeap) Pop() int {
	x := h.arr[0]
	h.arr = h.arr[1:]

	// Restore the heap property
	heapifyDown(h, 0)
	return x
}

// heapifyUp: Restores the heap property once a new element is added to the heap.
// We call it heapifyUp because the element is initially added to the last index
// and if it's greater than its parent then it needs to be moved upwards. In the worst
// sceanario, the newly added element is the minimum in the entire heap and in that case,
// it needs to be shoved all the way up to the root.
func heapifyUp(h *MinHeap, index int) {
	//  Keep iterating until the newly added element's parent value is lesser than
	// the element's value
	for index > 0 {
		parentIdx := (index - 1) / 2

		// If the parents value is lesser than the newly child's value - we have found the correct spot
		if h.arr[parentIdx] <= h.arr[index] {
			break
		}

		// Swap if the parent's (roots) value is greater than the elements value
		h.arr[parentIdx], h.arr[index] = h.arr[index], h.arr[parentIdx]
		index = parentIdx
		heapifyUp(h, index)
	}
}

// heapifyDown: Restores the heap propery once the element is removed from the heap.
// We call it heapifyDown because before the root node is popped, we swap the last element
// and this swapped element (now located in the root) might have tobe shoved all the way to
// the bottom of the tree if its value is greater than its child
func heapifyDown(h *MinHeap, index int) {
	n := len(h.arr)

	// // Keep track of the left and right child's index at every iteration
	leftIdx := 2*index + 1
	rightIdx := 2*index + 2
	smallestIdx := index

	if leftIdx < n && h.arr[leftIdx] < h.arr[smallestIdx] {
		smallestIdx = leftIdx
	}

	if rightIdx < n && h.arr[rightIdx] < h.arr[smallestIdx] {
		smallestIdx = rightIdx
	}

	// If the smallest index is left unchanged after checking with the left and right child
	// we found the right spot for the node
	if smallestIdx != index {
		h.arr[smallestIdx], h.arr[index] = h.arr[index], h.arr[smallestIdx]
		heapifyDown(h, smallestIdx)
	}

	// Alternatively: We can heapify iteratively as well

	// for rightIdx < n {

	// 	if leftIdx < n && h.arr[leftIdx] < h.arr[smallestIdx] {
	// 		smallestIdx = leftIdx
	// 	}

	// 	if rightIdx < n && h.arr[rightIdx] < h.arr[smallestIdx] {
	// 		smallestIdx = rightIdx
	// 	}

	// 	// If the smallest index is left unchanged after checking with the left and right child
	// 	// we found the right spot for the node
	// 	if smallestIdx == index {
	// 		break
	// 	}

	// 	h.arr[smallestIdx], h.arr[index] = h.arr[index], h.arr[smallestIdx]
	// 	index = smallestIdx
	// 	leftIdx = 2*index + 1
	// 	rightIdx = 2*index + 2
	// }
}

func main() {

	heap := &MinHeap{}

	heap.Push(5)
	heap.Push(4)
	heap.Push(3)
	heap.Push(2)
	heap.Push(1)

	fmt.Println(heap.arr)

	fmt.Println("Popped element", heap.Pop())
	fmt.Println("Popped element", heap.Pop())

	fmt.Println(heap.arr)

	heap2 := &MinHeap{}
	heap2.Push(5)
	heap2.Push(4)
	heap2.Push(3)

	fmt.Println(heap2.arr)
}
