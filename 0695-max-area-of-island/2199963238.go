func maxAreaOfIsland(grid [][]int) int {
    arr := grid
    rows := len(arr)
    cols := len(arr[0])

    maximum := 0

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if arr[i][j] == 1 {
                maximum = max(maximum, dfs(arr, i, j, 1))
            }
        }
    }

    return maximum
}

type Cell struct {
    row int
    col int
}

func dfs(grid [][]int, row, col int, count int) int {
    curr := count
    grid[row][col] = 0

    if row < len(grid) - 1 && grid[row + 1][col] == 1 {
        curr = dfs(grid, row + 1, col, curr + 1)
    }

    if row > 0 && grid[row - 1][col] == 1 {
        curr = dfs(grid, row - 1, col, curr + 1)
    }

    if col < len(grid[0]) - 1 && grid[row][col + 1] == 1 {
        curr = dfs(grid, row, col + 1, curr + 1)
    }

    if col > 0 && grid[row][col - 1] == 1 {
        curr = dfs(grid, row, col - 1, curr + 1)
    }

    return curr
}
