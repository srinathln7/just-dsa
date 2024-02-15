package main

type Node struct {
	val  string
	next *Node
	prev *Node
}

type BrowserHistory struct {
	list *Node
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{list: &Node{val: homepage}}
}

func (this *BrowserHistory) Visit(url string) {
	var newList *Node
	if this == nil {
		newList = &Node{val: url, next: nil, prev: nil}
		this.list = newList
		return
	}

	// MISTAKE : Did not disconnect the pointer to existing `this.next` that is pointing to nil to the newList
	// newList = &Node{val: url, next: nil, prev: this.list}
	// this.list = newList

	newList = &Node{val: url, next: nil, prev: this.list}
	this.list.next = newList
	this.list = this.list.next
}

func (this *BrowserHistory) Back(steps int) string {
	for i := steps; i > 0; i-- {
		// Encountered head
		if this.list.prev == nil {
			break
		}
		this.list = this.list.prev
	}
	return this.list.val
}

func (this *BrowserHistory) Forward(steps int) string {
	for i := steps; i > 0; i-- {
		// Encountered tail
		if this.list.next == nil {
			break
		}
		this.list = this.list.next
	}

	return this.list.val
}
