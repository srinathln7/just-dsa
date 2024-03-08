package main

import "fmt"

func experiement() {
	subset := []int{1, 2}

	// Append to res
	var res [][]int
	res = append(res, append([]int{}, subset...))

	// Append to res1
	var res1 [][]int
	res1 = append(res1, subset)

	// Modify the original subset
	subset[0] = 3

	// Print res and res1
	fmt.Println("res:", res)   // res: [[1 2]]
	fmt.Println("res1:", res1) // res1: [[3 2]]
}
