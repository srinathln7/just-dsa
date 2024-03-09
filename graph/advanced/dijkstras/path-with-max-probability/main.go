package main

import "container/heap"

// Node: Each node has an ID and a cost associated with it
type Node struct {
	id          int
	probability float64
}

type MaxHeap []Node

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].probability > h[j].probability }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Top() Node          { return h[0] }
func (h MaxHeap) Empty() bool        { return len(h) == 0 }

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
	adj := make(map[int][]Node, n+1)
	for i := 0; i < len(edges); i++ {
		src, dest, prob := edges[i][0], edges[i][1], succProb[i]
		adj[src] = append(adj[src], Node{id: dest, probability: prob})
		adj[dest] = append(adj[dest], Node{id: src, probability: prob}) // Undirected graph
	}

	// If no neighbour exists from `start_node`, then we cannot travel anywhere
	if len(adj[start_node]) == 0 {
		return 0
	}

	// Initialize a max-heap with the `start_node` same as end node with max probability 1.0
	maxHeap := &MaxHeap{}
	heap.Push(maxHeap, Node{id: start_node, probability: 1.0})

	maxProb := make(map[int]Node)
	for !maxHeap.Empty() {
		visiting := maxHeap.Top()
		heap.Pop(maxHeap)

		if _, visited := maxProb[visiting.id]; visited {
			continue
		}

		// IMPORTANT: Only mark the node as visited and store the max probability associated
		// with that path AFTER it has been popped out out of the max. heap.
		// For eg: To reach node 0 from node 0 => probability is 1.0
		maxProb[visiting.id] = Node{id: visiting.id, probability: visiting.probability}

		// Visit the neighbours of the recently visited node
		for _, node := range adj[visiting.id] {
			// Make sure the neighbour is not visited before
			if _, visited := maxProb[node.id]; !visited {
				// Mulitply the probability
				node.probability *= visiting.probability
				heap.Push(maxHeap, node)
			}
		}
	}

	// Make sure the path exists from `start_node` to `end_node`
	if node, exists := maxProb[end_node]; exists {
		return node.probability
	}

	return 0
}
