package main

import (
	"fmt"
	"sort"
)

func Diameter(root *Node) int {

	// Basecase
	if root == nil {
		return 0
	}

	if len(root.Children) == 0 {
		return 0
	}

	var diamater int

	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}

		for _, child := range node.Children {
			dfs(child)
			maxDepth1, maxDepth2 := maxTwoDepths(child)
			diamater = max(diamater, maxDepth1+maxDepth2)
		}
	}

	dfs(root)
	return diamater
}

func maxTwoDepths(node *Node) (int, int) {

	if node == nil {
		return 0, 0
	}

	var depths []int
	for _, child := range node.Children {
		depths = append(depths, maxDepth(child))
	}

	fmt.Println("Depths", depths)
	// For uni-ary tree
	if len(depths) == 1 {
		return depths[0], 0
	}

	sort.Slice(depths, func(i, j int) bool {
		return depths[i] >= depths[j]
	})

	return depths[0], depths[1]
}

func maxDepth(root *Node) int {
	// Key Idea: Use recursion
	if root == nil {
		return 0
	}

	// At the root level `depth`
	depth := 1
	for _, child := range root.Children {
		// Plus 1 is for the child node that will be treated as the `root` in the next recursive stack frame
		depth = max(depth, 1+maxDepth(child))
	}
	return depth
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
