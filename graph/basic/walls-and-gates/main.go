package main

import (
	"fmt"
	"math"
)

const INF = math.MaxInt32

func wallsAndGates(rooms [][]int) {

	// Key Idea: Since each cell in the grid has three possibilities `-1` => wall or obstacle  `0` => gate
	// and `INF` => empty and to fill each empty room with the distance to its nearest gate, let us start
	// from the reverse i.e. let us start from the gates and mark every adjacnt empty rooms with the nearest
	// distance values. Since the algorithm involves finding the shortest distance to its nearest gate, we can
	// deploy a BFS algorithm

	m, n := len(rooms), len(rooms[0])
	if m == 0 || n == 0 {
		return
	}

	var queue [][2]int
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if rooms[r][c] == 0 {
				queue = append(queue, [2]int{r, c})
			}
		}
	}

	// If no gates are found
	if len(queue) == 0 {
		return
	}

	// Run a BFS
	dist := 0
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(queue) > 0 {

		levelSize := len(queue)
		dist++

		for i := 0; i < levelSize; i++ {
			cell := queue[0]
			row, col := cell[0], cell[1]

			// Dequeue
			queue = queue[1:]

			for _, dir := range dirs {
				newRow, newCol := row+dir[0], col+dir[1]

				// Allow only in-bound cells and calculate distance only for empty rooms
				if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && rooms[newRow][newCol] == math.MaxInt32 {
					rooms[newRow][newCol] = dist
					queue = append(queue, [2]int{newRow, newCol})
				}
			}
		}
	}
}

func main() {

	rooms1 := [][]int{{INF, -1, 0, INF}, {INF, INF, INF, -1}, {INF, -1, INF, -1}, {0, -1, INF, INF}}
	wallsAndGates(rooms1)
	fmt.Println("Output 1:", rooms1)

	rooms2 := [][]int{{-1}}
	wallsAndGates(rooms2)
	fmt.Println("Output 2:", rooms2)

	//	Output 1: [[3 -1 0 1] [2 2 1 -1] [1 -1 2 -1] [0 -1 3 4]]
	// Output 2: [[-1]]
}
