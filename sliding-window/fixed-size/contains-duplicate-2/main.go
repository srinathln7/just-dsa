package main

func containsNearbyDuplicate(nums []int, k int) bool {
	// i and j => distinct
	// |i-j| <= k  => At max we can have a window of size `k+1`
	// If we start from `j = 0`, `i` can be atmost `k` => We can have atmost `k+1` elements at a given time
	// If we account for the maximum window size of `k+1`, we can easily tackle lesser window sizes

	// Window set
	window := make(map[int]struct{})

	// Distinct indices are gauranteed through the use of a simple for loop
	for i, num := range nums {
		if _, exists := window[num]; exists {
			return true
		}

		window[num] = struct{}{}

		// Key Idea: If this start getting true when `i=k`, we already have `k+1` elements now
		// Before we proceed to the next iteration, we should remove the first element from the
		// window so that in the next iteration we can maintain a constant window of size `k+1`.
		// Had it been the case when the `k+1`st element was the same in that window, the fn would have
		// already returned true
		if len(window) == k+1 {
			delete(window, nums[i-k])
		}
	}

	return false
}
