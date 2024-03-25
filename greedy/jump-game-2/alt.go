package main

type KV struct {
	idx int
	val int
}

func Jump(nums []int) int {

	n := len(nums)
	start := KV{idx: 0, val: nums[0]}
	dst := KV{idx: n - 1, val: nums[n-1]}

	var minlen int

	visited := make(map[KV]bool)
	queue := []KV{start}
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			curr := queue[0]
			queue = queue[1:]

			if visited[curr] {
				continue
			}

			visited[curr] = true

			if curr == dst {
				return minlen
			}

			for j := curr.idx + 1; curr.val > 0 && j <= curr.idx+curr.val; j++ {
				if j >= 0 && j < n {
					neighbour := KV{idx: j, val: nums[j]}
					if !visited[neighbour] {
						queue = append(queue, neighbour)
					}
				}
			}
		}
		minlen++
	}

	// If dst is not reachable
	return -1
}
