package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {

	// Key Idea: Reverse the middle portion of the linkedlist and have a pointer `l2` pointing to the head of this linkedlist.
	// Now set l1 at the head of the linkedlist and traverse until l2 is empty. Note len(l2) = len(l1)/2 i.e. len(l2) < len(l1)
	// Calculate the sum and update max sum accorindgly.
	var maxSum int

	// Find the middle portion of the linkedlist
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse the mid-portion to extract the twin sum and travese till the end of this linked list
	// REM MISTAKE: `l2.Next != nil` will ignore the last element
	l1, l2 := head, reverseLinkedList(slow)
	for l2 != nil {
		maxSum = max(maxSum, l1.Val+l2.Val)
		l1 = l1.Next
		l2 = l2.Next
	}

	return maxSum
}

func reverseLinkedList(head *ListNode) *ListNode {
	var newHead, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = newHead
		newHead = head
		head = next
	}
	return newHead
}
