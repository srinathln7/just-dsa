package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {

	// Key Idea: Left spectrum represents the good versions and right spectrum maintains the bad version. We keep moving
	// both the pointers towards each other until the pointers cross or matches. The minute the left pointer shifts from
	// good spectrum to bad spectrum is out first bad version.

	l, r := 1, n
	for l <= r {
		mid := l + (r-l)/2
		if isBadVersion(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l
}
