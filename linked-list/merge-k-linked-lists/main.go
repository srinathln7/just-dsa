package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {

	// Key Idea: A straight forward approach would be to just merge two lists and then keep merging that list with the following lists
	// subsequently. However, a faster approach would be to divide the entire `k` lists in to one and keep recursively dividing it using
	// ideas from merge sort. By this way, we can have a O(n.log k) algorithm

	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	if len(lists) == 2 {
		return mergeTwoLists(lists[0], lists[1])
	}

	l, r := 0, len(lists)
	mid := l + (r-l)/2

	firstHalf := mergeKLists(lists[:mid])
	secondHalf := mergeKLists(lists[mid:])

	return mergeTwoLists(firstHalf, secondHalf)
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	dummy := &ListNode{}
	tail := dummy

	for list1 != nil && list2 != nil {

		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}

		// Progress tail further
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	}

	if list2 != nil {
		tail.Next = list2
	}

	return dummy.Next
}
