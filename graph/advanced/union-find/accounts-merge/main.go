package main

import "sort"

type UnionFind struct {
	parent []int
	rank   []int // Optional parameter
}

func newUnionFind(account [][]string) *UnionFind {
	n := len(account)
	uf := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
	}

	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.rank[i] = len(account[i])
	}

	return uf
}

func (uf *UnionFind) find(n int) int {
	p := uf.parent[n]
	for p != uf.parent[p] {
		uf.parent[p] = uf.parent[uf.parent[p]]
		p = uf.parent[p]
	}

	return p
}

func (uf *UnionFind) union(n1, n2 int) {
	p1, p2 := uf.find(n1), uf.find(n2)
	if p1 != p2 {
		switch {
		case uf.rank[p1] > uf.rank[p2]:
			uf.parent[p2] = p1
		case uf.rank[p1] < uf.rank[p2]:
			uf.parent[p1] = p2
		default:
			uf.parent[p2] = p1
			uf.rank[p1] += uf.rank[p2]
		}
	}
}

func accountsMerge(accounts [][]string) [][]string {

	// Create a union find data structure
	uf := newUnionFind(accounts)
	emailIdx := make(map[string]int)

	// Range over the accounts
	for i, account := range accounts {
		emails := account[1:]
		for _, email := range emails {
			// Merge the two accounts into one account if atleast one common email exists
			if idx, exists := emailIdx[email]; exists {
				uf.union(idx, i)
			} else {
				emailIdx[email] = i
			}
		}
	}

	// Collect the merged accounts
	merged := make(map[int][]string)
	for email, idx := range emailIdx {
		// Find the email owner by finding the root (parent) of the corresponding email index
		root := uf.find(idx)
		merged[root] = append(merged[root], email)
	}

	// Format the merged accounts
	var res [][]string
	for root, emails := range merged {
		sort.Strings(emails)
		name := accounts[root][0]
		account := append([]string{name}, emails...)
		res = append(res, account)
	}

	return res
}
