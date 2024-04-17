package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {

	// Key Idea: Use a stack

	var stack []*TreeNode
	curr := root
	for curr != nil {
		stack = append(stack, curr)
		curr = curr.Left
	}

	return BSTIterator{stack: stack}
}

func (this *BSTIterator) Next() int {
	n := len(this.stack)
	node := this.stack[n-1]
	this.stack = this.stack[:n-1]

	// Push the left descendants of the right child, if it exists
	if node.Right != nil {
		curr := node.Right
		for curr != nil {
			this.stack = append(this.stack, curr)
			curr = curr.Left
		}
	}

	return node.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
