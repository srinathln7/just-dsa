package main

import (
	"crypto/sha256"
	"fmt"
)

type Node struct {
	children map[string]*Node
	Val      string
	Hash     string
}

type MPT struct {
	root *Node
}

func newMPT() *MPT {
	return &MPT{
		root: &Node{
			children: make(map[string]*Node),
		},
	}
}

func (mpt *MPT) insert(key, val string) {

	curr := mpt.root
	for _, ch := range key {
		chStr := string(ch)
		if _, exists := curr.children[chStr]; !exists {
			curr.children[chStr] = &Node{children: make(map[string]*Node)}
		}

		curr = curr.children[chStr]
	}

	updateHashes(mpt.root)
	curr.Val = val
}

func (mpt *MPT) search(key string) (string, bool) {
	curr := mpt.root
	for _, ch := range key {
		chStr := string(ch)
		if _, exists := curr.children[chStr]; !exists {
			return "", false
		}
		curr = curr.children[chStr]
	}

	return curr.Val, true
}

func updateHashes(node *Node) string {
	if node == nil {
		return ""
	}

	if len(node.children) == 0 {
		node.Hash = calcHash(node.Val)
		return node.Hash
	}

	var combinedHashes string
	for _, child := range node.children {
		combinedHashes += updateHashes(child)
	}

	node.Hash = calcHash(combinedHashes)
	return node.Hash
}

func calcHash(val string) string {
	hash := sha256.Sum256([]byte(val))
	return fmt.Sprintf("%x", hash)
}

func main() {
	mpt := newMPT()
	fmt.Printf("MPT Root hash: %s\n", mpt.root.Hash)

	mpt.insert("apple", "A fruit")
	fmt.Printf("MPT Root hash: %s\n", mpt.root.Hash)

	mpt.insert("app", "Application")
	fmt.Printf("MPT Root hash: %s\n", mpt.root.Hash)

	mpt.insert("banana", "A fruit")
	fmt.Printf("MPT Root hash: %s\n", mpt.root.Hash)

	if value, found := mpt.search("apple"); found {
		fmt.Println("apple:", value)
	} else {
		fmt.Println("apple: not found")
	}

	if value, found := mpt.search("app"); found {
		fmt.Println("app:", value)
	} else {
		fmt.Println("app: not found")
	}

	if value, found := mpt.search("banana"); found {
		fmt.Println("banana:", value)
	} else {
		fmt.Println("banana: not found")
	}

	if value, found := mpt.search("apricot"); found {
		fmt.Println("apricot:", value)
	} else {
		fmt.Println("apricot: not found")
	}
}
