func eventualSafeNodes(graph [][]int) []int {
    n := len(graph)
    visited := make([]bool, n)
    pathVisited := make([]bool, n)

    for idx := range n {
        if !visited[idx] {
            dfs(graph, idx, visited, pathVisited)
        }
    }

    var result []int

    for i := range n {
        if pathVisited[i] == false {
            result = append(result, i)
        }
    }

    return result
}

func dfs(graph [][]int, idx int, visited []bool, pathVisited []bool) bool { 
    visited[idx] = true
    pathVisited[idx] = true

    for _, key := range graph[idx] {
        if !visited[key] {
            flag := dfs(graph, key, visited, pathVisited)
            if !flag {
                return flag
            }    
        } else if pathVisited[key] {
            return false
        } 
    }

    pathVisited[idx] = false
    return true
}


