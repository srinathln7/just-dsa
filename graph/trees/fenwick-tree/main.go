package main

import "fmt"

// BIT: Binary Index Tree (aka) Fenwick tree
type BIT struct {
	sum []int
}

func initBIT(n int) *BIT {
	return &BIT{
		sum: make([]int, n+1),
	}
}

func (ft *BIT) update(idx, delta int) {

	// Update the sum of the node and all its children
	// To get the children, add the current idx with the result of the AND operation between the current idx and its two's complement
	for i := idx; i < len(ft.sum); i += i & -i {
		ft.sum[i] += delta
	}

}

func (ft *BIT) getPrefixSum(idx int) int {
	var sum int

	// Retrieve the sum from the current node and up until its parent
	// To get the parent, subract the result of the AND operation between the current idx and its two's complement from its current index
	for i := idx; i > 0; i -= i & -i {
		sum += ft.sum[i]
	}

	return sum
}

func main() {

	arr := []int{1, 2, 3, 4, 5}
	ft := initBIT(len(arr))

	for i, num := range arr {
		ft.update(i+1, num)
	}

	// Get prefix sum
	fmt.Println(ft.getPrefixSum(3))

	// Update value
	arr[3] = 10
	ft.update(3, arr[3])

	// Get latest updated prefix sum
	fmt.Println(ft.getPrefixSum(3))

}
