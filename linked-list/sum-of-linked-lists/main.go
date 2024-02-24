package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sum *ListNode
	var val, carry int

	for l1 != nil || l2 != nil || carry > 0 {
		val = carry
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}

		carry = val / 10
		val %= 10

		insertToEnd(&sum, val)
	}

	return sum
}

func insertToEnd(head **ListNode, val int) {
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
