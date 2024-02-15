package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	head *ListNode
}

func Constructor() MyLinkedList {
	// Mistake: You're initializing the head of the linked list to an empty node, but the head pointer itself should be nil initially.
	// Otherwise, it creates an extra node that doesn't represent an actual element in the linked list.
	//return MyLinkedList{head: &ListNode{}}

	return MyLinkedList{head: nil}
}

func (this *MyLinkedList) getLength() int {
	var n int
	curr := this.head
	for curr != nil {
		n++
		curr = curr.Next
	}

	return n
}

func (this *MyLinkedList) isValidIndex(index int) bool {
	curr := this.head
	if curr == nil || index < 0 {
		return false
	}
	return index < this.getLength()
}

func (this *MyLinkedList) Get(index int) int {
	if !this.isValidIndex(index) {
		return -1
	}

	curr := this.head
	idx := 0
	for curr != nil {
		if idx == index {
			break
		}
		idx++
		curr = curr.Next
	}

	return curr.Val
}

func (this *MyLinkedList) AddAtHead(val int) {

	// Add to empty list
	if this.head == nil {
		newNode := &ListNode{Val: val, Next: nil}
		this.head = newNode
		return
	}

	// Non-empty list
	newNode := &ListNode{Val: val, Next: this.head}
	this.head = newNode
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &ListNode{Val: val, Next: nil}
	// Add to Empty list
	if this.head == nil {
		this.head = newNode
		return
	}

	// Non.empty list
	curr := this.head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	// Index validation
	if index > this.getLength() {
		return
	}

	// Special cases
	if index == this.getLength() {
		this.AddAtTail(val)
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	}

	curr := this.head

	// To add element before the i^th index, traverse until `i-1`
	// 0 -> 1 -> 2 -> 3 -> 4
	for i := 0; i < index-1; i++ {
		curr = curr.Next
	}
	newNode := &ListNode{Val: val, Next: nil}
	newNode.Next = curr.Next
	curr.Next = newNode
}

func (this *MyLinkedList) DeleteAtIndex(index int) {

	if index >= this.getLength() {
		return
	}

	// Special case Delete at beginning
	// 0 -> 1-> 2 -> 3

	if index == 0 {
		this.head = this.head.Next
		return
	}

	curr := this.head
	for i := 0; i < index-1; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next

}
