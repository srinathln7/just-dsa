package main

import "math"

func minFlips(s string) int {

	// Key Idea: The key idea in this problem is to realize that for a given length of the string - say `n`
	// there are only two possibilities of alternating string "010101..." or "101010...". Now, the brute force
	// solution is to compare the difference between the original string and these two alternating strings, calculate
	// the `diff` and take the minimum of the two Then we perform a type 1 operation every character by character and
	// do the same and then return the solution. This will be O(n^2) solution. HOWEVER, we can achieve a linear time
	// complexity O(n) solution by simply turning this problem into fixed-sliding window problem if we append the string
	// to itself (length=2n) and then consider a window of size `n` and do the same operation of calculating the diff's.

	n := len(s)
	if n == 1 {
		return 0
	}

	// Append the string to itself to make a sliding window
	s += s

	a0, a1 := make([]byte, 2*n), make([]byte, 2*n)
	for i := 0; i < 2*n; i++ {
		if i&1 == 0 {
			a0[i] = '0'
			a1[i] = '1'
		} else {
			a0[i] = '1'
			a1[i] = '0'
		}
	}

	result := math.MaxInt32
	alt0, alt1 := string(a0), string(a1)

	// First window

	diff0, diff1 := calcDiff(s[:n], alt0[:n]), calcDiff(s[:n], alt1[:n])
	result = min(result, diff0, diff1)

	// Sliding window of size `k=n`
	for i := 1; i < n; i++ {

		if s[i-1] != alt0[i-1] {
			diff0--
		}

		if s[i-1] != alt1[i-1] {
			diff1--
		}

		// REMARK: Same as comparing s[i-1] with alt0[i+n-1] since we are appending s with itself at the begining
		if s[i+n-1] != alt0[i+n-1] {
			diff0++
		}

		if s[i+n-1] != alt1[i+n-1] {
			diff1++
		}

		result = min(result, diff0, diff1)
	}

	return result
}

func calcDiff(s, p string) int {
	if s == p {
		return 0
	}

	diff := 0
	for i := 0; i < len(s); i++ {
		if s[i] != p[i] {
			diff++
		}
	}

	return diff
}
