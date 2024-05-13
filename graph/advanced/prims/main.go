package main

import "container/heap"

type Point struct {
	idx  int // Represents the DESTINATION index of the point. Rather than storing the actual (x,y) point co-ordinates, it is much cheaper and efficient to store the point index.
	cost int // Represents the manhattan distance in this case.
}

type MinHeap []Point

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Top() Point         { return h[0] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func abs(val int) int {
	if val > 0 {
		return val
	}

	return -val
}

func minCostConnectPoints(points [][]int) int {

	// Key Idea: To find the min. cost to connect all the points in the graph is equivalent to finding the minimum spanning tree (MST)
	// We can deploy Prim's algorithm which is very similar to Dijkstra's algorithm. The main difference between them is Dijksjtra's is
	// exclusively used for finding the shortest path between the specified `src` and `dst` node where as in Prims, we must ensure
	// we visit all nodes in the graph only once. Since, visiting all nodes in the graph without any cycles means that the graph is connected
	// the resulting grah will be a tree. Since we are minimizing the cost associated with visiting every node in the graph, the resulting
	// tree is the MINIMUM SPANNING TREE (MST).

	n := len(points)

	// Form the adjacency list of the graph
	// adj[i] represents the adjacency list of the `i`th point in the given `points` slice of type [][]int
	adj := make([][]Point, n)
	for i, point := range points {
		x1, y1 := point[0], point[1]
		for j := i + 1; j < n; j++ {
			x2, y2 := points[j][0], points[j][1]
			manHatDist := abs(x2-x1) + abs(y2-y1)
			// Since it is undirected graph, we need to account for both directions
			adj[i] = append(adj[i], Point{idx: j, cost: manHatDist})
			adj[j] = append(adj[j], Point{idx: i, cost: manHatDist})
		}
	}

	// Init a minHeap and push the starting point. We can choose arbitarily any point as starting point from the points slice
	// Here we choose the first point i.e. zeroth index and the cost associated with visiting the starting point from starting
	// point is always zero (manhattan dist. is zero).
	minHeap := &MinHeap{}
	heap.Push(minHeap, Point{idx: 0, cost: 0})

	// Declare a visited map to keep track of the nodes visited
	// VERY IMPORTANT: Key is the point `idx` of type `int` and not the entire `node`. It is not only because
	// it is efficient in doing that but also struct equality implies checking equality for not just the
	// `idx` field but also the `cost` field. For ex: Point{1,4} is different from Point{1,3}.  If we had
	//  used `node` as key rather than `node.idx` then the heap would still calculate the cost for visiting
	// point `1` although point `1` is already visited. Heap would visit {1,4} and then {1,3} again which is
	// not what we wanted and will output the wrong result.
	visited := make(map[int]bool)
	var minCost int
	for len(visited) != n {
		visiting := minHeap.Top()
		heap.Pop(minHeap)

		if visited[visiting.idx] {
			continue
		}

		// Mark the `visiting` point as `visited`
		minCost += visiting.cost
		visited[visiting.idx] = true

		// Visit the neighbouring nodes of the recently visited node
		for _, point := range adj[visiting.idx] {
			if !visited[point.idx] {
				heap.Push(minHeap, point)
			}
		}
	}

	return minCost
}
