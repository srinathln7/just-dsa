package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func diameter(root *Node) int {

	// PREREQ: Diameter of binary tree
	// Key Idea: Similar to binary trees to find the diameter we need to find the depth of the tree. The catch here is
	// since the tree may have more tahn two children unlike binary nodes, we will have to find the maximum two depths
	// among all the nodes in the tree. Then the path length passing through that node will be equal to `maxDepth1+maxDepth2`

	// REMEMBER: The length of the longest path may or may not pass through the root. Hence it is important to run DFS among all nodes

	var diameter int

	var dfs func(node *Node) int
	dfs = func(node *Node) int {

		if node == nil {
			return 0
		}

		if len(node.Children) == 0 {
			return 1
		}

		maxDepth1, maxDepth2 := 0, 0
		for _, child := range node.Children {
			parentDepth := dfs(child)
			if parentDepth > maxDepth1 {
				maxDepth1, maxDepth2 = parentDepth, maxDepth1
			} else if parentDepth > maxDepth2 {
				maxDepth2 = parentDepth
			}
		}

		diameter = max(diameter, maxDepth1+maxDepth2)
		return 1 + maxDepth1
	}

	dfs(root)
	return diameter
}

func main() {
	// Define trees
	root1 := &Node{
		Val: 1,
		Children: []*Node{
			nil,
			{Val: 3, Children: []*Node{
				{Val: 5, Children: nil},
				{Val: 6, Children: nil},
			}},
			{Val: 2, Children: nil},
			{Val: 4, Children: nil},
		},
	}

	root2 := &Node{
		Val: 1,
		Children: []*Node{
			nil,
			{Val: 2, Children: nil},
			nil,
			{Val: 3, Children: []*Node{
				{Val: 5, Children: nil},
				nil,
				{Val: 6, Children: nil},
			}},
			{Val: 4, Children: nil},
		},
	}

	root3 := &Node{
		Val: 1,
		Children: []*Node{
			nil,
			{Val: 2, Children: []*Node{
				{Val: 3, Children: []*Node{
					{Val: 6, Children: []*Node{
						nil,
						nil,
						{Val: 11, Children: []*Node{
							nil,
							{Val: 14, Children: nil},
						}},
					}},
					nil,
				}},
				{Val: 4, Children: nil},
				{Val: 5, Children: []*Node{
					{Val: 7, Children: nil},
					{Val: 8, Children: []*Node{
						nil,
						{Val: 12, Children: []*Node{
							nil,
							nil,
							{Val: 13, Children: nil},
						}},
					}},
					{Val: 9, Children: nil},
					{Val: 10, Children: nil},
				}},
			}},
		},
	}

	fmt.Println(diameter(root1))
	fmt.Println(diameter(root2))
	fmt.Println(diameter(root3))
}
