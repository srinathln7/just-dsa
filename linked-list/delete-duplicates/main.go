package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	// Key Idea: Since the given list is in sorted order, check if the adjacent elements are equal and if yes
	// skip that node "else" traverse normally

	// Base case
	if head == nil || head.Next == nil {
		return head
	}

	curr := head
	for curr != nil {
		// Account for multiple duplicate elements through `if/else` block
		if curr.Next != nil && curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return head
}
