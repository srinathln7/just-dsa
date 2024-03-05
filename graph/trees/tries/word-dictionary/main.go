package main

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: &TrieNode{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if curr.children[index] == nil {
			curr.children[index] = &TrieNode{}
		}
		curr = curr.children[index]
	}

	curr.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {

	var dfs func(*TrieNode, int) bool
	dfs = func(node *TrieNode, index int) bool {
		// Base condition
		if index == len(word) {
			return node.isEnd
		}

		ch := word[index]
		switch ch {

		// Wild card pattern
		case '.':
			for _, childNode := range node.children {
				if childNode != nil && dfs(childNode, index+1) {
					return true
				}
			}
			return false

		// Standard characters
		default:
			childNode := node.children[ch-'a']
			if childNode == nil {
				return false
			}
			return dfs(childNode, index+1)
		}
	}

	return dfs(this.root, 0)
}
