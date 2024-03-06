package main

type SegmentTree struct {
	val                       int
	leftIdx, rightIdx         int
	leftSegTree, rightSegTree *SegmentTree
}

type NumArray struct {
	root *SegmentTree
}

// BIG MISTAKE:
// func Constructor(nums []int) NumArray {
// 	l, r := 0, len(nums)-1
// 	// Base case - Leaf node
// 	if l == r {
// 		return NumArray{
// 			root: &SegmentTree{
// 				val:      nums[l],
// 				leftIdx:  l,
// 				rightIdx: r,
// 			},
// 		}
// 	}

// 	// Construct the left and right segment trees recursively
// 	m := l + (r-l)/2
// 	leftSegTree := Constructor(nums[:m+1])
// 	rightSegTree := Constructor(nums[m+1:])
// 	return NumArray{
// 		root: &SegmentTree{
// 			val:          leftSegTree.root.val + rightSegTree.root.val,
// 			leftIdx:      l,
// 			rightIdx:     r,
// 			leftSegTree:  leftSegTree.root,
// 			rightSegTree: rightSegTree.root,
// 		},
// 	}
// }

func Constructor(nums []int) NumArray {
	l, r := 0, len(nums)-1
	root := buildSegmentTree(nums, l, r)
	return NumArray{root: root}
}

func (this *NumArray) Update(index int, val int) {
	update(this.root, index, val)
}

func (this *NumArray) SumRange(left int, right int) int {
	return sumRange(this.root, left, right)
}

func buildSegmentTree(nums []int, l, r int) *SegmentTree {

	// Base case - Leaf node
	if l == r {
		return &SegmentTree{
			val:      nums[l],
			leftIdx:  l,
			rightIdx: r,
		}
	}

	m := l + (r-l)/2
	leftSegTree := buildSegmentTree(nums, l, m)
	rightSegTree := buildSegmentTree(nums, m+1, r)

	return &SegmentTree{
		val:          leftSegTree.val + rightSegTree.val,
		leftIdx:      l,
		rightIdx:     r,
		leftSegTree:  leftSegTree,
		rightSegTree: rightSegTree,
	}
}
func update(this *SegmentTree, index, val int) {
	// Base case - Leaf node
	if this.leftIdx == this.rightIdx {
		this.val = val
		return
	}

	midIdx := this.leftIdx + (this.rightIdx-this.leftIdx)/2
	switch {
	case index > midIdx:
		update(this.rightSegTree, index, val)
	default:
		update(this.leftSegTree, index, val)
	}

	// Once the sub-segments are updates, update the main root's sum as well
	this.val = this.leftSegTree.val + this.rightSegTree.val
}

func sumRange(this *SegmentTree, left, right int) int {
	// Base case: Leaf node
	if this.leftIdx == left && this.rightIdx == right {
		return this.val
	}

	midIdx := this.leftIdx + (this.rightIdx-this.leftIdx)/2
	switch {
	case left > midIdx:
		return sumRange(this.rightSegTree, left, right)
	case right <= midIdx:
		return sumRange(this.leftSegTree, left, right)
	default:
		return sumRange(this.leftSegTree, left, midIdx) + sumRange(this.rightSegTree, midIdx+1, right)
	}
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
