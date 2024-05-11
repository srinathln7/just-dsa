package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	// Key Idea: We need to arrive at the starting position of the left and maintain a pointer that is pointing to the
	// list that needs to be reversed. Number of elements in the portion to be reversed equal `right-left+1`

	// Base cases
	if head == nil || left == right {
		return head
	}

	sentinel := &ListNode{Next: head}
	leftPrev, curr := sentinel, head
	for i := 0; i < left-1; i++ {
		curr = curr.Next
		leftPrev = leftPrev.Next
	}

	// Now `curr` node is at left position and `leftPrev` is the node just before the left

	// `newHead` will represent the list that needs to be reversed
	var newHead *ListNode

	// We reverse each element and hence run the loop `r-l+1` times equal to the number of elements in the portion to be reversed
	for i := 0; i < right-left+1; i++ {
		next := curr.Next
		curr.Next = newHead
		newHead = curr
		curr = next
	}

	// After this iteration `curr` node point to the remainder of the list
	leftPrev.Next.Next = curr
	leftPrev.Next = newHead
	return sentinel.Next
}
