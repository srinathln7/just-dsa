package main

func replaceElements(arr []int) []int {

	n, rightMax := len(arr), -1
	for i := n - 1; i >= 0; i-- {
		newMax := findRightMax(rightMax, arr[i])
		arr[i] = rightMax
		rightMax = newMax
	}

	return arr
}

func findRightMax(currMax, element int) int {
	if element > currMax {
		currMax = element
	}
	return currMax
}
