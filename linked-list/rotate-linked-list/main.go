package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {

	// Base case
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// Determine the length of the linked list
	n, curr := 0, head
	for curr != nil {
		n++
		curr = curr.Next
	}

	// Determine the number of times we have to rotate
	k = k % n

	// No need to rotate
	if k == 0 {
		return head
	}

	// Find the (n-k)th node from the begininning
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	// Traverse until the last node
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// `fast` now represents the last node
	newHead := slow.Next
	slow.Next = nil
	fast.Next = head

	return newHead
}
