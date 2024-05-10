package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {

	// Key Idea: Maintain a sentinel node whose next pointer is pointing to it head. Traverse the linkedlist
	// and as we traverse, skip the iteration whenever the linkedlist value is equal to the target value

	result := &ListNode{Next: head}
	curr := result
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return result.Next
}
