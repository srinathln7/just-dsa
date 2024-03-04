package main

func maxTurbulenceSize(arr []int) int {

	// Key Idea: To use sliding window approach of variable size with two pointers `l` and `r`
	// We set `l` and `r` pointers initially at the beginning. And then traverse the `r` pointer along the array if array is greater than 1.
	// If a turbulent subarray condition is met, we update the maxLen accordingly otherwise we update the `l` pointer to
	// the current position of `r` pointer and continue our search until we meet the end of the array

	n := len(arr)
	l, r := 0, 0
	maxLen := 1

	// When we have at least 2 elements in the array
	for r < n-1 {
		// Check the relationship between the current and next elements to determine if they form a turbulent subarray
		switch {
		// If the two adjacent elements are equal, it can never form a turbulent subarray
		// and so we move the left pointer `l` to the next position of the right pointer
		case arr[r] == arr[r+1]:
			l = r + 1

		// If the conditions for a turbulent subarray are met, update the maximum length
		// We do `r-l+2` because `r-l+1` elements denotes the number of elements b/w left and right index
		// It does not account for element `r+1` which is also a part of the turbulent subarray.
		// Hence `(r-l+1)+1`
		case r == l || (arr[r-1] < arr[r] && arr[r] > arr[r+1]) || (arr[r-1] > arr[r] && arr[r] < arr[r+1]):
			maxLen = max(maxLen, r-l+2)

		// If the current elements don't form a turbulent subarray, update the left pointer `l` to the current position of the right pointer
		default:
			l = r
		}

		r++
	}

	return maxLen
}
