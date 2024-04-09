package main

/**
 * @param n: An integer
 * @param edges: a list of undirected edges
 * @return: true if it's a valid tree, or false
 */

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

func ValidTree(n int, edges [][]int) bool {
	uf := newUnionFind(n)
	for _, edge := range edges {
		n1, n2 := edge[0], edge[1]

		// Redundant connection or cycle detected
		if !uf.union(n1, n2) {
			return false
		}
	}

	return true
}
