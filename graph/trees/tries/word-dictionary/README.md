# Word Dictionary

## Intuition
We can solve this problem using a trie data structure. Each node in the trie represents a character, and the presence of a word is indicated by marking the last node as the end of the word. To support wildcard search with '.', we can perform a depth-first search (DFS) on the trie, checking each character in the word and branching into all possible characters if '.' is encountered.

## Approach
1. Define a TrieNode struct with an array of pointers to child nodes representing each character and a boolean flag to indicate if it's the end of a word.
2. Implement the WordDictionary struct with a root node pointing to the trie's root.
3. For adding a word, traverse the trie starting from the root and create new nodes for each character if they don't exist.
4. Mark the last node as the end of the word.
5. For searching a word, perform a DFS starting from the root of the trie.
6. At each level of the trie, check if the current character matches the corresponding character in the word.
7. If '.', branch into all possible characters. If a match is found, continue DFS.
8. If the end of the word is reached and the node is marked as the end of a word, return true.
9. If the DFS completes without finding a match, return false.

## Time Complexity
- Adding a word: O(M), where M is the length of the word.
- Searching a word: O(N), where N is the number of characters in the word.

## Space Complexity
- The space complexity for the trie structure is O(N*M), where N is the number of words and M is the average length of the words.