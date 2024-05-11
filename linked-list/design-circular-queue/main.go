package main

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

type MyCircularQueue struct {
	head     *Node // Sentinel nodes
	tail     *Node // Sentinel nodes
	length   int
	capacity int
}

func Constructor(k int) MyCircularQueue {
	cq := MyCircularQueue{
		head:     &Node{},
		tail:     &Node{},
		capacity: k,
	}

	// Connect the head and tail since the queue is circular
	cq.head.Prev, cq.head.Next = cq.tail, cq.tail
	cq.tail.Prev, cq.tail.Next = cq.head, cq.head

	return cq
}

func (this *MyCircularQueue) EnQueue(value int) bool {

	// Enqueue to the end
	if this.length == this.capacity {
		return false
	}

	node := &Node{Val: value, Prev: this.tail.Prev, Next: this.tail}
	this.tail.Prev.Next = node
	this.tail.Prev = node
	this.length++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	// Dequeue from the front
	if this.length == 0 {
		return false
	}

	next, prev := this.head.Next.Next, this.head
	next.Prev, prev.Next = this.head, next
	this.length--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.length > 0 {
		return this.head.Next.Val
	}

	return -1
}

func (this *MyCircularQueue) Rear() int {
	if this.length > 0 {
		return this.tail.Prev.Val
	}
	return -1
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.length == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.length == this.capacity
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
