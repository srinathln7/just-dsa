package main

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	// Assume a Linkedlist with length N, removing `n`th from the end is equivalent  to removing `N-n+1`st node from the beginning
	// Constraints => n <= N,  N-n+1 <= N => n >= 1
	if n == 0 {
		return head
	}
	curr := head
	var N int
	for curr != nil {
		N++
		curr = curr.Next
	}
	if n > N || N == 0 {
		return nil
	}
	if N == 1 && n == 1 {
		return nil
	}
	k := N - n + 1
	if k == 1 {
		head = head.Next
		return head
	}
	curr = head
	for i := 1; i < k-1; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next
	return head
}
