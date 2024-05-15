package main

import "fmt"

type TreeNode struct {
	Val      string
	Children []*TreeNode
}

func getEdges(sequences [][]string) [][2]string {

	edgeSet := make(map[[2]string]bool)
	var edges [][2]string

	for i := 0; i < len(sequences); i++ {

		for j := 0; j < len(sequences[i])-1; j++ {

			edge := [2]string{sequences[i][j], sequences[i][j+1]}
			if edgeSet[edge] {
				continue
			}

			edgeSet[edge] = true
			edges = append(edges, edge)
		}
	}

	return edges

	// var pairs [][2]string
	// seen := make(map[[2]string]bool)

	// for _, seq := range sequences {
	// 	for i := 0; i < len(seq)-1; i++ {
	// 		for j := i + 1; j < len(seq); j++ {
	// 			pair := [2]string{seq[i], seq[j]}

	// 			if !seen[pair] {
	// 				pairs = append(pairs, pair)
	// 				seen[pair] = true
	// 			}
	// 		}
	// 	}
	// }

	// return pairs
}

func findRoot(edges [][2]string) *TreeNode {
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
	sequences := [][]string{
		{"root", "food"},
		{"root", "food", "chips"},
		{"root", "food", "chips", "doritos"},
		{"root", "drink", "soda"},
		{"root", "drink", "soda", "diet coke"},
	}

	edges := getEdges(sequences)
	root := findRoot(edges)
	printTree(root, 0)

	// Output
	// Root val is root
	// Node root at level 0
	// Node food at level 1
	// Node chips at level 2
	// Node doritos at level 3
	// Node drink at level 1
	// Node soda at level 2
	// Node diet coke at level 3

	//	root
	//	/               \
	//
	// food                 drink
	// /                               |
	// chips                    soda
	// |                                |
	// doritos                   diet coke
}
