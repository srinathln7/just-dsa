package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {

	// Base case
	if head == nil || head.Next == nil {
		return head
	}

	sortedList := sortList(head.Next)
	return insert(sortedList, head.Val)
}

func insert(head *ListNode, val int) *ListNode {

	// Use the sentinel node technique to insert the elements properly

	result := &ListNode{Next: head}
	node := &ListNode{Val: val}

	curr := result
	for curr.Next != nil {
		if val <= curr.Next.Val {
			node.Next = curr.Next
			curr.Next = node
			break
		}
		curr = curr.Next
	}

	// Insertion at the end
	if curr.Next == nil {
		curr.Next = node
	}

	return result.Next
}
