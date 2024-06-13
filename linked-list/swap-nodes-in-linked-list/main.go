package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {

	// Calculate the length of linkedlist
	n, curr := 0, head
	for curr != nil {
		n++
		curr = curr.Next
	}

	// Get the kth node from the beginning
	first := head
	for i := 1; i < k; i++ {
		first = first.Next
	}

	// `k` the node from the end => `n-k+1`st node from the beginning
	second := head
	for i := 1; i < n-k+1; i++ {
		second = second.Next
	}

	// No need to change pointers => Just swap the values
	first.Val, second.Val = second.Val, first.Val
	return head
}
