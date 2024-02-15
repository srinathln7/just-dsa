package main

type MyStack struct {
	Queue []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.Queue = append(this.Queue, x)
}

func (this *MyStack) Pop() int {
	var deqElement int

	// Dequeue and Enqueue until the last element is encountered
	for i := 0; i < len(this.Queue)-1; i++ {
		deqElement = this.Queue[0]
		this.Queue = this.Queue[1:]
		this.Push(deqElement)
	}
	popElement := this.Queue[0]
	this.Queue = this.Queue[1:]
	return popElement
}

func (this *MyStack) Top() int {
	return this.Queue[len(this.Queue)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.Queue) == 0
}
