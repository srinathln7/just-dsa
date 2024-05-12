package main

func openLock(deadends []string, target string) int {

	// Key Idea: Since each slot can be incremented by 1 clockwise or anti-clockwise at every turn, for a given level(turn)
	// each node can have 8 different possibilites.Each slot has 2 possbilities (+ or - 1) and hence 4 slots will have a total
	// of 8 combinations. Now since we only need the MINIMUM number of turns to get to target without => This problem is equialent
	// to finding the shortest path from an init state ("0000") to the `target` while avoiding deadends. Hence we use BFS to effectively
	// solve the problem.

	deadendSet := make(map[string]bool)
	for _, deadend := range deadends {
		deadendSet[deadend] = true
	}

	// Base case - Init state is in deadend
	initState := "0000"
	if deadendSet[initState] {
		return -1
	}

	visited, count := make(map[string]bool), 0

	queue := []string{initState}
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {

			// IMPORTANT: Compare it with conventional BFS where we dequeue every element one-by-one from the front
			// However, in this case, this will TIME_OUT because of large volume of combinations and hence we do it for every level
			curr := queue[i]

			// Skip if already visited
			if visited[curr] {
				continue
			}

			visited[curr] = true

			if curr == target {
				return count
			}

			// Form the child nodes i.e. 8 possible combinations for each slot and append to the queue
			for j := 0; j < 4; j++ {

				// Turn the slot `j` clockwise and anti-clockwise
				nextForwardState, nextBackState := turnWheels(curr, j, 1), turnWheels(curr, j, -1)

				if !visited[nextForwardState] && !deadendSet[nextForwardState] {
					queue = append(queue, nextForwardState)
				}

				if !visited[nextBackState] && !deadendSet[nextBackState] {
					queue = append(queue, nextBackState)
				}

			}

		}

		// Move to the next level
		queue = queue[levelSize:]
		count++
	}

	return -1
}

func turnWheels(state string, slot int, dir int) string {

	digits := []byte(state)
	digit := digits[slot] - '0'

	// Account for wrapping around
	digit = ((digit + byte(dir)) + 10) % 10
	digits[slot] = digit + '0'

	return string(digits)
}
