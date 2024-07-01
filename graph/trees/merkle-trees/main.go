package main

import (
	"crypto/sha256"
	"fmt"
)

type Node struct {
	val        string
	left_idx   int
	right_idx  int
	left_node  *Node
	right_node *Node
}

type Merkle struct {
	root *Node
}

func calcHash(tx []byte) string {
	hash := sha256.Sum256(tx)
	return fmt.Sprintf("%x", hash)
}

func newMerkleTree(txs [][]byte) *Merkle {
	n := len(txs)
	if n == 0 {
		return nil
	}

	l, r := 0, n-1
	root := buildMerkleTree(txs, l, r)
	return &Merkle{
		root: root,
	}
}

func buildMerkleTree(txs [][]byte, l, r int) *Node {
	if l == r {
		return &Node{
			val:       calcHash(txs[l]),
			left_idx:  l,
			right_idx: r,
		}
	}

	mid := l + (r-l)/2
	left := buildMerkleTree(txs, l, mid)
	right := buildMerkleTree(txs, mid+1, r)
	return &Node{
		val:        calcHash(append([]byte(left.val), []byte(right.val)...)),
		left_idx:   l,
		right_idx:  r,
		left_node:  left,
		right_node: right,
	}
}

func findParent(root, node *Node) *Node {
	if root == nil || node == nil {
		return nil
	}

	if root == node || (root.left_node == node || root.right_node == node) {
		return root
	}

	if left := findParent(root.left_node, node); left != nil {
		return left
	}

	return findParent(root.right_node, node)
}

func findSibling(root, node *Node) *Node {
	if root == nil || node == nil || root == node {
		return nil
	}

	parent := findParent(root, node)
	if parent.left_node == node {
		return parent.right_node
	}

	return parent.left_node
}

func findLeaf(root *Node, leafIdx int) *Node {
	if root == nil || leafIdx < root.left_idx || leafIdx > root.right_idx {
		return nil
	}

	if root.left_idx == leafIdx && root.right_idx == leafIdx {
		return root
	}

	midIdx := root.left_idx + (root.right_idx-root.left_idx)/2
	if leafIdx <= midIdx {
		return findLeaf(root.left_node, leafIdx)
	}

	return findLeaf(root.right_node, leafIdx)
}

func genProofs(root *Node, leafIdx int) []*Node {

	if root == nil || leafIdx < root.left_idx || leafIdx > root.right_idx {
		return nil
	}

	if root.left_node == nil && root.right_node == nil {
		return []*Node{root}
	}

	var result []*Node
	leaf := findLeaf(root, leafIdx)
	sibling := findSibling(root, leaf)
	result = []*Node{sibling}

	parent := findParent(root, leaf)
	for parent != root {
		sibling = findSibling(root, parent)
		result = append(result, sibling)
		parent = findParent(root, parent)
	}

	return result
}

func (mt *Merkle) verifyProof(rootHash, txHash string, txIdx int, proofs []*Node) bool {
	if txIdx < mt.root.left_idx || txIdx > mt.root.right_idx {
		return false
	}

	leaf := findLeaf(mt.root, txIdx)
	if leaf.val != txHash {
		return false
	}

	merkleHash := txHash
	curr := leaf
	for _, proof := range proofs {
		if curr.left_idx < proof.left_idx || curr.right_idx < proof.right_idx {
			merkleHash = calcHash(append([]byte(merkleHash), []byte(proof.val)...))
		} else {
			merkleHash = calcHash(append([]byte(proof.val), []byte(merkleHash)...))
		}
		curr.left_idx = min(curr.left_idx, proof.left_idx)
		curr.right_idx = max(curr.right_idx, proof.right_idx)
	}

	return merkleHash == rootHash && mt.root.val == rootHash
}

func main() {
	// Example transactions
	txs := [][]byte{
		[]byte("tx1"),
		[]byte("tx2"),
		[]byte("tx3"),
		[]byte("tx4"),
	}

	// Create a new Merkle tree
	mt := newMerkleTree(txs)

	// Display the root hash
	fmt.Println("Merkle Root Hash:", mt.root.val)

	// Generate proofs for a specific transaction (e.g., tx2)
	txIdx := 0 // Index of "tx2"
	proofs := genProofs(mt.root, txIdx)

	// Display the proofs
	fmt.Println("Proofs for tx2:")
	for _, proof := range proofs {
		fmt.Println(proof.val)
	}

	// Verify the proof for "tx2"
	txHash := calcHash(txs[txIdx])
	if mt.verifyProof(mt.root.val, txHash, txIdx, proofs) {
		fmt.Printf("Proof verified successfully for tx idx:%d \n", txIdx)
	} else {
		fmt.Printf("Proof verification failed for tx idx:%d \n", txIdx)
	}
}
