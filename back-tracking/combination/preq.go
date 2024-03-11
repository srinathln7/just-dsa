package main

import "fmt"

func main() {
	var a []int
	a = append(a, 1)
	a = append(a, 2)
	fmt.Println(a)

	changeMe(a)
	fmt.Println(a)

	changeMeRef(&a)
	fmt.Println(a)

	// Output
	// 	[1 2]
	// [3 4 5 6]
	// [3 4]
	// [7 8 9 10]
	// [7 8 9 10]
}

func changeMe(a []int) {
	a[0], a[1] = 3, 4
	a = append(a, 5)
	a = append(a, 6)
	fmt.Println(a)
}

func changeMeRef(a *[]int) {
	(*a)[0], (*a)[1] = 7, 8
	*a = append(*a, 9)
	*a = append(*a, 10)
	fmt.Println(*a)
}
