package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// Key Idea: There is a very elegant solution to this problem called as Recurring cyclic pattern.
	// We maintain two pointers to the two linked lists and keep traversing them. As we move and reach
	// the end of the linked list, we move this pointer the other linked list and start traversing. The
	// the intersection point is when the two nodes meet

	l1, l2 := headA, headB
	for l1 != l2 {
		if l1 == nil {
			l1 = headB
		} else {
			l1 = l1.Next
		}

		if l2 == nil {
			l2 = headA
		} else {
			l2 = l2.Next
		}
	}

	return l1
}
