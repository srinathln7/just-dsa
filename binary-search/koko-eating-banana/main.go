package main

import "math"

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, getMax(piles) // left and right pointers

	for l <= r {
		mid := l + (r-l)/2

		// Remember the range for k starts with -> [1...r]
		// Left portion represents slow eating but cannot finish all
		// Right portion represents fast eating but can finish all
		// We want to find min. k where can eat all
		if canEatAll(piles, h, mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l
}

func canEatAll(piles []int, totalHours, k int) bool {
	var eatingHours int
	for _, pile := range piles {
		eatingHours += int(math.Ceil(float64(pile) / float64(k)))
	}
	return eatingHours <= totalHours
}

func getMax(nums []int) int {
	max := nums[0]
	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
	}
	return max
}
