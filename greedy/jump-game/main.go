package main

func canJump(nums []int) bool {

	// Key Idea: Let us start from backwards where our goal post is at the last index `n-1`.
	// Say, we are just one step behind the goal post i.e. at index `n-2`, then the max jump
	// we can make from that index is `idx + nums[idx]` and if we can reach our goal post and beyond
	// from that then we can shift our goal post to this spot. We continue this and check if we can shift the goal post
	// to the first spot

	n := len(nums)
	goal := n - 1

	for i := n - 2; i >= 0; i-- {
		// If we can reach the goal post or beyond from the current index with our jump, then we can shift our goal post
		currJump := i + nums[i]
		if currJump >= goal {
			goal = i // Shift the goal post
		}
	}

	// Check if the goal post is at the beginning
	return goal == 0
}
