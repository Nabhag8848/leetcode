func uniquePaths(m int, n int) int {
    memo := make([][]int, m)
    for i := range memo {
        memo[i] = make([]int, n)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    return helper(m, n, 0, 0, memo)
}

func helper(m, n, row, col int, memo [][]int) int {  
    if row == m - 1  && col == n - 1 {
        return 1
    }

    if memo[row][col] != -1 {
        return memo[row][col]
    }

    var count int
    if row < m - 1 {
       count += helper(m, n, row + 1, col, memo)
    }

    if col < n - 1 {
        count += helper(m, n, row, col + 1, memo)
    }

    memo[row][col] = count
    return count
}