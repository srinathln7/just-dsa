# Range query mutable

You are given an integer array `nums` and need to handle multiple queries of the following types:

1. Update the value of an element in `nums`.
2. Calculate the sum of the elements of `nums` between indices `left` and `right` inclusive where `left <= right`.

Implement the NumArray class with the following methods:

- `NumArray(nums []int)`: Initializes the object with the integer array `nums`.
- `Update(index int, val int)`: Updates the value of `nums[index]` to be `val`.
- `SumRange(left int, right int) int`: Returns the sum of the elements of `nums` between indices `left` and `right` inclusive.

## Intuition

To efficiently handle updates and range sum queries, we can use the segment tree data structure. The segment tree is a binary tree where each node represents a segment of the array. Each node stores the sum of elements within its range. With segment trees, both update and range sum queries can be performed in O(log n) time complexity.

## Approach

1. **Construction**: We will construct the segment tree recursively. At each node, we will divide the range into two halves and recursively build the left and right subtrees until we reach leaf nodes.
2. **Update**: To update a value at index `index`, we will recursively traverse the tree until we find the leaf node representing the index and update its value. Then, we will propagate the update upwards by updating the parent nodes' values accordingly.
3. **Sum Range Query**: To calculate the sum of elements between indices `left` and `right`, we will recursively traverse the tree. At each node, we will check if the current segment lies completely within the query range, lies completely outside the query range, or partially overlaps with the query range. Based on this, we will recursively traverse the left and right subtrees and calculate the sum accordingly.

## Time Complexity

- Construction: O(n)
- Update: O(log n)
- Range Sum Query: O(log n)

## Space Complexity

- O(n) for storing the segment tree.


## Remark

### Mistake made 

```
BIG MISTAKE:
func Constructor(nums []int) NumArray {
	l, r := 0, len(nums)-1
	// Base case - Leaf node
	if l == r {
		return NumArray{
			root: &SegmentTree{
				val:      nums[l],
				leftIdx:  l,
				rightIdx: r,
			},
		}
	}

	// Construct the left and right segment trees recursively
	m := l + (r-l)/2
	leftSegTree := Constructor(nums[:m+1])
	rightSegTree := Constructor(nums[m+1:])
	return NumArray{
		root: &SegmentTree{
			val:          leftSegTree.root.val + rightSegTree.root.val,
			leftIdx:      l,
			rightIdx:     r,
			leftSegTree:  leftSegTree.root,
			rightSegTree: rightSegTree.root,
		},
	}
}

```

The code snippet you provided does not work due to the way pointers are being handled in the Constructor function. The issue lies in the fact that when you call Constructor recursively to construct the left and right segment trees, you are accessing the root field of leftSegTree and rightSegTree directly, assuming they have a root field. However, since leftSegTree and rightSegTree are of type NumArray, they do not have a direct root field. The error causing the nil pointer dereference is  due to the incorrect assignment of pointers in the Constructor function. When constructing the segment tree recursively, the pointers leftSegTree and rightSegTree were being assigned the addresses of local variables instead of their actual values. This led to the pointers pointing to invalid memory locations once the local variables went out of scope, resulting in nil pointer dereference when accessing their fields later in the code.
By ensuring that the pointers are correctly initialized and assigned with the actual values of the left and right segment trees during construction, as shown in the revised code snippet provided earlier, you resolved the issue and prevented the nil pointer dereference from occurring. This adjustment ensures that valid references are stored in leftSegTree and rightSegTree, allowing proper traversal and manipulation of the segment tree