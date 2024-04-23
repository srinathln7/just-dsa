package main

func mergeTriplets(triplets [][]int, target []int) bool {

	// Key Idea: We adopt a greedy approach where in we compare every triplet with the target and discard them if any of the element in the
	// triplet are greater than the element in the target triplet. This is because since taking `max` of any other triplet with that will never
	// result in the target - we can safely dsicard it.

	var candidates [][]int
	for _, triplet := range triplets {
		if triplet[0] > target[0] || triplet[1] > target[1] || triplet[2] > target[2] {
			continue
		}
		candidates = append(candidates, triplet)
	}

	result := make([]bool, 3)
	for i := 0; i < 3; i++ {
		for _, candidate := range candidates {
			if target[i] == candidate[i] {
				result[i] = true
				break
			}
		}
	}

	return result[0] && result[1] && result[2]
}
