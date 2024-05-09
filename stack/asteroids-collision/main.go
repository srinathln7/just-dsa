package main

func asteroidCollision(asteroids []int) []int {

	// Key Idea: Frame the condition for collision and use a stack to keep track of the remaining asteroids
	// One important observation is that the when the first asteroid moves to the left, it will never collide
	// which means the condition for collision is simply not just two asteroids moving in opposite directions
	// We need the asteroid we stack up to move to the right and the incoming asteroid to move to the left.
	// This is the condition for collision

	stack := []int{asteroids[0]}
	for i := 1; i < len(asteroids); i++ {
		incoming := asteroids[i]

		for len(stack) > 0 && incoming < 0 && stack[len(stack)-1] > 0 {
			if abs(incoming) > abs(stack[len(stack)-1]) {
				// Incoming asteroid wins and hence pop from the stack
				stack = stack[:len(stack)-1]
			} else if abs(incoming) < abs(stack[len(stack)-1]) {
				// Top of the stack wins and no need to append to the stack as its already present
				incoming = 0
			} else {
				// Condition for equality both asterioids are destroyed and hence pop from the stack
				incoming = 0
				stack = stack[:len(stack)-1]
			}

		}

		if incoming != 0 {
			stack = append(stack, incoming)
		}
	}

	return stack
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
