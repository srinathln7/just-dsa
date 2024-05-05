package main

import "fmt"

type TreeNode struct {
	Val      int
	Children []*TreeNode
}

func findRoot(edges [][]int) *TreeNode {
	adj := make(map[int][]int)
	childset := make(map[int]struct{})

	for _, edge := range edges {
		p, c := edge[0], edge[1]
		adj[p] = append(adj[p], c)
		childset[c] = struct{}{}
	}

	var rootVal int
	for _, edge := range edges {
		p := edge[0]
		if _, exists := childset[p]; !exists {
			rootVal = p
			break
		}
	}

	fmt.Printf("Root val is %d \n", rootVal)
	visited := make(map[int]bool)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		visited[node.Val] = true

		for _, child := range adj[node.Val] {
			// IMPORTANT
			if !visited[child] {

				// Form the child node
				childNode := &TreeNode{Val: child}

				// Recurse
				dfs(childNode)

				// And then append
				node.Children = append(node.Children, childNode)
			}
		}
	}

	root := &TreeNode{Val: rootVal}
	dfs(root)

	return root
}

func printTree(root *TreeNode, level int) {
	if root == nil {
		return
	}
	fmt.Printf("Node %d at level %d\n", root.Val, level)
	for _, child := range root.Children {
		printTree(child, level+1)
	}
}

func main() {
	edges := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{1, 5},
		{1, 6},
		{1, 7},
		{2, 4},
		{2, 5},
		{2, 7},
		{3, 6},
		{4, 7},
	}
	tree := findRoot(edges)

	printTree(tree, 0)

	// Output
	//	Root val is 1
	// Node 1 at level 0
	// Node 2 at level 1
	// Node 4 at level 2
	// Node 7 at level 3
	// Node 5 at level 2
	// Node 3 at level 1
	// Node 6 at level 2
}
