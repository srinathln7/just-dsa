package main

import "sort"

type UnionFind struct {
	parent []int
	rank   []int
}

func newUnionFind(n int) *UnionFind {

	parent := make([]int, n)
	rank := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}

	return &UnionFind{parent: parent, rank: rank}
}

func (uf *UnionFind) find(n int) int {
	p := uf.parent[n]
	for p != uf.parent[p] {
		uf.parent[p] = uf.parent[uf.parent[p]]
		p = uf.parent[p]
	}

	return p
}

func (uf *UnionFind) union(n1, n2 int) bool {
	p1, p2 := uf.find(n1), uf.find(n2)
	if p1 == p2 {
		return false
	}

	switch {
	case uf.rank[p2] < uf.rank[p1]:
		uf.parent[p2] = p1
	case uf.rank[p1] < uf.rank[p2]:
		uf.parent[p1] = p2
	default:
		uf.parent[p2] = p1
		uf.rank[p1]++
	}

	return true
}

func abs(val int) int {
	if val > 0 {
		return val
	}

	return -val
}

func minCostConnectPoints(points [][]int) int {

	// Key Idea: We deploy Kruskal's algorithm which internally uses the union-find data structure to find the minimum spanning tree
	// Sort the edges in ascending order and accumalte the cost until we are able to perform a union on the pair of the nodes.

	n := len(points)

	// Edges: Contains the index of the two points and their respective manhattan distance as their third point
	// For ex: Suppose we have edge =[0,1, 5], It means we have  points[0], points[1], and manhattan dist. b/w them is 5.
	var edges [][3]int
	for i, point := range points {
		x1, y1 := point[0], point[1]
		for j := i + 1; j < n; j++ {
			x2, y2 := points[j][0], points[j][1]
			edges = append(edges, [3]int{i, j, abs(x1-x2) + abs(y1-y2)})
		}
	}

	// Sort the edges in ascending order based on their cost
	// By this way, when we perform a union on the respective nodes,
	// we first pick the edges with the least weights.
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	var mincost int

	// Create a union find datastructure
	uf := newUnionFind(n)
	for _, edge := range edges {
		n1, n2, cost := edge[0], edge[1], edge[2]
		// Union will always operate on dis-join sets. Since the edges slice is sorted in ascending order
		// we always perform the union b/e two nodes with least weights first and connect them
		if uf.union(n1, n2) {
			mincost += cost
		}
	}

	return mincost
}
