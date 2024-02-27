package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	// preorder => root => left_subtree -> right_subtree
	// inorder =>  left_subtree -> root -> right_subtree

	// Key Idea: We can extract the root of the entire tree from the first index of the preorder traversal array
	// Locate the index in which that element lies in the inorder traversal arrray. Then all elements to the left
	// of it form a left subtree and to the right of it form the right subtree. We can then build the tree recursively

	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	var rootIdx, n int = 0, len(preorder)
	for i := 0; i < n; i++ {
		if inorder[i] == root.Val {
			rootIdx = i
			break
		}
	}

	// Build the tree recursively.
	// Ex: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7] => rootIdx = 1
	// All elements to the left of rootIdx in inorder traversal array will form the left subtree and to the right
	// will form the right subtree. No. of elements in the left subtree will equal to the value of rootIdx (since
	// array index starts at zero) and right elements will be `n-rootIdx-1` where n is the length of the traversal array.

	// Picking the index is tricker. Key is the following observation
	// FOR THE PREORDER TRAVERSAL ARRAY: We know we have to drop the first index in and there are
	// in total `rootIdx` items in the left subtree. So the index ranges from 1 (incl.) to 1+rootIdx (excl.).
	// The last index=1+rootIdx means that by the time we finish building the left subtree and do another preorder
	// traversal we will be in the root of the full tree. Next for the right subtree we start from 1+rootIdx (incl.)
	// FOR INORDER TRAVERSAL ARRAY: Rather than asking where do we start we ask where do we end after building the left subtree?
	// Just before the root of the full tree. And for the right subtree, you start from the next element to your root `1+rootIdx`
	root.Left = buildTree(preorder[1:1+rootIdx], inorder[:rootIdx])
	root.Right = buildTree(preorder[1+rootIdx:], inorder[1+rootIdx:])
	return root
}
