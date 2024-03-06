package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	Stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {

	// Key Idea: Is to use a Stack and push all values to the left of the root to the stack
	// This is because since it is a BST all values to the left < root < all values to the right
	stack := make([]*TreeNode, 0)
	bstIter := &BSTIterator{Stack: stack}
	bstIter.Push(root)
	return *bstIter
}

func (this *BSTIterator) Next() int {

	// IMPORTANT: To note that once we pop the left-most element we need to push its right child (if its exists) to the stack
	// This is because the right child of the left-subtree will still be smaller in value compared to the previous root value
	// Rem in BST: All left subtree val < root val < All right subtree val
	topNode := this.Pop()
	if topNode.Right != nil {
		this.Push(topNode.Right)
	}

	return topNode.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.Stack) > 0
}

func (this *BSTIterator) Push(node *TreeNode) {
	for node != nil {
		this.Stack = append(this.Stack, node)
		node = node.Left
	}
}

func (this *BSTIterator) Pop() *TreeNode {
	topNode := this.Stack[len(this.Stack)-1]
	this.Stack = this.Stack[:len(this.Stack)-1]
	return topNode
}
