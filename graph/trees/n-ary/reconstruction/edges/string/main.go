package main

import "fmt"

type TreeNode struct {
	Val      string
	Children []*TreeNode
}

func findRoot(edges [][]string) *TreeNode {
	adj := make(map[string][]string)
	childset := make(map[string]struct{})

	for _, edge := range edges {
		p, c := edge[0], edge[1]
		adj[p] = append(adj[p], c)
		childset[c] = struct{}{}
	}

	var rootVal string
	for _, edge := range edges {
		p := edge[0]
		if _, exists := childset[p]; !exists {
			rootVal = p
			break
		}
	}

	fmt.Printf("Root val is %s \n", rootVal)
	visited := make(map[string]bool)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		visited[node.Val] = true

		for _, child := range adj[node.Val] {
			if !visited[child] {
				childNode := &TreeNode{Val: child}
				dfs(childNode)
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
	fmt.Printf("Node %s at level %d\n", root.Val, level)
	for _, child := range root.Children {
		printTree(child, level+1)
	}
}

func main() {
	edges := [][]string{
		{"A", "B"},
		{"A", "C"},
		{"A", "D"},
		{"A", "E"},
		{"B", "F"},
		{"B", "G"},
		{"C", "H"},
		{"C", "I"},
		{"C", "J"},
		{"D", "K"},
		{"E", "L"},
		{"E", "M"},
		{"L", "N"},
	}

	tree := findRoot(edges)

	printTree(tree, 0)

	// Output
	// 	Root val is A
	// Node A at level 0
	// Node B at level 1
	// Node F at level 2
	// Node G at level 2.
	// Node C at level 1
	// Node H at level 2
	// Node I at level 2
	// Node J at level 2
	// Node D at level 1
	// Node K at level 2
	// Node E at level 1
	// Node L at level 2
	// Node N at level 3
	// Node M at level 2

}
