package main

import "sort"

func threeSum(nums []int) [][]int {

	// KEY IDEA: Since we have to find 3 elements with distinct indices, the brute force would be to loop over three loops
	// at a time resulting in O(n^3) solution. We can essentially turn this into a two sum problem using a hash map.
	// Range over each `num` in the array and find the two sum solution for array nums[i+1:], -num since collectively
	// they should sum up to zero.  Since we should avoid duplicates, it is important to sort the array first.

	var result [][]int

	// IMPORTANT: Since, we are going to use a hash map of key with type [3]int, it is important to sort the array first
	// to ensure no duplicates are stored
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] <= nums[j]
	})

	resMap := make(map[[3]int]struct{})
	for i, num := range nums {

		// Skip duplicates - Avoid duplicates condition since the array is is already sorted
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		sums2 := twoSum(nums[i+1:], -num)
		for _, res := range sums2 {
			if _, exists := resMap[[3]int{num, res[0], res[1]}]; !exists && len(res) != 0 {
				result = append(result, append([]int{num, res[0]}, res[1]))
				resMap[[3]int{num, res[0], res[1]}] = struct{}{}
			}
		}
	}

	return result
}

func twoSum(nums []int, target int) [][2]int {
	var result [][2]int
	sumMap := make(map[int]struct{})
	for _, num := range nums {
		if _, exists := sumMap[target-num]; exists {
			result = append(result, [2]int{num, target - num})
		}
		sumMap[num] = struct{}{}
	}

	return result
}
