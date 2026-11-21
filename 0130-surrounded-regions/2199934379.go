func solve(board [][]byte)  {
    sr := 0
    sc := 0
    l_sr := len(board)
    l_sc := len(board[0])
    visited := make([][]bool, len(board))

    for i := range board {
        visited[i] = make([]bool, len(board[i]))
    }

    for i := range board {
        for j := range board[i] {
           if i == sr || j == sc || i == l_sr - 1 || j == l_sc - 1 {
                if board[i][j] == 'O' && !visited[i][j] {
                    dfs(board, visited, i, j)     
                }
           }
        }
    }

    for i := range board {
        for j := range board[i] {
            if board[i][j] == 'O' && !visited[i][j] {
                board[i][j] = 'X'
            }
        }
    }

}

func dfs(board[][]byte, visited[][]bool, row int, col int) {
    visited[row][col] = true

    if row > 0 && !visited[row - 1] [col] && board[row - 1][col] == 'O' {
        dfs(board, visited, row - 1, col)
    }

    if col < len(board[0]) - 1 && !visited[row][col + 1] && board[row][col + 1] == 'O' {
        dfs(board, visited, row, col + 1)
    }

    if row < len(board) - 1 && !visited[row + 1][col] && board[row + 1][col] == 'O' {
        dfs(board, visited, row + 1, col)
    }

    if col > 0 && !visited[row][col - 1] && board[row][col - 1] == 'O' {
        dfs(board, visited, row, col - 1)
    }
}



