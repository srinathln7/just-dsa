package main

func numOfSubarrays(arr []int, k int, threshold int) int {
	var count, avg, windowSum int

	// First, calculate the avg for the first window
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}
	avg = windowSum / k
	if avg >= threshold {
		count++
	}

	// Slide the window
	for i := k; i < len(arr); i++ {
		windowSum = windowSum + arr[i] - arr[i-k]
		avg = windowSum / k
		if avg >= threshold {
			count++
		}
	}

	return count
}
