func isBipartite(graph [][]int) bool {
    n := len(graph)
    color := make([]int, n)

    for i := range n {
        if color[i] == 0 {
            color[i] = 1
            flag := dfs(graph, color, i)
            if !flag {
                return flag
            }
        }
    }    

    return true
}

func dfs(graph [][]int, color []int, idx int) bool {
    var flag bool = true
    
    for _, node := range graph[idx] {
        if color[node] == 0 {
            if color[idx] == 1 {
                color[node] = 2
            } else {
                color[node] = 1
            }

            flag = dfs(graph, color, node)
            if !flag {
                return flag
            }
        } else if color[node] == color[idx] {
            return false
        }
    }

    return flag
}
