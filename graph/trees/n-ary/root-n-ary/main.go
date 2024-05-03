package main

type Node struct {
	Val      string
	Children []*Node
}

func findRoot(tree []*Node) *Node {

	// Key Idea: To find the root of n-ary tree - if we do a level-order traversal of each of the nodes starting from the root
	// and apply XOR operations all along, we can turn this problem to a problem where we are given a slice of integers with each
	// integer repeating twice expecting only one integer. This `integer` x -> is the root of the tree.
	// XOR propeorties: x^0 = x, x^x=0 and it is associative and commutative

	charToInt := func(str string) int {
		return int(str[0] - 'A')
	}

	x := 0
	for _, node := range tree {
		x ^= charToInt(node.Val)
		for _, child := range node.Children {
			x ^= charToInt(child.Val)
		}
	}

	// `x` represents the root of the tree
	for _, node := range tree {
		if charToInt(node.Val) == x {
			return node
		}
	}

	return nil
}
