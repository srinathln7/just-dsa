# Trie (Prefix Tree) Implementation

## Intuition
A trie is a tree-like data structure used for efficient retrieval of keys in a dataset of strings. It is commonly used in applications involving autocomplete features, spell checkers, and IP routing tables. This implementation of a trie supports basic operations such as insertion, search, and prefix search.

## Approach
1. Initialize a trie with a root node and an array of 26 children, each representing a letter in the alphabet.
2. Implement insertion by traversing the trie from the root to the corresponding node for each character in the word to be inserted. Mark the last node as an end node.
3. Implement search by traversing the trie from the root to the corresponding node for each character in the word. Return true if the end node is reached.
4. Implement prefix search by traversing the trie from the root to the corresponding node for each character in the prefix. Return true if the prefix is found in the trie.

## Time Complexity
1. Insertion: O(m), where m is the length of the word to be inserted.
2. Search: O(m), where m is the length of the word to be searched.
3. Prefix Search: O(p), where p is the length of the prefix.

## Space Complexity
The space complexity of the trie depends on the number of words inserted and the length of the words. In this implementation, the space complexity is O(n*m), where n is the number of words and m is the average length of the words.

