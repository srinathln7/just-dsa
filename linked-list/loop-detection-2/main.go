package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {

	// Edge case
	if head == nil || head.Next == nil {
		return nil
	}

	// Key Idea: We maintain two pointers fast and slow pointer. `fast` pointer moves twice as fast as the slow pointer
	// If the linkedlist ever has a cycle, then eventually fast and slow pointer will meet. Else there is no cycle
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		// If fast's Next pointer ever becomes nil during this process then there is no cycle
		if fast == nil || fast.Next == nil {
			return nil
		}

		// This is the meeting point
		if fast == slow {
			break
		}
	}

	// We can mathematically demonstrate that the meeting point, where the tortoise and hare intersect,
	// is equidistant from both the starting point of the cycle and the head of the linked list. So to detect
	// the point where the cycle begins we move slow to `head` and now traverse slow and fast pointers at equal pace

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return fast
}
