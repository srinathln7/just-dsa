package main

func maxSlidingWindow(nums []int, k int) []int {

	// Key Idea: As we slide through the window of constant size `k`, we need a data strucuture
	// where we can have access to the max. element of that window in constant time. Besides that
	// as we shrink through the window, we need to be able to add/delete from both front and back.
	// We use a deque (Double-ended queue) to accomplish this where the deque stores the element itself.

	var deque, result []int

	// As we move the right pointer `r` forward, we can access the index of the element outside the current window by `r-k``
	cleanDeque := func(r int) {
		// Dequeue from the FRONT: Drop the first element to remove element outside the current window.
		n := len(deque)
		if n > 0 && deque[0] == r-k {
			deque = deque[1:]
			n = len(deque) // Update the length
		}

		// Keep Dequeueing from the BACK if the number we are adding is greater than the rest of the numbers in the current window
		// By this way, we ensure constant time access to maximum element in a window
		for n > 0 && nums[r] >= nums[deque[n-1]] {
			deque = deque[:len(deque)-1]
			n = len(deque) // Update the length
		}
	}

	for i := 0; i < len(nums); i++ {

		cleanDeque(i)

		// Deque only stores the index and not the actual element
		deque = append(deque, i)

		// Once you have processed atleast `k-1` elements  => we have finished processing the first window and will start processing all subsequent windows
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}
