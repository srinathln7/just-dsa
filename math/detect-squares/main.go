package main

type DetectSquares struct {
	ptsMap map[[2]int]int
	pts    [][]int
}

func Constructor() DetectSquares {
	return DetectSquares{
		ptsMap: make(map[[2]int]int),
		pts:    make([][]int, 0),
	}
}

func (this *DetectSquares) Add(point []int) {
	px, py := point[0], point[1]
	this.pts = append(this.pts, point)
	this.ptsMap[[2]int{px, py}]++
}

func (this *DetectSquares) Count(point []int) int {

	// Key Idea: For a given query point, we consider every other avaiable point in the list as if it was its diagonal point
	// Then we check for the square condition for other two points. If the square condition is satified, then we count the
	// freq of the other two points in top-left and bottom-right.

	count := 0
	px, py := point[0], point[1]
	for _, point := range this.pts {
		x, y := point[0], point[1]

		// Check for square condition
		if abs(py-y) != abs(px-x) || x == px || y == py {
			continue
		}
		count += this.ptsMap[[2]int{px, y}] * this.ptsMap[[2]int{x, py}]
	}
	return count
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */
