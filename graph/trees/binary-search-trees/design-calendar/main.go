package main

// Binasry Search Tree
type BST struct {
	start, end  int
	Left, Right *BST
}

type MyCalendar struct {
	root *BST
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(start int, end int) bool {

	// Key Idea: Once the first booking is made, we must to always ensure that the next subsequent booking intervals never overlaps.
	// This essentially means - either your new booking's start index is be greater than the root and all sub-roots end index OR
	// your new bookings end index is lesser than the root and all sub-trees start index. => This means we need a data structure
	// where all left sub_tree <= root < right sub_tree => BST
	if this.root == nil {
		this.root = &BST{start: start, end: end}
		return true
	}
	return book(this.root, start, end)
}

func book(root *BST, start, end int) bool {

	switch {

	// When my new booking end index is lesser than the current root's start index
	// Add a left sub-tree i.e. new booking (BST) to the left of the root
	case end <= root.start:

		// If the left child does not exist - add it immediately
		if root.Left == nil {
			root.Left = &BST{start: start, end: end}
			return true
		}

		// IMPORTANT: If a left child already exists : recurse to ensure that the new booking does not overlap in the left sub-tree
		return book(root.Left, start, end)

	// When my new booking start index is greater than the current root's end index
	// Add a right sub-tree i.e.new booking (BST) to the right of the root
	case start >= root.end:
		if root.Right == nil {
			root.Right = &BST{start: start, end: end}
			return true
		}

		// IMPORTANT: If a right child already exists : recurse to ensure that the new booking does not overlap in the right sub-tree
		return book(root.Right, start, end)

	default:
		return false
	}
}
