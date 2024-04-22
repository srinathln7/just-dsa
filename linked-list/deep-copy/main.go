package main

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {

	// Key Idea: Since the linkedlist in this case also contains a `Random` pointer which can potentially point to any node on the linkedlist
	// we can treat the list as a graph with possible cycles. To avoid recursing endlessly, we introduce `memoization` technique

	if head == nil {
		return nil
	}

	// store mapping from curr node to copy node
	visited := make(map[*Node]*Node)
	var dfs func(curr *Node) *Node
	dfs = func(curr *Node) *Node {

		if val, exists := visited[curr]; exists {
			return val
		}

		copyNode := &Node{Val: curr.Val}

		visited[curr] = copyNode

		if curr.Next != nil {
			copyNode.Next = dfs(curr.Next)
		}

		if curr.Random != nil {
			copyNode.Random = dfs(curr.Random)
		}

		return copyNode
	}

	return dfs(head)
}
