package main

import "container/heap"

// Node: Each node has an ID and a cost associated with it
type Node struct {
	id   int
	cost int
}

type MinHeap []Node

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Top() Node          { return h[0] }
func (h MinHeap) Empty() bool        { return len(h) == 0 }

func (h *MinHeap) Push(x any) { *h = append(*h, x.(Node)) }
func (h *MinHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	(*h) = (*h)[:n-1]
	return x
}

// shortestPath: Returns a map containing the optimum cost to reach all the
// `n` vertices in the graph from a source vertex `src`
func shortestPath(edges [][3]int, n, src int) map[int]int {

	shortest := make(map[int]int)

	// First form the adjacency list
	adjacencyList := make(map[int][]Node, n)
	for _, edge := range edges {
		start, dest, cost := edge[0], edge[1], edge[2]
		adjacencyList[start] = append(adjacencyList[start], Node{id: dest, cost: cost})
	}

	// Initialize the min heap and push the origin `src` node
	minHeap := &MinHeap{}
	heap.Push(minHeap, Node{id: src, cost: 0})

	for !minHeap.Empty() {
		visiting := minHeap.Top()
		heap.Pop(minHeap)

		// If the visiting node is already visited - skip the current iteration
		// This is because the node has already been visited and the minimum cost
		// associated with visiting that node has already been captured while poping that node from the minheap.
		if _, visited := shortest[visiting.id]; visited {
			continue
		}

		// IMPORTANT: Mark the visiting node as `visited` and store the "optimum" cost associated with visiting that node
		// The cost will always be optimum (minimum) since we are poping the element from min. heap
		shortest[visiting.id] = visiting.cost

		// Visit the neighbouring nodes of the visited node
		for _, node := range adjacencyList[visiting.id] {

			// Check if the node has not been visited yet
			if _, visited := shortest[node.id]; !visited {

				// MOST IMPORTANT: NEVER mark the node as visited here since we might have to push the same node (with different costs) multiple times into the heap.
				// We only mark it as `visited` while poping the element from the min. heap since we always want to record the min. cost
				// Accumulate the cost at every iteration. For example : If  A->B = 1 ; B->C =2 , the path A->B->C = 1+2 = 3
				node.cost += visiting.cost
				heap.Push(minHeap, node)
			}
		}

	}

	return shortest
}

func main() {
	edges := [][3]int{{1, 2, 1}, {1, 3, 4}, {2, 3, 2}, {3, 4, 5}, {2, 4, 7}}
	n := 4
	src := 1
	shortest := shortestPath(edges, n, src)
	for node, dist := range shortest {
		println("Shortest distance from node", src, "to node", node, ":", dist)
	}

	// Output:
	// Shortest distance from node 1 to node 1 : 0
	// Shortest distance from node 1 to node 2 : 1
	// Shortest distance from node 1 to node 3 : 3
	// Shortest distance from node 1 to node 4 : 8
}
