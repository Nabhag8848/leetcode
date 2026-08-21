func uniquePathsIII(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])

    visited := make([][]int, m)
    total := 0
    startRow, startCol := -1, -1

    for i := range grid {
        visited[i] = make([]int, n)
        for j := range grid[i] {
            switch grid[i][j] {
            case -1:
                visited[i][j] = -1
            case 0:
                total++
            case 1:
                startRow, startCol = i, j
                total++
            }
        }
    }

    return helper(grid, m, n, startRow, startCol, 0, total, visited)
}

func helper(grid [][]int, m, n, row, col, steps, total int, visited [][]int) int {
    if grid[row][col] == 2 {
        if steps == total {
            return 1
        }
        return 0
    }

    var count int
    visited[row][col] = 1

    if row < m-1 && visited[row+1][col] == 0 {
        count += helper(grid, m, n, row+1, col, steps+1, total, visited)
    }
    if row > 0 && visited[row-1][col] == 0 {
        count += helper(grid, m, n, row-1, col, steps+1, total, visited)
    }
    if col < n-1 && visited[row][col+1] == 0 {
        count += helper(grid, m, n, row, col+1, steps+1, total, visited)
    }
    if col > 0 && visited[row][col-1] == 0 {
        count += helper(grid, m, n, row, col-1, steps+1, total, visited)
    }

    visited[row][col] = 0

    return count
}