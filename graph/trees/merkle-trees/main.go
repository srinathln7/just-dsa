package main

import (
	"crypto/sha256"
	"fmt"
)

type Node struct {
	Hash     string
	Left     *Node
	Right    *Node
	LeftIdx  int
	RightIdx int
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

	// Leaf nodes
	if l == r {
		return &Node{
			Hash:     calcHash(txs[l]),
			LeftIdx:  l,
			RightIdx: r,
		}
	}

	mid := l + (r-l)/2
	left := buildMerkleTree(txs, l, mid)
	right := buildMerkleTree(txs, mid+1, r)
	return &Node{
		Hash: calcHash(append([]byte(left.Hash), []byte(right.Hash)...)),
		//Hash:     calcHash(fmt.Append([]byte(left.Hash), []byte(right.Hash))),
		Left:     left,
		Right:    right,
		LeftIdx:  l,
		RightIdx: r,
	}
}

func findParent(root, node *Node) *Node {

	if root == nil || node == nil {
		return nil
	}

	if root == node || root.Left == node || root.Right == node {
		return root
	}

	if left := findParent(root.Left, node); left != nil {
		return left
	}

	return findParent(root.Right, node)
}

func findSibling(root, node *Node) *Node {

	if root == nil || node == nil || root == node {
		return nil
	}

	parent := findParent(root, node)
	if parent == nil {
		return nil
	}

	if parent.Left == node {
		return parent.Right
	}

	return parent.Left
}

func findLeaf(root *Node, leafIdx int) *Node {
	if root == nil || (leafIdx < root.LeftIdx || leafIdx > root.RightIdx) {
		return nil
	}

	if root.LeftIdx == leafIdx && root.RightIdx == leafIdx {
		return root
	}

	midIdx := root.LeftIdx + (root.RightIdx-root.LeftIdx)/2
	if leafIdx <= midIdx {
		return findLeaf(root.Left, leafIdx)
	}
	return findLeaf(root.Right, leafIdx)
}

func (mt *Merkle) genMerkleProofs(txIdx int) []*Node {
	var proofs []*Node
	leaf := findLeaf(mt.root, txIdx)
	proofs = append(proofs, findSibling(mt.root, leaf))

	parent := findParent(mt.root, leaf)
	for parent != mt.root {
		proofs = append(proofs, findSibling(mt.root, parent))
		parent = findParent(mt.root, parent)
	}
	return proofs
}

func (mt *Merkle) verifyMerkleProofs(txHash string, txIdx int, proofs []*Node) bool {

	leaf := findLeaf(mt.root, txIdx)
	if leaf == nil {
		return false
	}

	if leaf.Hash != txHash {
		return false
	}

	merkleHash := txHash

	curr := leaf
	for _, proof := range proofs {
		if curr.LeftIdx < proof.LeftIdx && curr.RightIdx < proof.RightIdx {
			merkleHash = calcHash(append([]byte(merkleHash), []byte(proof.Hash)...))
		} else {
			merkleHash = calcHash(append([]byte(proof.Hash), []byte(merkleHash)...))
		}

		// IMPORTANT: To update the index
		curr.LeftIdx = min(curr.LeftIdx, proof.LeftIdx)
		curr.RightIdx = max(curr.RightIdx, proof.RightIdx)
	}

	return merkleHash == mt.root.Hash
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
	fmt.Println("Merkle Root Hash:", mt.root.Hash)

	// Generate proofs for a specific transaction (e.g., tx2)
	txIdx := 3 // Index of "tx2"
	proofs := mt.genMerkleProofs(txIdx)

	// Display the proofs
	fmt.Printf("Proofs for tx:%d\n", txIdx)
	for _, proof := range proofs {
		fmt.Println(proof.Hash)
	}

	// Verify the proof for "tx2"
	txHash := calcHash(txs[txIdx])
	if mt.verifyMerkleProofs(txHash, txIdx, proofs) {
		fmt.Printf("Proof verified successfully for tx%d \n", txIdx)
	} else {
		fmt.Printf("Proof verification failed for tx:%d \n", txIdx)
	}
}
