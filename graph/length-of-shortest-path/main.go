package main

import "container/list"

func bfs(node int, target int, adjList map[int][]int) int {
	length := 0
	visit := make(map[int]bool)
	visit[node] = true
	queue := list.New()
	queue.PushBack(node)

	for queue.Len() > 0 {
		levelSize := queue.Len()
		for i := 0; i < levelSize; i++ {
			curr := queue.Remove(queue.Front()).(int)
			if curr == target {
				return length
			}
			for _, neighbor := range adjList[curr] {
				if !visit[neighbor] {
					visit[neighbor] = true
					queue.PushBack(neighbor)
				}
			}
		}
		length++
	}
	return length
}
