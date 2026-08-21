func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m := len(obstacleGrid)
    n := len(obstacleGrid[0])

    memo := make([][]int, m)
    for i := range memo {
        memo[i] = make([]int, n)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    return helper(obstacleGrid, m, n, 0, 0, memo)
}

func helper(obstacleGrid [][]int, m, n, row, col int, memo [][]int) int {  
    if row == m - 1 && col == n - 1 {
        if obstacleGrid[row][col] == 0 {
            return 1
        }

        return 0
    }

    if memo[row][col] != -1 {
        return memo[row][col]
    }

    var count int
    if row < m - 1 && obstacleGrid[row][col] == 0 {
       count += helper(obstacleGrid, m, n, row + 1, col, memo)
    }

    if col < n - 1 && obstacleGrid[row][col] == 0 {
        count += helper(obstacleGrid, m, n, row, col + 1, memo)
    }

    memo[row][col] = count
    return count
}