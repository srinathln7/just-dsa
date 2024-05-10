package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {

	// Key Idea: We init two pointers `slow` and `fast` where the `fast` pointer moves twice as fast as the slow pointer
	// By the time, fast pointer reaches the end, slow pointer is at the middle
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
