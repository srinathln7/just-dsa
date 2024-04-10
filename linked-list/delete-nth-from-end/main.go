package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	// Key Idea: Keep track of two pointers `fast` and `slow` and move the slow pointer by a delay of `n` steps
	// By the time fast pointer reaches the end slow pointer would be pointing to the `nth` node from the end

	dummy := &ListNode{Next: head}
	fast, slow := head, dummy

	// Move fast pointer by `n` steps
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// Now start moving both fast and slow pointer until `fast` pointer reaches the end
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// Skip the `nth` node from the end
	slow.Next = slow.Next.Next
	return dummy.Next
}
