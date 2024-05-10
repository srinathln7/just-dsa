package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {

	// Key Idea: Simplest and most elegant possible solution to this problem is to solve this problem recursively.

	// Base case: Empty node or single node
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next

	// Swap the pairs recursively
	swapped := swapPairs(next.Next)

	// Point the head to the swapped pairs and the second node back to head
	head.Next, next.Next = swapped, head

	// Return the new head
	return next
}
