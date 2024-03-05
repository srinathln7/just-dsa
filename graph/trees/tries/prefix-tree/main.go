package main

// Trie: Consists of a root trie and 26 children representing `26` letters in alphabet
type Trie struct {
	root     *Trie
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{
		root:     &Trie{},
		children: [26]*Trie{},
	}
}

func (this *Trie) Insert(word string) {

	// First create a pointer to the trie's root/
	// REM: `this` points to the entire trie while `curr` points only to the trie's root
	// As we traverse the trie using `curr` pointer and insert every character until the end of the word
	// `curr` will reach the end while the original trie `this` will have the newly inserted word
	curr := this.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		// If the letter does not exist in the corresponding trie yet
		if curr.children[index] == nil {
			curr.children[index] = &Trie{}
		}

		// Traverse the trie and insert every character until the end of the word
		curr = curr.children[index]
	}

	curr.isEnd = true
}

func (this *Trie) Search(word string) bool {
	curr := this.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		// If a particular character in the word does not exist
		if curr.children[index] == nil {
			return false
		}

		curr = curr.children[index]
	}

	return curr.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this.root
	for i := 0; i < len(prefix); i++ {
		index := prefix[i] - 'a'
		// If a particular character in the prefix does not exist
		if curr.children[index] == nil {
			return false
		}
		curr = curr.children[index]
	}

	return true
}
