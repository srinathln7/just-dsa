package main

import "sort"

type car struct {
	pos int
	spd int
}

func carFleet(target int, position []int, speed []int) int {

	// Key Idea: Conceptually, we need to find if two car fleets intersect each other or not. To find this, we first
	// sort the car based on its position and then start from the reverse. This is because, only the last car's speed
	// can be easily determined depedning on speed of the previous car and the time it takes to arrive at the destination.

	n := len(position)
	cars := make([]car, n)
	for i := 0; i < n; i++ {
		cars[i] = car{pos: position[i], spd: speed[i]}
	}

	// Sort the car's based on it's position
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos <= cars[j].pos
	})

	// Stack contains the time that each car will take to arrive at the `target`
	var stack []float32

	// Start from right to left
	for i := n - 1; i >= 0; i-- {

		// Calculate time to reach the target
		time := float32(target-cars[i].pos) / float32(cars[i].spd)

		stack = append(stack, time)

		// Push this time to the stack
		m := len(stack)

		// IMPORTANT: If the penultimate car  reached the target before the ultimate car, it means
		// based on the given conditions, they must have definitely collided and became one car fleet
		// along the journey.
		if m >= 2 && stack[m-1] <= stack[m-2] {
			// Pop from the stack
			stack = stack[:m-1]
		}
	}

	return len(stack)
}
