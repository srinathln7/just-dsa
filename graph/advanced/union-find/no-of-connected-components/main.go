package main

type UnionFind struct {
	parent []int
	rank   []int // Default value is zero
}

func newUnionFind(n int) UnionFind {

	uf := UnionFind{}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}

	return uf
}

func (uf *UnionFind) find(n int) int {

	p := uf.parent[n]

	// Path compression algorithm
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

	// Perform union operation

	switch {
	case uf.rank[p2] < uf.rank[p1]:
		uf.parent[p2] = p1
	case uf.rank[p1] > uf.rank[p2]:
		uf.parent[p1] = p2
	default:
		uf.parent[p2] = p1
		uf.rank[p1]++
	}

	return true
}

func countComponents(n int, edges [][]int) int {

	uf, count := newUnionFind(n), 0
	for _, edge := range edges {
		n1, n2 := edge[0], edge[1]
		if uf.union(n1, n2) {
			count++
		}
	}

	return n - count
}
