package main

func subarraySum(nums []int, k int) int {
	// Key Idea: To keep track of every prefix sums in a hash map and check if the difference b/w k and the current sum
	// in the iteration exists or not. If it exists, it means there is a subarry with `sum=k` and we increment the count
	// by the value associated with that sum in the hash map

	prefixSum := make(map[int]int)

	// Base case: When a single element equals `k`
	prefixSum[0] = 1

	var count, cumSum int
	for _, num := range nums {
		cumSum += num
		if val, exists := prefixSum[cumSum-k]; exists {
			count += val
		}

		prefixSum[cumSum]++
	}

	return count
}
