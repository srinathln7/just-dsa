package main

import "container/heap"

type Node struct {
	id   int
	prob float64
}

type MaxHeap []Node

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].prob > h[j].prob }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Top() Node          { return h[0] }
func (h MaxHeap) IsEmpty() bool      { return len(h) == 0 }

func (h *MaxHeap) Push(x any) { *h = append(*h, x.(Node)) }
func (h *MaxHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	(*h) = (*h)[:n-1]
	return x
}

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {

	//  Key Idea: This problem is a variation of the classic Dijkstra's problem where instead of using min heap
	// we use a max. heap and we multiply the probability along the path rather than adding the cumulative cost.
	// IMPORTANT: The graph is undirected which means we need to form the adjaceny list both ways from `src` to
	// `dest` and vice-versa.

	// First, form the adjacency list of the given UNDIRECTED graph
	adj := make(map[int][]Node)
	for i, edge := range edges {
		src, dst, prob := edge[0], edge[1], succProb[i]

		// Bi-directional graph
		adj[src] = append(adj[src], Node{id: dst, prob: prob})
		adj[dst] = append(adj[dst], Node{id: src, prob: prob})
	}

	if len(adj[start_node]) == 0 {
		return 0
	}

	maxHeap := &MaxHeap{}
	maxProb := make(map[int]float64)
	heap.Push(maxHeap, Node{id: start_node, prob: 1.0})

	for !maxHeap.IsEmpty() {
		visiting := heap.Pop(maxHeap).(Node)

		// If already visited
		if _, exists := maxProb[visiting.id]; exists {
			continue
		}

		// IMPORTANT: Only mark the node as visited and store the max probability associated
		// with that path AFTER it has been popped out out of the max. heap.
		// For eg: To reach node 0 from node 0 => probability is 1.0
		// Mark the node as visited
		maxProb[visiting.id] = visiting.prob

		// Visit its neighbours
		for _, neighbour := range adj[visiting.id] {
			// IMPORTANT: Check if the node is already visited. If yes, the max. probabality associated with
			// visiting that node has already been captured
			if _, visited := maxProb[neighbour.id]; !visited {
				neighbour.prob *= visiting.prob
				heap.Push(maxHeap, neighbour)
			}
		}
	}

	// Make sure the path exists from `start_node` to `end_node`
	if val, exists := maxProb[end_node]; exists {
		return val
	}

	return 0
}
