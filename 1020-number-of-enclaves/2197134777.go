func numEnclaves(grid [][]int) int {
    arr := grid
    sr := 0
    sc := 0
    l_sr := len(arr)
    l_sc := len(arr[0])

    for i := range arr[sr] {
        if arr[sr][i] == 1 {
            dfs(arr, sr, i)
        }
    }

    for i := range arr[l_sr - 1] {
        if arr[l_sr - 1][i] == 1 {
            dfs(arr, l_sr - 1, i)
        }
    }

    for i := range l_sr {
        if arr[i][sc] == 1{
            dfs(arr, i, sc)
        }
    }

    for i := range l_sr {
        if arr[i][l_sc - 1] == 1 {
            dfs(arr, i, l_sc - 1)
        }
    }

    count := 0

    for i := range arr {
        for j := range arr[i] {
            if arr[i][j] == 1 {
                count++
            }
        }
    }

    return count
}

func dfs(grid [][]int, sr int, sc int) {
    grid[sr][sc] = 0

    if sr < len(grid) - 1 && grid[sr + 1][sc] == 1 {
        dfs(grid, sr + 1, sc)
    }

    if sr > 0 && grid[sr - 1][sc] == 1 {
        dfs(grid, sr - 1, sc)
    }   

    if sc < len(grid[0]) - 1 && grid[sr][sc + 1] == 1 {
        dfs(grid, sr, sc + 1)
    }

    if sc > 0 && grid[sr][sc - 1] == 1 {
        dfs(grid, sr, sc - 1)
    }

}



