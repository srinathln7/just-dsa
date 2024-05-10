package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {

	// Key Idea: O(n) time and O(1) space => Use two pointer approach to traverse until the middle of the linked list
	// Then reverse the middle of the linked list and then if the values are the same. IMPORTANT: Account for odd and
	// even number of elements in the linkedlist

	// Base cases - One element
	if head == nil || head.Next == nil {
		return true
	}

	// Only two elements
	if head.Next != nil && head.Next.Next == nil {
		return head.Val == head.Next.Val
	}

	// Two pointer technique to find the middle of the linkedlist
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Slow represents the middle node
	var secondHalf *ListNode

	// If linkedlist has even number of elements
	if fast == nil {
		secondHalf = slow
	} else {
		// Odd number of elements => fast != nil
		secondHalf = slow.Next
	}

	first, second := head, reverse(secondHalf)
	for first != nil && second != nil {
		if first.Val != second.Val {
			return false
		}

		first = first.Next
		second = second.Next
	}

	return true
}

func reverse(head *ListNode) *ListNode {

	var newHead *ListNode
	for head != nil {
		next := head.Next
		head.Next = newHead
		newHead = head
		head = next
	}

	return newHead
}
