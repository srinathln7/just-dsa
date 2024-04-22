package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {

	// Key Idea: Reverse the middle portion of the linked list and then start merging the two lists alternatively

	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	secondHalf := slow.Next

	// Disconnect the first half from second half
	slow.Next = nil

	firstHalf, secondHalf := head, revList(secondHalf)

	// Merge the first half with second half by interleaving
	result := &ListNode{}
	tail := result
	for firstHalf != nil && secondHalf != nil {
		tail.Next = firstHalf
		firstHalf = firstHalf.Next
		tail = tail.Next

		tail.Next = secondHalf
		secondHalf = secondHalf.Next
		tail = tail.Next
	}

	if firstHalf != nil {
		tail.Next = firstHalf
	}

	head = result.Next
}

func revList(head *ListNode) *ListNode {
	var newHead, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = newHead
		newHead = head
		head = next
	}

	return newHead
}
