package main

func jump(nums []int) int {

	// Key Idea: We adopt a greedy approach. To find the min. number of steps needed to reach from idx 0 to idx `n-1`
	// we can deploy a simple 1D-BFS to output the shortest path length.

	n, l, r := len(nums), 0, 0
	count := 0
	// Run this until we reach the destination which is the last idx
	for r < n-1 {
		maxJump := 0
		// Run the loop within the incl range [l, r]
		for i := l; i <= r; i++ {
			currJump := i + nums[i]

			// maxJump represents the farthest we can jump from the current idx
			maxJump = max(maxJump, currJump)
		}

		// Since it is gauranteed a solution always exists, the min. step we can jump is `l=r+1`
		l = r + 1
		r = maxJump
		count++
	}

	return count
}
