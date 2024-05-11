package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {

	// Base case
	if head == nil || head.Next == nil {
		return head
	}

	sentinel := &ListNode{Next: head}

	// This is because in insertion sort we always consider the first element to be sorted
	prev, curr := head, head.Next
	for curr != nil {
		if curr.Val >= prev.Val {
			prev, curr = curr, curr.Next
			continue
		}

		// Insertion can possibly happen starting from the head
		tmp := sentinel

		// Progress further if the elements are in sorted order
		for curr.Val > tmp.Next.Val {
			tmp = tmp.Next
		}

		// Now the curr node needs to be added right after `tmp`
		prev.Next = curr.Next
		curr.Next = tmp.Next
		tmp.Next = curr
		curr = prev.Next
	}

	return sentinel.Next
}
