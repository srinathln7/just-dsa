package main

func canCompleteCircuit(gas []int, cost []int) int {

	// Key Idea: incoming -> gas, outgoing -> cost, net -> incoming[i] - outcoming[i]. To complete the entire trip
	// out net should always be positive or zero. A valid solution can exist only when the total amount of gas is greater than
	// the total cost of the gas spent.

	var incoming, outgoing int
	for i := 0; i < len(gas); i++ {
		incoming += gas[i]
		outgoing += cost[i]
	}

	if incoming < outgoing {
		return -1
	}

	startIdx, net := 0, 0
	for i := 0; i < len(gas); i++ {
		net += gas[i] - cost[i]

		// If net gas ever dips below zero, reset net and try the next index as the startIdx index
		if net < 0 {
			startIdx = i + 1
			net = 0
		}
	}

	return startIdx
}
