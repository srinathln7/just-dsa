package main

func maximumRemovals(s string, p string, removable []int) int {

	// Key Idea: Brute force is to just remove the characters at the specified index in the removable slice
	// and then check if p is still a subsequence of the resulting string. An OPTIMAL approach can be found
	// by utilizing binary search. Since the question asks only for the max. k that we can remove so that
	// `p` is still a subsequence of `s`, we can run a binary search on the index of the removable array
	// and check if the resulting string is a subsequence

	n := len(s)
	var isSubsequence func(int) bool
	isSubsequence = func(k int) bool {

		// Form the removable set first
		removableSet := make(map[int]bool)
		for i := 0; i < k; i++ {
			removableSet[removable[i]] = true
		}

		// Check if the given string `p` is still a sub-sequence of the trimmed string `s`
		idx := 0
		for i := 0; i < n && idx < len(p); i++ {
			if !removableSet[i] && s[i] == p[idx] {
				idx++
			}
		}

		return idx == len(p)
	}

	// IMPORTANT:  Perform binary search on the removable array
	// To start with `l` represents the min. number of characters we can remove from the removable array
	// and `r` represents the max we can remove from the removable array
	l, r := 0, len(removable)
	for l <= r {
		mid := l + (r-l)/2
		if isSubsequence(mid) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	// `r` represents the max. we can remove while `p` still remains a subsequence of `s`
	return r
}
