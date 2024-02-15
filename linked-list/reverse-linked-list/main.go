package main

// * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	linkedlistItr := IreverseList(head)
	linkedlistRecur := RreverseList(head)
	if linkedlistItr != linkedlistRecur {
		panic("something went wrong")
	}

	return linkedlistItr
}

func IreverseList(head *ListNode) *ListNode {
	var next, newHead *ListNode
	for head != nil {
		next = head.Next
		head.Next = newHead
		newHead = head
		head = next
	}
	return newHead
}

func RreverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := RreverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
