package main

func canPlaceFlowers(flowerbed []int, n int) bool {

	// Key Idea: With general observation, we can see that if the empty spots are in the middle we need three contiguous empty
	// spots to place a flower. To handle corner cases, we extrapolate the end spots with zeros

	var f []int
	f = append(append([]int{0}, flowerbed...), 0)
	for i := 1; i < len(f)-1; i++ {
		if f[i-1] == 0 && f[i] == 0 && f[i+1] == 0 {
			f[i] = 1
			n--
		}
	}

	return n <= 0
}
