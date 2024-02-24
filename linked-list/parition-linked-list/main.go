package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	var leftList, rightList *ListNode
	if head == nil {
		return nil
	}

	curr := head
	for curr != nil {
		// Add to left list
		if curr.Val < x {
			InsertToEnd(&leftList, curr.Val)
		} else {
			// Add to right list
			InsertToEnd(&rightList, curr.Val)
		}
		curr = curr.Next
	}

	return mergeList(leftList, rightList)
}

func InsertToEnd(head **ListNode, val int) {
	newNode := &ListNode{Val: val, Next: nil}

	// Add to head
	if *head == nil {
		*head = newNode
		return
	}

	curr := *head
	// Traverse until the last element
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = newNode
}

func mergeList(left, right *ListNode) *ListNode {

	if left == nil && right == nil {
		return nil
	}

	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	var mergedList *ListNode = left
	curr := left
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = right
	return mergedList
}
