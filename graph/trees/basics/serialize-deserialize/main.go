package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {

	// Key Idea: To serialize and deserialize a tree - let us perform a preorder traversal of the tree
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {

	var res []string

	var dfs func(curr *TreeNode)
	dfs = func(curr *TreeNode) {

		// IMPORTANT: Base case - After encountering a `nil` node, append and then `return`
		if curr == nil {
			res = append(res, "N")
			return
		}

		res = append(res, strconv.Itoa(curr.Val))
		dfs(curr.Left)
		dfs(curr.Right)
	}

	dfs(root)
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {

	res := strings.Split(data, ",")
	idx := 0

	var dfs func() *TreeNode
	dfs = func() *TreeNode {

		// IMPORTANT: The reason why we keep this function as a empty argument is because the `idx` needs to be global

		// Prints 0,1,2,3,4,5,6
		// fmt.Printf("index: %d \n", idx)

		if res[idx] == "N" {
			idx += 1
			return nil
		}

		val, _ := strconv.Atoi(res[idx])
		node := &TreeNode{Val: val}
		idx += 1
		node.Left = dfs()
		node.Right = dfs()

		return node
	}

	return dfs()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
