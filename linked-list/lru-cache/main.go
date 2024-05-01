package main

type Node struct {
	key  int
	val  int
	next *Node
	prev *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node

	head *Node // Sentinel node representing the least-recently-used (LRU)
	tail *Node // Sentinel node representing the most-recently-used  (MRU)
}

func Constructor(capacity int) LRUCache {

	lru := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     &Node{},
		tail:     &Node{},
	}

	// Connect head and tail sentinels
	lru.head.next, lru.tail.prev = lru.tail, lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.cache[key]; exists {

		// Move the node to most recently used
		this.Remove(node)
		this.Insert(node)

		return node.val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {

	if node, exists := this.cache[key]; exists {

		// Update the node val if it already exists
		node.val = value

		// Move the node to most recently used
		this.Remove(node)
		this.Insert(node)
	} else {

		node := &Node{key: key, val: value}
		this.cache[key] = node

		// Insert just before the tail sentinel
		this.Insert(node)

		// If cache exceeds capacity
		if len(this.cache) > this.capacity {
			lru := this.head.next
			this.Remove(lru)

			// Delete the lru from the cache to maintain the capacity
			delete(this.cache, lru.key)
		}
	}

}

// Helper functions

// Remove : Removes the node from the linkedlist
func (this *LRUCache) Remove(node *Node) {
	prev, next := node.prev, node.next
	prev.next, next.prev = next, prev
}

// Insert: Inserts the node just before the tail sentinel
func (this *LRUCache) Insert(node *Node) {
	prevTail := this.tail.prev
	prevTail.next = node
	node.prev = prevTail
	node.next = this.tail
	this.tail.prev = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
