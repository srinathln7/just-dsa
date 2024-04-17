package main

type bSTIterator struct {
	Stack []*TreeNode
}

func constructor(root *TreeNode) bSTIterator {

	// Key Idea: Is to use a Stack and push all values to the left of the root to the stack
	// This is because since it is a BST all values to the left < root < all values to the right
	stack := make([]*TreeNode, 0)
	bstIter := &bSTIterator{Stack: stack}
	bstIter.Push(root)
	return *bstIter
}

func (this *bSTIterator) Next() int {

	// IMPORTANT: To note that once we pop the left-most element we need to push its right child (if its exists) to the stack
	// This is because the right child of the left-subtree will still be smaller in value compared to the previous root value
	// Rem in BST: All left subtree val < root val < All right subtree val
	topNode := this.Pop()
	if topNode.Right != nil {
		this.Push(topNode.Right)
	}

	return topNode.Val
}

func (this *bSTIterator) hasNext() bool {
	return len(this.Stack) > 0
}

func (this *bSTIterator) Push(node *TreeNode) {
	for node != nil {
		this.Stack = append(this.Stack, node)
		node = node.Left
	}
}

func (this *bSTIterator) Pop() *TreeNode {
	topNode := this.Stack[len(this.Stack)-1]
	this.Stack = this.Stack[:len(this.Stack)-1]
	return topNode
}
