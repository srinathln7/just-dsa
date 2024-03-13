package main 

func allPathsSourceTarget(graph [][]int) [][]int {
    var result [][]int
    var path []int

    visited := make(map[int]bool)
    var dfs func(src, dst int)
    dfs = func(src, dst int) {
        visited[src] = true
        path = append(path, src)

       // Destination reached
       if src == dst { 
        // IMPORTANT: To make a copy of the current path and not directly assign path to it
        // This is because in the subsequent recursive stack frames, path will be modified
        // and in case we are appending only path, then all subsequent modifications will 
        // also be reflected in the `results`. 
         newPath := make([]int, len(path))
         copy(newPath, path)   
         result = append(result, newPath)    
       } 

        for _, neighbour := range graph[src] {
            if !visited[neighbour] {
                dfs(neighbour, dst)
            }
        }

       // Back track
       visited[src] = false
       path = path[:len(path)-1] 
    }

    n := len(graph)
    dfs(0, n-1)    
    return result
}
