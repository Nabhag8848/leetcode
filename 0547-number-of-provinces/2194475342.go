
func findCircleNum(isConnected [][]int) int {
    n := len(isConnected)
    visited := make([]bool, n)
    provinces := 0

    for i := 0; i < n; i++ {
        if !visited[i] {
            dfs(isConnected, visited, i)
            provinces++
        }
    }

    return provinces
}

func dfs(isConnected [][]int, visited []bool, curr int) {
    visited[curr] = true

    for idx, edge := range isConnected[curr] {
        if edge == 1 && !visited[idx] {
            dfs(isConnected, visited, idx)
        }
    }
}
