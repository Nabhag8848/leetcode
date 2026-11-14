type Cell struct {
    row int
    col int
}

func numIslands(grid [][]byte) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }

    rows := len(grid)
    cols := len(grid[0])

    islands := 0

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == '1' {
                islands++
                dfs(grid, i, j)
            }
        }
    }

    return islands
}

func dfs(grid [][]byte, row, col int) {
    rows := len(grid)
    cols := len(grid[0])

    if row < 0 || row >= rows || col < 0 || col >= cols || grid[row][col] == '0' {
        return
    }

    grid[row][col] = '0'

    dfs(grid, row + 1, col)
    dfs(grid, row - 1, col)
    dfs(grid, row, col + 1)
    dfs(grid, row, col - 1)
}
