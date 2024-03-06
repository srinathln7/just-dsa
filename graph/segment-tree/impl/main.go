package main

import "fmt"

type SegmentTree struct {
	val, leftIdx, rightIdx int
	leftSeg, rightSeg      *SegmentTree
}

func newSegmentTree(nums []int, l, r int) *SegmentTree {

	// Reached the leaf node
	if l == r {
		return &SegmentTree{val: nums[l], leftIdx: l, rightIdx: r}
	}

	m := l + (r-l)/2

	// Build the left and right segment trees recursively
	leftSegmentTree := newSegmentTree(nums, l, m)
	rightSegmentTree := newSegmentTree(nums, m+1, r)

	return &SegmentTree{
		val:      leftSegmentTree.val + rightSegmentTree.val,
		leftIdx:  l,
		rightIdx: r,
		leftSeg:  leftSegmentTree,
		rightSeg: rightSegmentTree,
	}

}

func (seg *SegmentTree) update(idx, val int) {
	// If it is a leaf node i.e only one node, update the value at the index
	if seg.leftIdx == seg.rightIdx {
		seg.val = val
		return
	}

	// Similar to binary Search : if the specified index to be updated lies to the right of the mid index
	// of the tree, then we search and update the right segment. Otherwise, we update the left segment
	midIdx := seg.leftIdx + (seg.rightIdx-seg.leftIdx)/2
	switch {
	case idx > midIdx:
		seg.rightSeg.update(idx, val)
	default:
		seg.leftSeg.update(idx, val)
	}

	// Once the value at the specified index is updated, we need to update all the corresponding root query (sum)as well
	seg.val = seg.leftSeg.val + seg.rightSeg.val
}

// rangeQuery: In this case, is the sum of all ranges
func (seg *SegmentTree) rangeQuery(startIdx, endIdx int) int {

	// Base case: Destination reached - Leaf node
	if seg.leftIdx == startIdx && seg.rightIdx == endIdx {
		return seg.val
	}

	midIdx := seg.leftIdx + (seg.rightIdx-seg.leftIdx)/2
	switch {
	// If the start index of the query range is greater than the mid, then our target lies in the right segment
	case startIdx > midIdx:
		return seg.rightSeg.rangeQuery(startIdx, endIdx)

	// If the end index of the query range is lesser than or equal to the mid, then our target lies in the left segment
	case endIdx <= midIdx:
		return seg.leftSeg.rangeQuery(startIdx, endIdx)

	// When there is a overlap b/w the two segments we build the two segments and calculate the query (`sum` function) in this case
	default:
		return seg.rangeQuery(startIdx, midIdx) + seg.rangeQuery(midIdx+1, endIdx)
	}

}

func main() {
	nums := []int{1, 3, 5, 7, 9, 11}
	root := newSegmentTree(nums, 0, len(nums)-1)

	fmt.Println("Sum from index 1 to 3:", root.rangeQuery(1, 3))

	root.update(2, 10)
	fmt.Println("Sum from index 1 to 3 after update:", root.rangeQuery(1, 3))

	// Output
	// Sum from index 1 to 3: 15
	// Sum from index 1 to 3 after update: 20
}
