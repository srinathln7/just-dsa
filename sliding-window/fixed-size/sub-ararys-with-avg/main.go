package main

func numOfSubarrays(arr []int, k int, threshold int) int {
	var count, avg, windowSum int

	// First, calculate the sum uptil for first `k-1` elements
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}

	// Calculate avg for the first window
	avg = windowSum / k
	if avg >= threshold {
		count++
	}

	// Note the loop exits before calculating avg for the last window
	for i := k; i < len(arr); i++ {
		// Slide the window
		windowSum = windowSum + arr[i] - arr[i-k]

		avg = windowSum / k
		if avg >= threshold {
			count++
		}
	}

	return count
}
